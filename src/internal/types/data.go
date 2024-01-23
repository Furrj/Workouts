package types

import "fmt"

type MetaData struct {
	WorkoutCount uint64 `json:"workout_count"`
}

func (m MetaData) ToString() string {
	return fmt.Sprintf("%d", m.WorkoutCount)
}
