package controllers

import (
	"log"
	"net/http"

	"example.com/school/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/schools", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
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

func schoolNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "school_new")
	}
}

func schoolSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		operator, err := sess.GetOperatorBySession()
		if err != nil {
			log.Println(err)
		}
		name := r.PostFormValue("name")
		if err := operator.CreateSchool(name); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/schools", 302)
	}
}

func schoolEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetOperatorBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetSchool(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "school_edit")
	}
}

func schoolUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		operator, err := sess.GetOperatorBySession()
		if err != nil {
			log.Println(err)
		}
		name := r.PostFormValue("name")
		t := &models.School{ID: id, Name: name, OperatorID: operator.ID}
		if err := t.UpdateSchool(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/schools", 302)
	}
}
