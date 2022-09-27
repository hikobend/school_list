package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"text/template"

	"example.com/school/app/models"
	"example.com/school/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(_ http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("invalid Error")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/schools/(edit|update|delete)/([0-9]+)")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
}

func StartmainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/schools", index)
	http.HandleFunc("/schools/new", schoolNew)
	http.HandleFunc("/schools/save", schoolSave)
	http.HandleFunc("/schools/edit/", parseURL(schoolEdit))
	http.HandleFunc("/schools/update/", parseURL(schoolUpdate))
	http.HandleFunc("/schools/delete/", parseURL(schoolDelete))
	http.HandleFunc("/schools/club/new", clubNew)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
