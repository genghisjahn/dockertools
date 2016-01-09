package main

import (
	"flag"
	"log"
	"os"
	"testing"
)

func TestInsertRecrod(t *testing.T) {

}

func TestMain(m *testing.M) {
	persistDB := flag.Bool("persistdb", false, "True, leave the DB container running")
	flag.Parse()
	keepDB, errSetup := setup("DBNAME")
	if errSetup != nil {
		log.Println("Setup Error:", errSetup)
	}
	var code int
	if errSetup == nil {
		code = m.Run()
	}
	//args := os.Args
	// for _, v := range args {
	// 	if v == "-persist=true" {
	// 		*persistDB = true
	// 		break
	// 	}
	// }
	if !*persistDB && !keepDB {
		errShutdown := shutdown()
		if errShutdown != nil {
			log.Println(errShutdown)
		}
	}
	os.Exit(code)
}
