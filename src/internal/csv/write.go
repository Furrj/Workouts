package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/Furrj/Workouts/src/internal/types"
	"log"
	"os"
	"strconv"
)

func Write(sets []types.ReqSet) error {
	var workout [][]string
	fmt.Printf("ReqSets: %+v\n", sets)
	// TODO: workoutID
	var workoutID uint64 = 1

	for _, s := range sets {
		data := []string{strconv.FormatUint(workoutID, 10), strconv.FormatUint(uint64(s.SetID), 10), strconv.FormatUint(s.Timestamp, 10), s.Name, s.Text, s.Reps, s.Weights}

		workout = append(workout, data)
	}

	dataCSVURL := os.Getenv("DATA_CSV_URL")
	file, err := os.OpenFile(dataCSVURL, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panicf("Error opening csv file writer: %+v\n", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panicf("Error closing csv file writer: %+v\n", err)
		}
	}(file)

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, r := range workout {
		err = w.Write(r)
		if err != nil {
			log.Panicf("Error writing csv file row: %+v\n", err)
		}
	}

	return nil
}
