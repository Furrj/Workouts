package main

import (
	"encoding/json"
	"fmt"
	"github.com/Furrj/Workouts/src/internal/types"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/add", addWorkout)
	http.HandleFunc("/api/post", postWorkout)

	log.Panic(http.ListenAndServe(":5000", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)

	http.ServeFile(w, r, os.Getenv("HOME_URL"))
}

func addWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)

	http.ServeFile(w, r, os.Getenv("ADD_URL"))
}

func postWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /api/post")
	var req []types.ReqSet

	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading request: %+v\n", err)
		return
	}
	err = json.Unmarshal(reqBytes, &req)
	if err != nil {
		fmt.Printf("Error unmarshalling request: %+v\n", err)
		return
	}
	fmt.Printf("%+v\n", req)

	res := types.ResBase{
		Status: "Success",
	}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error marshalling res: %+v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonRes); err != nil {
		fmt.Printf("Error sending response: %+v\n", err)
		return
	}
}
