#!/bin/bash

go build -o wrk cmd/worker/main.go
./wrk cmd/worker/config.json