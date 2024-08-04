package main

import (
    "html/template"
    "net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")

        // Здесь должна быть проверка логина и пароля
        // Для примера используем простой условный оператор
        if username == "admin" && password == "password" {
            http.Redirect(w, r, "/", http.StatusFound)
            return
        }

        // Если логин или пароль неправильные, перенаправляем обратно на страницу логина
        http.Redirect(w, r, "/login", http.StatusFound)
        return
    }
    renderTemplate(w, "auth", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    templates := template.Must(template.ParseFiles("./templates/base.html", "./templates/"+tmpl+".html"))
    err := templates.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
