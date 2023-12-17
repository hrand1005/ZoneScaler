#!/bin/bash

go build -o crd cmd/coordinator/main.go
./crd cmd/coordinator/config.json