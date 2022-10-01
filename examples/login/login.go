package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"swt/swt"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	log.Println("listening on http://localhost:80")
	http.ListenAndServe(":80", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	safeToken := r.URL.Query().Get("token")
	token := swt.DecodeSWT(safeToken)
	payload := token.Payload

	fmt.Println(payload)
	if len(payload) > 1 {
		if payload["isLogged"].(bool) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "Your username is: %s", payload["username"].(string))
		}
	} else {
		fmt.Fprintf(w, "You dont have access to this website.")
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		m := make(map[string]interface{})
		m["username"] = r.FormValue("uname")
		m["isLogged"] = true
		swt_cargo := swt.EncodeSWT(m)

		http.Redirect(w, r, fmt.Sprintf("http://localhost:80?token=%s", swt_cargo), 301)
	}

	tmpl := template.Must(template.ParseFiles("login.html"))
	tmpl.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	payload := r.URL.Query().Get("swt")
	res := swt.DecodeSWT(payload)

	mJson, err := json.Marshal(res.Payload)
	if err != nil {
		return
	}

	fmt.Fprintf(w, string(mJson))
}