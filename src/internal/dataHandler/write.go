package dataHandler

import (
	"encoding/csv"
	"fmt"
	"github.com/Furrj/Workouts/src/internal/types"
	"log"
	"os"
	"strconv"
)

func WriteSetsCSV(dataDir string, sets []types.ReqSet) error {
	setsCSVURL := fmt.Sprintf("%s/sets.csv", dataDir)
	var workout [][]string
	fmt.Printf("ReqSets: %+v\n", sets)
	// TODO: workoutID
	var workoutID uint64 = 1

	for _, s := range sets {
		data := []string{strconv.FormatUint(workoutID, 10), strconv.FormatUint(uint64(s.SetID), 10), strconv.FormatUint(s.Timestamp, 10), s.Name, s.Text, s.Reps, s.Weights}

		workout = append(workout, data)
	}

	file, err := os.OpenFile(setsCSVURL, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panicf("Error opening dataHandler file writer: %+v\n", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panicf("Error closing dataHandler file writer: %+v\n", err)
		}
	}(file)

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, r := range workout {
		err = w.Write(r)
		if err != nil {
			log.Panicf("Error writing dataHandler file row: %+v\n", err)
		}
	}

	return nil
}
