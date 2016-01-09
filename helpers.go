package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/genghisjahn/dockertools/docker"
	_ "github.com/lib/pq"
)

type DBInfo struct {
	ContainerName string
	Host          string
	DBName        string
	UserName      string
	Password      string
}

func getDBConn(info DBInfo) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", info.Host, info.UserName, info.DBName, info.Password))
}

func setup(name string) (bool, error) {
	var err error

	info, infoErr := docker.InspectContainer(name)
	if infoErr != nil {
		if _, ok := infoErr.(*docker.ContainerNotFoundError); !ok {
			return false, infoErr
		}
	}
	if info.State.Running {
		log.Printf("Container %s is already running. Started At: %s\n", name, info.State.StartedAt)
		return true, nil
	}
	err = setupDBContainer()
	if err != nil {
		return false, err
	}
	err = initData()
	if err != nil {
		return false, err
	}
	return false, nil
}

var dbInfo DBInfo

func setupDBContainer() error {
	hostip, err := docker.GetHostIP("default")
	if err != nil {
		return err
	}

	dbInfo.ContainerName = "dockerDemoDB"
	dbInfo.Host = hostip
	dbInfo.DBName = "dockerdemo"
	dbInfo.UserName = "demo"
	dbInfo.Password = "abcd1234"

	argtemplate := "-p 5432:5432  --name %s  -e POSTGRES_PASSWORD=%s -e POSTGRES_DB=%s -e POSTGRES_USER=%s -d postgres"
	runargs := fmt.Sprintf(argtemplate, dbInfo.ContainerName, dbInfo.Password, dbInfo.DBName, dbInfo.UserName)
	runErr := docker.Run(runargs, false)
	if runErr != nil {
		return runErr
	}
	return nil
}

func shutdown() error {
	var err error
	err = docker.StopContainer(dbInfo.ContainerName, false)
	if err != nil {
		return err
	}
	err = docker.RemoveContainer(dbInfo.ContainerName, false)
	if err != nil {
		return err
	}
	return nil
}

func initData() error {
	hostip, err := docker.GetHostIP("default")
	if err != nil {
		return err
	}
	dbInfo.ContainerName = "dockerDemoDB"
	dbInfo.Host = hostip
	dbInfo.DBName = "dockerdemo"
	dbInfo.UserName = "demo"
	dbInfo.Password = "abcd1234"
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", hostip, dbInfo.UserName, dbInfo.Password, dbInfo.DBName)
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
	_, errExec = cn.Exec("INIT DATA")
	if errExec != nil {
		return errExec
	}
	_, errExec = cn.Exec("Insert DATA")
	return errExec
}
