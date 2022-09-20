package main

import (
	"log"
	"net/http"
	"os"
)

// type InMemoryPlayerStore struct {
// 	store map[string]int
// }

// func newInMemoryPlayerStore() *InMemoryPlayerStore {
// 	return &InMemoryPlayerStore{map[string]int{}}
// }

// func(i *InMemoryPlayerStore) GetPlayerScore(name string) int {
// 	return i.store[name]
// }

// func(i *InMemoryPlayerStore) RecordWin(name string) {
// 	i.store[name] ++
// }

// func(i *InMemoryPlayerStore) GetLeague() League {
// 	var league League
// 	for name, wins := range i.store {
// 		league = append(league, Player{name, wins})
// 	}
// 	return league
// }

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening db, %v", err)
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := NewPlayerServer(store)
	if err := http.ListenAndServe(":5001", server); err != nil{
		log.Fatalf("couldnt listen on port 5001, %v", err)
	}
}