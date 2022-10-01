package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/wspirrat/swt/swt"
)

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	log.Println("listening on http://localhost:8090")
	http.ListenAndServe(":8090", nil)
}

func encode(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	m := make(map[string]interface{})
	m[key] = value
	mencode := swt.EncodeSWT(m)
	fmt.Fprintf(w, mencode)
}

func decode(w http.ResponseWriter, r *http.Request) {
	payload := r.URL.Query().Get("swt")
	res := swt.DecodeSWT(payload)

	mJson, err := json.Marshal(res.Payload)
	if err != nil {
		return
	}

	fmt.Fprintf(w, string(mJson))
}
