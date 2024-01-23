package dataHandler

import (
	"github.com/Furrj/Workouts/src/internal/types"
	"os"
	"strconv"
	"strings"
)

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
