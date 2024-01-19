#!/usr/bin/env bash

rm -rf build &&
mkdir build &&
cp src/html/* build &&
echo "workoutID,setID,timestamp,text,name,reps,weights" >> build/data.csv &&
go build -o build/server src/server.go &&
cd build &&
export HOME_URL=home.html &&
export ADD_URL=add_workout.html &&
./server