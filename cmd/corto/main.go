package main

import (
	"log"

	co "github.com/alwindoss/corto"
	"github.com/alwindoss/corto/internal/corto"
	"github.com/gorilla/mux"
)

func main() {
	log.Printf("welcome to corto short url service")
	log.Printf("Added as part of Language Server")
	cfg := &co.Config{
		DBLoc: "",
	}
	r := mux.NewRouter()
	corto.Run(cfg, r, "8080")
}
