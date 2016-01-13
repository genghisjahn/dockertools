package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	persistDB := flag.Bool("persistdb", false, "True, leave the DB container running")
	flag.Parse()
	keepDB, errSetup := setup()
	if errSetup != nil {
		log.Println("Setup Error:", errSetup)
	}
	var code int
	if errSetup == nil {
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

func TestInsertMovie(t *testing.T) {
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", db.Host, db.UserName, db.Password, db.DBName)
	cn, _ := sql.Open("postgres", constr)
	defer cn.Close()
	insertMovie := "insert into movie (title) values ('Raiders of the Lost Ark');"
	_, errExec := cn.Exec(insertMovie)
	if errExec != nil {
		t.Error(errExec)
	}

}

func TestInsertActor(t *testing.T) {
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", db.Host, db.UserName, db.Password, db.DBName)
	cn, _ := sql.Open("postgres", constr)
	defer cn.Close()
	insertActor := "insert into actor (name) values ('Karen Allen');"
	_, errExec := cn.Exec(insertActor)
	if errExec != nil {
		t.Error(errExec)
	}
}

func TestAddActorToMovie(t *testing.T) {
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", db.Host, db.UserName, db.Password, db.DBName)
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
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", db.Host, db.UserName, db.Password, db.DBName)
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
	constr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", db.Host, db.UserName, db.Password, db.DBName)
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
