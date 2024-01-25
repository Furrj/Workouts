#!/usr/bin/env bash

source set_vars_dev.env &&
cd "$PROJECT_ROOT" &&
rm -rf "$BUILD_DIR" &&
mkdir "$BUILD_DIR" &&
mkdir "$HTML_DIR" &&
mkdir "$DATA_DIR" &&
mkdir "$LOGS_DIR" &&
echo "workoutID,setID,timestamp,text,name,reps,weights" >> "$SETS_CSV_URL" &&
echo "0" >> "$META_CSV_URL" &&go build -o "$BUILD_DIR"/server.exe src/server.go
cp src/html/* "$HTML_DIR" &&
rm "$VIEW_WORKOUT_HTML_URL" &&
go build -o "$BUILD_DIR"/server.exe src/server.go &&
cd build &&
./server.exe