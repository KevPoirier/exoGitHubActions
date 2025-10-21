package main

import (
	poker "exoGitHubActions/app"
	_ "exoGitHubActions/docs"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// /test
const dbFileName = "game.db.json"

func main() {

	log.Println("🏁 [INFO] Starting Poker Server version 0.1.15...")

	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("❌ [ERROR] Failed to load player store: %v", err)
	}
	defer close()
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	server := poker.NewPlayerServer(store)
	mux := http.NewServeMux()
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	mux.Handle("/", server)
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("could not listent on port 8081 %v", err)
	}
}
