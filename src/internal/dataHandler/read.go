package dataHandler

import (
	"encoding/csv"
	"github.com/Furrj/Workouts/src/internal/types"
	"os"
	"strconv"
	"strings"
)

func ReadWorkouts(setsCsvUrl string) ([][]string, error) {
	var records [][]string

	file, err := os.Open(setsCsvUrl)
	if err != nil {
		return records, err
	}

	reader := csv.NewReader(file)
	records, err = reader.ReadAll()
	if err != nil {
		return records, err
	}

	if err := file.Close(); err != nil {
		return records, err
	}
	return records, nil
}

func ReadMeta(setsCsvUrl string) (types.MetaData, error) {
	var meta = types.MetaData{}

	contents, err := os.ReadFile(setsCsvUrl)
	if err != nil {
		return meta, err
	}

	split := strings.Split(strings.Trim(string(contents), "\n"), ",")
	meta.WorkoutCount, err = strconv.ParseUint(split[0], 10, 64)
	if err != nil {
		return meta, err
	}

	return meta, nil
}
