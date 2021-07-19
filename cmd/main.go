package main

import (
	"fmt"
	poker "lexcao.io/learn-go-with-tests/cmd/webserver"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type '{name} wins' to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("opening %s %v", dbFileName, err)
	}

	store := poker.NewFileSystemPlayerStore(db)
	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
