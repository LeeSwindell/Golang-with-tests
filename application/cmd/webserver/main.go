package main

import (
	poker "application"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5001", server); err != nil{
		log.Fatalf("couldnt listen on port 5001, %v", err)
	}
}