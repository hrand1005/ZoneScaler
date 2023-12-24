#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <image-name>"
    exit 1
fi

IMAGE_NAME="$1"

running_containers=$(docker ps -q --filter "ancestor=$IMAGE_NAME")
if [ -n "$running_containers" ]; then
    echo "Removing running containers for image: $IMAGE_NAME"
    docker rm -f $running_containers
else
    echo "No running containers found for image: $IMAGE_NAME"
fi

stopped_containers=$(docker ps -aq --filter "ancestor=$IMAGE_NAME")
if [ -n "$stopped_containers" ]; then
    echo "Removing stopped containers for image: $IMAGE_NAME"
    docker rm $stopped_containers
else
    echo "No stopped containers found for image: $IMAGE_NAME"
fi
