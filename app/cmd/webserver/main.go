package main

import (
	poker "exoGitHubActions/app"
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
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("could not listent on port 5000 %v", err)
	}
}
