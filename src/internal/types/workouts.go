package types

type Workout struct {
	WorkoutID uint64 `json:"workout_id"`
	Timestamp uint64 `json:"timestamp"`
	Sets      []Set  `json:"sets"`
}
type Set struct {
	WorkoutID uint64 `json:"workout_id"`
	ReqSet
}
