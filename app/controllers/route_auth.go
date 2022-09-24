package controllers

import (
	"log"
	"net/http"

	"example.com/school/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		operator := models.Operator{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := operator.CreateOperator(); err != nil {
			log.Fatalln(err)
		}

		http.Redirect(w, r, "/", 302)
	}
}
