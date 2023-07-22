package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("go bank")

	store, err := NewPostgresStrore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	apiServer := NewAPIServer(":3000", store)
	apiServer.Run()
}
