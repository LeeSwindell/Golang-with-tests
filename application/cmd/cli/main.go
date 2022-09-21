package main

import (
	poker "application"
	"fmt"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("lets play poker!")
	fmt.Println("Type {name} wins to record win")

	poker.NewCLI(store, os.Stdin).PlayPoker()
}