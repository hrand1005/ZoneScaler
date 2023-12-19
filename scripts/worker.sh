#!/bin/bash

export COORDINATOR_HOST="localhost"
export COORDINATOR_PORT="8080"

go run cmd/worker/main.go cmd/worker/config.json