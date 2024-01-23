package routeHandler

import (
	"encoding/json"
	"fmt"
	"github.com/Furrj/Workouts/src/internal/dataHandler"
	"github.com/Furrj/Workouts/src/internal/types"
	"io"
	"net/http"
)

type RouteHandler struct {
	EnvVars types.EnvVars
}

func NewRouteHandler(e types.EnvVars) RouteHandler {
	rh := RouteHandler{
		EnvVars: e,
	}

	return rh
}

func (rh RouteHandler) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)

	http.ServeFile(w, r, rh.EnvVars.HomeHtmlUrl)
}

func (rh RouteHandler) AddWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)

	http.ServeFile(w, r, rh.EnvVars.AddWorkoutHtmlUrl)
}

func (rh RouteHandler) PostWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /api/post")
	var sets []types.ReqSet

	meta, err := dataHandler.ReadMeta(rh.EnvVars.MetaCsvUrl)
	if err != nil {
		fmt.Printf("Error retrieving metadata: %+v\n", err)
		sendErr(w)
		return
	}
	meta.WorkoutCount++

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

	if err := dataHandler.WriteSetsCSV(rh.EnvVars.SetsCsvUrl, meta, sets); err != nil {
		fmt.Printf("Error writing to %s: %+v\n", rh.EnvVars.SetsCsvUrl, err)
		sendErr(w)
		return
	}
	if err := dataHandler.WriteMetaCSV(rh.EnvVars.MetaCsvUrl, meta); err != nil {
		fmt.Printf("Error writing to %s: %+v\n", rh.EnvVars.MetaCsvUrl, err)
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
