package dataHandler

import (
	"encoding/csv"
	"fmt"
	"github.com/Furrj/Workouts/src/internal/types"
	"log"
	"os"
	"strconv"
)

func WriteSetsCSV(setsCSVURL string, meta types.MetaData, sets []types.ReqSet) error {
	var workout [][]string

	for _, s := range sets {
		data := []string{strconv.FormatUint(meta.WorkoutCount, 10), strconv.FormatUint(uint64(s.SetID), 10), strconv.FormatUint(s.Timestamp, 10), s.Text, s.Name, s.Reps, s.Weights}

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

	fmt.Printf("Wrote sets: %+v\n", sets)
	return nil
}

func WriteMetaCSV(metaCsvUrl string, data types.MetaData) error {
	file, err := os.Create(metaCsvUrl)
	if err != nil {
		return err
	}

	if _, err := file.WriteString(data.ToString()); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	fmt.Printf("Wrote meta: %+v\n", data)
	return nil
}
