package types

type ResWorkout struct {
	WorkoutID uint64   `json:"workout_id"`
	Timestamp uint64   `json:"timestamp"`
	Sets      []ResSet `json:"sets"`
}

type ResSet struct {
	SetID   uint8  `json:"set_id"`
	Text    string `json:"text"`
	Name    string `json:"name"`
	Reps    string `json:"reps"`
	Weights string `json:"weights"`
}

type ResBase struct {
	Status string `json:"status"`
}
