package corto

import (
	"fmt"
	"net/http"

	"github.com/alwindoss/corto"
	"github.com/gorilla/mux"
)

type errorResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type createShortURLRequest struct {
	URL        string `json:"url"`
	APIKey     string `json:"api-key"`
	Alias      string `json:"alias"`
	UserName   string `json:"user-name"`
	ExpiryDate string `json:"expiry-date"`
}

type createShortURLResponse struct {
	ShortURL string    `json:"short-url"`
	Error    errorResp `json:"error"`
}

type deleteShortURLRequest struct {
	ShortURL string `json:"short-url"`
	APIKey   string `json:"api-key"`
}

type deleteShortURLResponse struct {
	ShortURL string    `json:"short-url"`
	Error    errorResp `json:"error"`
}

type createShortURLHandler struct {
	mgr corto.ShortURLManager
}

func (createShortURLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Created a short URL")
}

type deleteShortURLHandler struct {
	mgr corto.ShortURLManager
}

func (deleteShortURLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleted short URL")
}

type fetchShortURLHandler struct {
	mgr corto.ShortURLManager
}

func (fetchShortURLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	kv := mux.Vars(r)
	key := kv["key"]
	// fmt.Fprintf(w, "%s is the key you tried to fetch", key)
	http.Redirect(w, r, "https://google.com/"+key, http.StatusSeeOther)
}
