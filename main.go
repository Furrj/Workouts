package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/add", addWorkout)

	log.Panic(http.ListenAndServe(":5000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(r.FormValue("name"))
	}

	http.ServeFile(w, r, "home.html")
}

func addWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Panicf("Error: %+v\n", err)
		}

		for key, value := range r.PostForm {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	http.ServeFile(w, r, "add_workout.html")
}
