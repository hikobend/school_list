package controllers

import (
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/schools", http.StatusFound)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		operator, err := sess.GetOperatorBySession()
		if err != nil {
			log.Println(err)
		}
		schools, _ := operator.GetSchoolByOperator()
		operator.Schools = schools
		generateHTML(w, operator, "layout", "private_navbar", "index")
	}
}
