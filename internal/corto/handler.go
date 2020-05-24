package corto

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alwindoss/corto"
	"github.com/gorilla/mux"
)

type errorResp struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type createShortURLRequest struct {
	URL        string    `json:"url"`
	APIKey     string    `json:"api-key"`
	Alias      string    `json:"alias"`
	UserName   string    `json:"user-name"`
	ExpiryDate time.Time `json:"expiry-date"`
}

type createShortURLResponse struct {
	ShortURL string    `json:"short-url,omitempty"`
	Error    errorResp `json:"error,omitempty"`
}

type deleteShortURLRequest struct {
	ShortURL string `json:"short-url"`
	APIKey   string `json:"api-key"`
}

type deleteShortURLResponse struct {
	ShortURL string    `json:"short-url,omitempty"`
	Error    errorResp `json:"error,omitempty"`
}

type createShortURLHandler struct {
	mgr corto.ShortURLManager
}

func (h createShortURLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	createReq := &createShortURLRequest{}
	err := json.NewDecoder(r.Body).Decode(createReq)
	if err != nil {
		log.Printf("unable to decode the request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "unable to create url")
	}
	defer r.Body.Close()
	shortURLKey, err := h.mgr.Create(createReq.APIKey, createReq.URL, createReq.Alias, createReq.UserName, createReq.ExpiryDate)
	if err != nil {
		log.Printf("unable to create the short url")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "unable to create url")
	}
	resp := &createShortURLResponse{
		ShortURL: shortURLKey,
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("unable to create the short url")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "unable to create url")
	}
}

type deleteShortURLHandler struct {
	mgr corto.ShortURLManager
}

func (h deleteShortURLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleted short URL")
}

type fetchShortURLHandler struct {
	mgr corto.ShortURLManager
}

func (h fetchShortURLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	kv := mux.Vars(r)
	key := kv["key"]
	originalURL, err := h.mgr.FetchURL("", key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "URL not found")
	}
	// fmt.Fprintf(w, "%s is the key you tried to fetch", key)

	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}
