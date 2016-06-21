#!/bin/bash

set -e

DOCKER_IMAGE_NAME="terraform-for-arukas-build"
DOCKER_CONTAINER_NAME="terraform-for-arukas-build-container"

if [[ $(docker ps -a | grep $DOCKER_CONTAINER_NAME) != "" ]]; then
  docker rm -f $DOCKER_CONTAINER_NAME 2>/dev/null
fi

docker build -t $DOCKER_IMAGE_NAME .

docker run --name $DOCKER_CONTAINER_NAME \
       -e ARUKAS_JSON_API_TOKEN \
       -e ARUKAS_JSON_API_SECRET \
       -e ARUKAS_JSON_API_URL \
       -e ARUKAS_DEBUG \
       -e TF_LOG \
       -e TESTARGS \
       $DOCKER_IMAGE_NAME make "$@"


if [[ "$@" == *"build"* ]]; then
  docker cp $DOCKER_CONTAINER_NAME:/go/src/github.com/yamamoto-febc/terraform-provider-arukas/bin ./
fi
docker rm -f $DOCKER_CONTAINER_NAME 2>/dev/null
