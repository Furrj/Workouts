package main

import (
	"fmt"
	"github.com/Furrj/Workouts/src/internal/routeHandler"
	"log"
	"net/http"
	"os"
)

func main() {
	HtmlDir := os.Getenv("HTML_DIR")
	DataDir := os.Getenv("DATA_DIR")
	LogsDir := os.Getenv("LOGS_DIR")
	if HtmlDir == "" || DataDir == "" || LogsDir == "" {
		fmt.Println("Could not find env vars")
		os.Exit(1)
	}
	rh := routeHandler.NewRouteHandler(HtmlDir, DataDir, LogsDir)

	http.HandleFunc("/", rh.Home)
	http.HandleFunc("/add", rh.AddWorkout)
	http.HandleFunc("/api/post", rh.PostWorkout)

	log.Panic(http.ListenAndServe(":5000", nil))
}
