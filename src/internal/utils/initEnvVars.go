package utils

import (
	"fmt"
	"github.com/Furrj/Workouts/src/internal/types"
	"os"
)

func InitEnvVars() types.EnvVars {
	env := types.EnvVars{
		HtmlDir:            getVar("HTML_DIR"),
		HomeHtmlUrl:        getVar("HOME_HTML_URL"),
		AddWorkoutHtmlUrl:  getVar("ADD_WORKOUT_HTML_URL"),
		ViewWorkoutHtmlUrl: getVar("VIEW_WORKOUT_HTML_URL"),
		SetsCsvUrl:         getVar("SETS_CSV_URL"),
		MetaCsvUrl:         getVar("META_CSV_URL"),
	}

	return env
}

func getVar(name string) string {
	key := os.Getenv(name)
	if key == "" {
		fmt.Printf("could not find env var %s\n", name)
		os.Exit(2)
	}
	return key
}
