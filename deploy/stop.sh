#!/bin/bash

DOCKER_PATH="$(pwd)"

HOST_MEDIA=/workspace/mediafiles/  \
HOST_INFO=/workspace/mediainfo/  \
docker-compose -f "$DOCKER_PATH/docker/docker-compose-stage.yml" down