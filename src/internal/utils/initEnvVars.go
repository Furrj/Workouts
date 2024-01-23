package utils

import (
	"github.com/Furrj/Workouts/src/internal/types"
	"os"
)

func InitEnvVars() types.EnvVars {
	env := types.EnvVars{
		HtmlDir:           os.Getenv("HTML_DIR"),
		HomeHtmlUrl:       os.Getenv("HOME_HTML_URL"),
		AddWorkoutHtmlUrl: os.Getenv("ADD_WORKOUT_HTML_URL"),
		SetsCsvUrl:        os.Getenv("SETS_CSV_URL"),
		MetaCsvUrl:        os.Getenv("META_CSV_URL"),
	}

	return env
}
