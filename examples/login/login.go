package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/wspirrat/swt/swt"
)

func main() {
	swt.EXPIRE_TIME = time.Second * 2
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	log.Println("listening on http://localhost:80")
	http.ListenAndServe(":80", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	safeToken := r.URL.Query().Get("token")
	token := swt.DecodeSWT(safeToken)
	if !token.IsPayloadNil() {
		payload := token.Payload.(map[string]interface{})
		fmt.Println(payload)

		if payload["isLogged"].(bool) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, `username: %s <br> password: %s <br> isLogged: %t
				<br><br>
				This token is only valid for 2 seconds, refresh it and it will be not valid.`,
				payload["username"].(string), payload["password"].(string), payload["isLogged"].(bool))
		}
	} else {
		fmt.Fprintf(w, "%s", `You dont have access to this website. go to <a href="/login">login page</a>`)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		m := make(map[string]interface{})
		m["username"] = r.FormValue("uname")
		m["password"] = r.FormValue("psw")
		m["isLogged"] = true
		swt_cargo := swt.EncodeSWTcustom(m)

		http.Redirect(w, r, fmt.Sprintf("http://localhost:80?token=%s", swt_cargo), http.StatusMovedPermanently)
	}

	tmpl := template.Must(template.ParseFiles("login.html"))
	tmpl.Execute(w, nil)
}
