package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/genghisjahn/dockertools/docker"
	_ "github.com/lib/pq"
)

type DBInfo struct {
	DockerMachine string `json:"docker_machine"`
	ContainerName string `json:"container_name"`
	Host          string `json:"host"`
	DBName        string `json:"db_name"`
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
}

//var db *DBInfo

func getDBConn(info DBInfo) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", info.Host, info.UserName, info.DBName, info.Password))
}

func getConnectInfo() (*DBInfo, error) {
	d := &DBInfo{}
	c, cErr := ioutil.ReadFile("data/connect.json")
	if cErr != nil {
		return nil, cErr
	}
	jErr := json.Unmarshal(c, &d)
	if jErr != nil {
		return nil, jErr
	}

	hostip, ipErr := docker.GetHostIP(d.DockerMachine)
	if ipErr != nil {
		return nil, ipErr
	}
	d.Host = hostip
	return d, nil
}

func setup() (bool, error) {
	var db *DBInfo
	var err error

	db, err = getConnectInfo()

	info, infoErr := docker.InspectContainer(db.ContainerName)
	if infoErr != nil {
		if _, ok := infoErr.(*docker.ContainerNotFoundError); !ok {
			return false, infoErr
		}
	}
	if info.State.Running {
		log.Printf("Container %s is already running. Started At: %s\n", db.ContainerName, info.State.StartedAt)
		return true, nil
	}
	err = setupDBContainer()
	if err != nil {
		return false, err
	}
	err = initData(*db)
	if err != nil {
		return false, err
	}
	return false, nil
}

func setupDBContainer() error {
	db, err := getConnectInfo()
	if err != nil {
		return err
	}
	argtemplate := "-p 5432:5432  --name %s  -e POSTGRES_PASSWORD=%s -e POSTGRES_DB=%s -e POSTGRES_USER=%s -d postgres"
	runargs := fmt.Sprintf(argtemplate, db.ContainerName, db.Password, db.DBName, db.UserName)
	runErr := docker.Run(runargs, false)
	if runErr != nil {
		return runErr
	}
	return nil
}

func shutdown() error {
	var err error
	var db *DBInfo
	db, err = getConnectInfo()
	if err != nil {
		return err
	}
	err = docker.StopContainer(db.ContainerName, false)
	if err != nil {
		if _, ok := err.(*docker.ContainerNotFoundError); ok {
			log.Printf("Error stoping %s. Container cannot be found.\n", db.ContainerName)
		}
		return err
	}
	err = docker.RemoveContainer(db.ContainerName, false)
	if err != nil {
		if _, ok := err.(*docker.ContainerNotFoundError); ok {
			log.Printf("Error removing %s. Container cannot be found.\n", db.ContainerName)
		}
		return err
	}
	return nil
}

func initData(db DBInfo) error {
	hostip, err := docker.GetHostIP(db.DockerMachine)
	if err != nil {
		return err
	}

	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", hostip, db.UserName, db.Password, db.DBName)
	maxAttempts := 20
	cn, _ := sql.Open("postgres", constr)
	defer cn.Close()
	var errExec error
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		errPing := cn.Ping()
		if errPing == nil {
			log.Println("DB connect succeeded!")
			break
		}
		log.Println(errPing)
		time.Sleep(time.Duration(attempts) * time.Second)
	}
	s, sErr := ioutil.ReadFile("data/schema.sql")
	if sErr != nil {
		return sErr
	}
	_, errExec = cn.Exec(string(s))
	if errExec != nil {
		return errExec
	}
	d, dErr := ioutil.ReadFile("data/data.sql")
	if dErr != nil {
		return dErr
	}
	//TODO what happens if data doesn't match schema?
	_, errExec = cn.Exec(string(d))
	return errExec
}
