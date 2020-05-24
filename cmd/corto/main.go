package main

import (
	"log"

	co "github.com/alwindoss/corto"
	"github.com/alwindoss/corto/internal/corto"
	"github.com/gorilla/mux"
)

func main() {
	log.Printf("welcome to corto short url service")
	cfg := &co.Config{
		DBLoc: "",
	}
	r := mux.NewRouter()
	corto.Run(cfg, r, "8080")
}
