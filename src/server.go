package main

import (
	"context"
	"fmt"
	"github.com/Furrj/Workouts/src/internal/DB"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Getenv("URL")
	fmt.Println(url)
	db := DB.InitDB(url)

	err := db.Conn.Ping(context.Background())
	if err != nil {
		log.Panicf("Error pinging db: %+v\n", err)
	} else {
		fmt.Printf("DB connection opened")
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/add", addWorkout)

	log.Panic(http.ListenAndServe(":5000", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)

	http.ServeFile(w, r, os.Getenv("HOME_URL"))
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

		http.Redirect(w, r, "http://localhost:5000/", http.StatusFound)
		return
	}

	http.ServeFile(w, r, os.Getenv("ADD_URL"))
}
