package loginView

import (
	"../controllers"
	"fmt"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.gtpl")
		t.Execute(w, nil)
	}
}

func Ajax(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		pwd := loginController.Select(r.Form.Get("login"))
		password := r.Form.Get("pwd")
		if pwd == password {
			fmt.Fprintf(w, "1")
		} else {
			fmt.Fprintf(w, "2")
		}
	}
}
