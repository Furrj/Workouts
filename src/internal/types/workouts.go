package types

type Workout struct {
	WorkoutID uint64 `json:"workout_id"`
	Timestamp uint64 `json:"timestamp"`
}

type Set struct {
	WorkoutID uint64 `json:"workout_id"`
	SetID     uint8  `json:"set_id"`
	Text      string `json:"text"`
	Name      string `json:"name"`
	Reps      string `json:"reps"`
	Weights   string `json:"weights"`
}
