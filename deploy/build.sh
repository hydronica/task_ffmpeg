#!/bin/bash

RED='\033[01;31m'
GREEN='\033[01;32m'
WHITE='\033[01;39m'
YELLOW='\033[01;33m'

# this script is used to build the go binaries, and build the docker images
printf $GREEN"building task ffmpeg service apps, building docker images\n"
printf $YELLOW"if images are not cached, this could take some time...\n"

# VERSION="testing"
VERSION="stage"
# VERSION=$(git describe --tags)

docker network create nsq_network_default

set -e

printf $GREEN"building repo apps for $VERSION: $WHITE"
# CGO_ENABLED=0 GOOS=linux make

docker build -t task_ffmpeg/build:stage .

build_docker_image () {
  printf $YELLOW"docker build $APPTYPE/$APPNAME/$VERSION: $WHITE"
  docker build -t "task_ffmpeg/$APPNAME:$VERSION" -f apps/"$APPTYPE/$APPNAME"/Dockerfile . 
}

APPTYPE=utils
APPNAME=filewatcher
build_docker_image

APPTYPE=taskmasters
APPNAME=files
build_docker_image

APPTYPE=workers
APPNAME=info
build_docker_image

APPNAME=tffmpeg
build_docker_image