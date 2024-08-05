package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	db "github.com/khudaibirdin/GoLangModules/database_actions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

type User struct {
	Login      string `db:"login"`
	Password   string `db:"password"`
	Categories string `db:"categories"`
}


type PageStruct struct {
	Categories struct {
		Code []string
		Name []string
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		status, _ := VerifyUser(username, password)
		if status {
			session, _ := store.Get(r, "session-name")
			session.Values["authenticated"] = true
			session.Values["username"] = username
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusFound)
		}
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	renderTemplate(w, "login", nil)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if r.Method == http.MethodPost {
		session, _ := store.Get(r, "session-name")
		session.Values["authenticated"] = false
		session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusFound)
	}

}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	db := db.Db{Path: "./database.db"}
	condition := fmt.Sprintf(`login = "%s"`, session.Values["username"])
	user := User{}
	db.GetRowByCondition("user", "*", condition, &user)

	view := PageStruct{
		Categories: struct{Code []string; Name []string}{
			Code: []string{""},
			Name: strings.Split(user.Categories, ", "),
		},
	}
	renderTemplate(w, "main", view)
}

func StatisticHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	renderTemplate(w, "statistic", nil)
}

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/auth", http.StatusFound)
		return
	}

	renderTemplate(w, "account", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates := template.Must(template.ParseFiles("./templates/base.html", "./templates/"+tmpl+".html"))
	err := templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
