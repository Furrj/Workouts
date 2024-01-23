package main

import (
	"github.com/Furrj/Workouts/src/internal/routeHandler"
	"github.com/Furrj/Workouts/src/internal/utils"
	"log"
	"net/http"
)

func main() {
	rh := routeHandler.NewRouteHandler(utils.InitEnvVars())

	http.HandleFunc("/", rh.Home)
	http.HandleFunc("/add", rh.AddWorkout)
	http.HandleFunc("/api/post", rh.PostWorkout)

	log.Panic(http.ListenAndServe(":5000", nil))
}
