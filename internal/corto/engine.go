package corto

import (
	"log"
	"net/http"

	"github.com/alwindoss/corto"
	"github.com/gorilla/mux"
)

// Run runs the engine
func Run(cfg *corto.Config, r *mux.Router, port string) {
	mgr := NewShortURLManager(cfg)
	createManagerServiceEngine(r, mgr)
	createFetchShortURLServiceEnginer(r, mgr)

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func createManagerServiceEngine(r *mux.Router, mgr corto.ShortURLManager) {
	r.Path("/url").Methods(http.MethodPost).Handler(createShortURLHandler{mgr})
	r.Path("/url").Methods(http.MethodDelete).Handler(deleteShortURLHandler{mgr})

}

func createFetchShortURLServiceEnginer(r *mux.Router, mgr corto.ShortURLManager) {
	r.Path("/{key}").Methods(http.MethodGet).Handler(fetchShortURLHandler{mgr})
}
