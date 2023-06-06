package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page!")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	fmt.Fprint(w, "User ID: ", userID)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/user/{userID}", UserHandler)

	server := &http.Server{
		Addr:         "127.0.0.1:3000",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
