package controllers

import (
	"log"
	"net/http"

	"example.com/school/app/models"
)

func schoolNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "school_new")
	}
}

func schoolSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
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

		http.Redirect(w, r, "/schools", http.StatusFound)
	}
}

func schoolEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
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
		http.Redirect(w, r, "/login", http.StatusFound)
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
		http.Redirect(w, r, "/schools", http.StatusFound)
	}
}

func schoolDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		_, err := sess.GetOperatorBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetSchool(id)
		if err != nil {
			log.Println(err)
		}
		if err := t.DeleteSchool(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/schools", http.StatusFound)
	}
}

func schoolShow(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		_, err := sess.GetOperatorBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetSchool(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "school_show")
	}
}
