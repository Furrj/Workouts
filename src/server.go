package main

import (
	"github.com/Furrj/Workouts/src/internal/routeHandler"
	"github.com/Furrj/Workouts/src/internal/utils"
	"log"
	"net/http"
)

func main() {
	rh := routeHandler.NewRouteHandler(utils.InitEnvVars())

	//records, err := dataHandler.ReadWorkouts(rh.EnvVars.SetsCsvUrl)
	//if err != nil {
	//	fmt.Printf("Error reading data.csv: %+v\n", err)
	//	os.Exit(1)
	//}
	//sets, err := dataHandler.ConvertToSets(records)
	//if err != nil {
	//	fmt.Printf("Error converting to sets: %+v\n", err)
	//}
	//workouts := dataHandler.ConvertToResWorkouts(sets)
	//for _, w := range workouts {
	//	fmt.Printf("%+v\n", w)
	//}

	http.HandleFunc("/", rh.Home)
	http.HandleFunc("/add", rh.AddWorkout)
	http.HandleFunc("/list", rh.ViewWorkout)
	http.HandleFunc("/api/post", rh.PostWorkout)

	log.Panic(http.ListenAndServe(":5000", nil))
}
