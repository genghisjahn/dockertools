package main

import (
	"flag"
	"log"
	"os"
	"testing"
)

func TestInsertMovie(t *testing.T) {

}

func TestInsertActor(t *testing.T) {

}

func TestCheckRaiders(t *testing.T) {

}

func TestCheckHarrisonFord(t *testing.T) {

}

func TestMain(m *testing.M) {
	persistDB := flag.Bool("persistdb", false, "True, leave the DB container running")
	flag.Parse()
	keepDB, errSetup := setup("DemoContainer")
	if errSetup != nil {
		log.Println("Setup Error:", errSetup)
	}
	var code int
	if errSetup == nil {
		code = m.Run()
	}

	if !*persistDB && !keepDB {
		errShutdown := shutdown()
		if errShutdown != nil {
			log.Println(errShutdown)
		}
	}
	os.Exit(code)
}
