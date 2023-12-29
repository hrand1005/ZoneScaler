#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <image-name>"
    exit 1
fi

IMAGE_NAME="$1"

RUNNING=$(docker ps -q --filter "ancestor=$IMAGE_NAME")
if [ -n "$RUNNING" ]; then
    echo "Removing running containers for image: $IMAGE_NAME"
    docker rm -f $RUNNING
else
    echo "No running containers found for image: $IMAGE_NAME"
fi

STOPPED=$(docker ps -aq --filter "ancestor=$IMAGE_NAME")
if [ -n "$STOPPED" ]; then
    echo "Removing stopped containers for image: $IMAGE_NAME"
    docker rm $STOPPED
else
    echo "No stopped containers found for image: $IMAGE_NAME"
fi
