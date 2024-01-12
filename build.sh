#!/bin/bash

rm -rf build &&
mkdir build &&
cp src/resources/* build &&
go build -o build/server src/server.go