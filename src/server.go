package main

import (
	"fmt"
	"github.com/Furrj/Workouts/src/internal/dataHandler"
	"github.com/Furrj/Workouts/src/internal/routeHandler"
	"github.com/Furrj/Workouts/src/internal/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	rh := routeHandler.NewRouteHandler(utils.InitEnvVars())

	_, err := dataHandler.ReadMeta(rh.EnvVars.MetaCsvUrl)
	if err != nil {
		fmt.Printf("Error reading meta: %+v\n", err)
		os.Exit(3)
	}

	http.HandleFunc("/", rh.Home)
	http.HandleFunc("/add", rh.AddWorkout)
	http.HandleFunc("/api/post", rh.PostWorkout)

	log.Panic(http.ListenAndServe(":5000", nil))
}
