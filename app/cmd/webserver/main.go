package main

import (
	poker "exoGitHubActions/app"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {

	log.Println("🏁 [INFO] Starting Poker Server version 0.1.14...")

	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("❌ [ERROR] Failed to load player store: %v", err)
	}
	defer close()

	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("could not listent on port 8080 %v", err)
	}
}
