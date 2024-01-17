package main

import (
	"fmt"
	"github.com/Furrj/Workouts/src/internal/types"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
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

		newWorkout := types.Workout{
			WorkoutID: 0,
			Timestamp: uint64(time.Now().Unix()),
		}

		var setCount uint8 = 1
		for key, value := range r.PostForm {
			fmt.Printf("%s: %s\n", key, value[0])

			newSet := types.ReqSet{
				SetID: setCount,
				Text:  value[0],
			}
			setCount++

			fmt.Printf("WorkoutID: %+v\n", newWorkout)
			fmt.Printf("Set: %+v\n", newSet)
		}

		http.Redirect(w, r, "http://localhost:5000/", http.StatusFound)
		return
	}

	http.ServeFile(w, r, os.Getenv("ADD_URL"))
}
