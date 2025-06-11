package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ishubhampatidar/students-api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// database setup
	// router setup
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})

	// server setup
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("Server started: %s", cfg.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}

}
