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
	SetsCSVURL := os.Getenv("SETS_CSV_URL")
	MetaCSVURL := os.Getenv("META_CSV_URL")
	if HtmlDir == "" || SetsCSVURL == "" || MetaCSVURL == "" {
		fmt.Println("Could not find env vars")
		os.Exit(1)
	}
	rh := routeHandler.NewRouteHandler(HtmlDir)

	http.HandleFunc("/", rh.Home)
	http.HandleFunc("/add", rh.AddWorkout)
	http.HandleFunc("/api/post", rh.PostWorkout)

	log.Panic(http.ListenAndServe(":5000", nil))
}
