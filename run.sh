#!/bin/bash
# Ensure that redis is running in a docker demon
# docker run --name some-redis -p 6379:6379 -d redis
# TODO: create a check to see if the redis is running

pushd gui
# run the qt python script
go run main.go
popd
