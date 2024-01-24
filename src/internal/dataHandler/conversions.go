package dataHandler

import (
	"fmt"
	"github.com/Furrj/Workouts/src/internal/types"
	"strconv"
)

func ConvertToSets(records [][]string) ([]types.Set, error) {
	var sets []types.Set

	for i, l := range records {
		if i == 0 {
			continue
		}

		set, err := convertToSet(l)
		if err != nil {
			return sets, err
		}
		fmt.Printf("%+v\n", set)
		sets = append(sets, set)
	}

	return sets, nil
}

func convertToSet(record []string) (types.Set, error) {
	var set types.Set

	workoutID, err := strconv.ParseUint(record[0], 10, 64)
	if err != nil {
		return set, err
	}
	set.WorkoutID = workoutID

	setID, err := strconv.ParseUint(record[1], 10, 64)
	if err != nil {
		return set, err
	}
	set.SetID = uint8(setID)

	timestamp, err := strconv.ParseUint(record[2], 10, 64)
	if err != nil {
		return set, err
	}
	set.Timestamp = timestamp

	set.Text = record[3]
	set.Name = record[4]
	set.Reps = record[5]
	set.Weights = record[6]

	return set, nil
}

func ConvertToWorkouts(sets []types.Set) []types.Workout {
	var workouts []types.Workout

	for _, s := range sets {
		if len(workouts) < int(s.WorkoutID) {
			workout := types.Workout{
				WorkoutID: s.WorkoutID,
				Timestamp: s.Timestamp,
				Sets: []types.Set{
					s,
				},
			}
			workouts = append(workouts, workout)
			continue
		}

		workouts[s.WorkoutID-1].Sets = append(workouts[s.WorkoutID-1].Sets, s)
	}

	return workouts
}
