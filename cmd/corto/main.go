package main

import (
	"log"

	"github.com/alwindoss/corto/internal/corto"
	"github.com/gorilla/mux"
)

func main() {
	log.Printf("welcome to corto short url service")
	r := mux.NewRouter()
	corto.Run(r, "8080")
}
