package routeHandler

import (
	"encoding/json"
	"fmt"
	"github.com/Furrj/Workouts/src/internal/csv"
	"github.com/Furrj/Workouts/src/internal/types"
	"io"
	"net/http"
)

type RouteHandler struct {
	HtmlDir string
	DataDir string
	LogsDir string
}

func NewRouteHandler(HtmlDir string, DataDir string, LogsDir string) RouteHandler {
	rh := RouteHandler{
		HtmlDir: HtmlDir,
		DataDir: DataDir,
		LogsDir: LogsDir,
	}

	return rh
}

func (rh RouteHandler) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)

	http.ServeFile(w, r, fmt.Sprintf("%s/home.html", rh.HtmlDir))
}

func (rh RouteHandler) AddWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)

	http.ServeFile(w, r, fmt.Sprintf("%s/add_workout.html", rh.HtmlDir))
}

func (rh RouteHandler) PostWorkout(w http.ResponseWriter, r *http.Request) {
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

	if err := csv.Write(fmt.Sprintf("%s/data.csv", rh.DataDir), sets); err != nil {
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
