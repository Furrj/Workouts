package main

import (
	"encoding/json"
	"fmt"
	"github.com/Furrj/Workouts/src/internal/csv"
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
	var sets []types.ReqSet

	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading request: %+v\n", err)
		sendErr(w)
		return
	}
	err = json.Unmarshal(reqBytes, &sets)
	if err != nil {
		fmt.Printf("Error unmarshalling request: %+v\n", err)
		sendErr(w)
		return
	}

	if err := csv.Write(sets); err != nil {
		fmt.Printf("Error writing csv: %+v\n", err)
		sendErr(w)
		return
	}

	jsonRes, err := json.Marshal(types.ResBase{
		Status: "Success",
	})
	if err != nil {
		fmt.Printf("Error marshalling res: %+v\n", err)
		sendErr(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonRes); err != nil {
		fmt.Printf("Error sending response: %+v\n", err)
		sendErr(w)
		return
	}
}

func sendErr(w http.ResponseWriter) {
	e := types.ResBase{
		Status: "Error",
	}
	errMsg, err := json.Marshal(e)
	if err != nil {
		fmt.Printf("Error marshalling ereMsg: %+v\n", err)
	}

	w.WriteHeader(http.StatusConflict)
	_, err = w.Write(errMsg)
	if err != nil {
		fmt.Printf("Error sending errMsg: %+v\n", err)
	}
}
