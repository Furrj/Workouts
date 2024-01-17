#!/usr/bin/env bash

rm -rf build &&
mkdir build &&
cp src/html/* build &&
go build -o build/server src/server.go &&
cd build &&
export HOME_URL=home.html &&
export ADD_URL=add_workout.html &&
./server