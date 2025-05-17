package main

import (
	"go-mvc-app/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/users", controllers.GetUsers)
	http.HandleFunc("/add-user", controllers.AddUser)

	http.ListenAndServe(":8080", nil)
}
