package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fName, lName, pass string) *Account {
	acc, err := NewAccount(fName, lName, pass)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account => ", acc.Number)

	return acc
}

// 52035
func seedAccounts(s Storage) {
	seedAccount(s, "jim", "page", "jalal1417")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStrore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("sending the database")
		seedAccounts(store)
	}
	// seed stuff

	apiServer := NewAPIServer(":3000", store)
	apiServer.Run()
}
