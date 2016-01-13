package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
)

var db *DBInfo

var testDBName string

func TestMain(m *testing.M) {
	var err error
	db, err = getConnectInfo()
	if err != nil {
		log.Fatal(err)
	}
	persistDB := flag.Bool("persistdb", false, "True, leave the DB container running")
	killDB := flag.Bool("killdb", false, "True, kill the DB Container and return.  No tests are run.")
	flag.Parse()
	if *killDB {
		log.Println("Shutting down container...")
		errShutdown := shutdown()
		if errShutdown != nil {
			log.Println(errShutdown)
		}
		log.Println("Shutdown complete.")
		return
	}
	keepDB, errSetup := setup()
	if errSetup != nil {
		log.Println("Setup Error:", errSetup)
	}
	var code int
	if errSetup == nil {
		testDB, errTest := getConnectInfo()
		if errTest != nil {
			log.Println(errTest)
			return
		}
		testDBName = fmt.Sprintf("%s_test", testDB.DBName)
		errCreate := createDB(testDBName)
		if errCreate != nil {
			log.Println(errCreate)
			return
		}
		code = m.Run()
		if !*persistDB && !keepDB {
			errShutdown := shutdown()
			if errShutdown != nil {
				log.Println(errShutdown)
			}
		}
	} else {
		log.Println(errSetup)
	}
	os.Exit(code)
}

func createDB(newdbname string) error {
	killconnstmt := `SELECT pg_terminate_backend(pg_stat_activity.pid)
										FROM pg_stat_activity
										WHERE datname = current_database()
  									AND pid <> pg_backend_pid();`
	dropstmt := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", newdbname)
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", db.Host, db.UserName, db.Password, db.DBName)
	cn, _ := sql.Open("postgres", constr)
	defer cn.Close()
	createtemplate := "CREATE DATABASE %s WITH TEMPLATE %s OWNER %s;"
	createstatment := fmt.Sprintf(createtemplate, newdbname, db.DBName, db.UserName)

	//Kill the connections so we can drop/copy the DB
	_, errKillCn := cn.Exec(killconnstmt)
	if errKillCn != nil {
		return errKillCn
	}

	//Drop the test database if it already exists, we don't want it anymore.
	_, errDrop := cn.Exec(dropstmt)
	if errDrop != nil {
		return errDrop
	}
	//Make a copy of the test database from the origin_db
	_, errExec := cn.Exec(createstatment)
	return errExec
}

func TestInsertMovie(t *testing.T) {
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", db.Host, db.UserName, db.Password, testDBName)
	cn, _ := sql.Open("postgres", constr)
	defer cn.Close()
	insertMovie := "insert into movie (title) values ('Raiders of the Lost Ark');"
	_, errExec := cn.Exec(insertMovie)
	if errExec != nil {
		t.Error(errExec)
	}

}

func TestInsertActor(t *testing.T) {
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", db.Host, db.UserName, db.Password, testDBName)
	cn, _ := sql.Open("postgres", constr)
	defer cn.Close()
	insertActor := "insert into actor (name) values ('Karen Allen');"
	_, errExec := cn.Exec(insertActor)
	if errExec != nil {
		t.Error(errExec)
	}
}

func TestAddActorToMovie(t *testing.T) {
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", db.Host, db.UserName, db.Password, testDBName)
	cn, _ := sql.Open("postgres", constr)
	defer cn.Close()
	insertMovieActor := `insert into movieactor (movie_id,actor_id) values (4,5);
	insert into movieactor (movie_id,actor_id) values (4,1);
	`
	_, errExec := cn.Exec(insertMovieActor)
	if errExec != nil {
		t.Error(errExec)
	}
}

func TestCheckRaiders(t *testing.T) {
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", db.Host, db.UserName, db.Password, testDBName)
	cn, _ := sql.Open("postgres", constr)
	defer cn.Close()
	queryRaiders := "select a.name from actor a, movieactor ma where ma.movie_id = 4 and ma.actor_id = a.id;"

	rows, err := cn.Query(queryRaiders)
	actornames := []string{}
	for rows.Next() {
		var name string
		if errScan := rows.Scan(&name); err != nil {
			t.Fatal(errScan)
		}
		actornames = append(actornames, name)

	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
	if actornames[0] != "Harrison Ford" {
		t.Errorf("Harrison Ford should have been the first row. Received: %s\n", actornames[0])
	}
	if actornames[1] != "Karen Allen" {
		t.Errorf("Karen Allen should have been the second row. Received: %s\n", actornames[0])
	}
}

func TestCheckHarrisonFord(t *testing.T) {
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", db.Host, db.UserName, db.Password, testDBName)
	cn, _ := sql.Open("postgres", constr)
	defer cn.Close()
	queryFord := "select m.title from movie m, movieactor ma where ma.movie_id = m.id and ma.actor_id = 1;"

	rows, err := cn.Query(queryFord)
	movies := []string{}
	for rows.Next() {
		var title string
		if errScan := rows.Scan(&title); err != nil {
			t.Fatal(errScan)
		}
		movies = append(movies, title)

	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
	expected := "[Star Wars Apocalypse Now Raiders of the Lost Ark]"
	result := fmt.Sprintf("%v", movies)
	if string(result) != expected {
		t.Errorf("Expected %s\nReceived: %s\n", expected, result)
	}
}

func TestInsertActorA(t *testing.T) {
	dbname := "testa"
	createDB(dbname)
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", db.Host, db.UserName, db.Password, dbname)
	cn, _ := sql.Open("postgres", constr)
	defer cn.Close()
	insertActor := "insert into actor (name) values ('Clint Eastwood');"
	_, errExec := cn.Exec(insertActor)
	if errExec != nil {
		t.Error(errExec)
	}
}

func TestInsertActorB(t *testing.T) {
	dbname := "testb"
	createDB(dbname)
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", db.Host, db.UserName, db.Password, dbname)
	cn, _ := sql.Open("postgres", constr)
	defer cn.Close()
	insertActor := "insert into actor (name) values ('Morgan Freeman');"
	_, errExec := cn.Exec(insertActor)
	if errExec != nil {
		t.Error(errExec)
	}
}
