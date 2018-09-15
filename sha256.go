package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"crypto/sha256"

	"github.com/gorilla/mux"
)

func transformToSha256(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// ip := r.FormValue("ip");

	h := sha256.New()

	sha := base64.URLEncoding.EncodeToString(h.Sum([]byte(params["ip"])))

	json.NewEncoder(w).Encode(sha)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/transform/{ip}", transformToSha256)

	log.Fatal(http.ListenAndServe(":8080", r))
}
