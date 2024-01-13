#!/usr/bin/env bash

rm -rf build &&
mkdir build &&
cp src/resources/* build &&
go build -o build/server src/server.go &&
cd build &&
export HOME_URL=home.html &&
export ADD_URL=add_workout.html &&
export URL=postgres://postgres:password@172.17.0.2:5432/workouts &&
./server