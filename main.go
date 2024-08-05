package main

import (
	"log"
	"net/http"
	"os/exec"
)


func main() {
	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/account", AccountHandler)
	http.HandleFunc("/statistic", StatisticHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	err := exec.Command("explorer", "http://localhost:8080/").Run()
	if err != nil {
		log.Println("Ошибка запуска браузера:", err)
	}
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
 