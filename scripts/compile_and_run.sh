#!/usr/bin/env bash

cd scripts &&
source set_vars_dev.env &&
cd .. &&
go build -o "$BUILD_DIR"/server.exe src/server.go &&
cd build &&
./server.exe