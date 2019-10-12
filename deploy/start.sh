#!/bin/bash

docker-compose -f "docker/docker-compose-nsq.yml" up -d

# make sure the topics exist in nsq
docker run -it --network nsq_network_default task_ffmpeg/build:stage /bin/sh -c \
  "curl -X POST 'nsqd:4151/topic/create?topic=mediainfo'; \
   curl -X POST 'nsqd:4151/topic/create?topic=done';  \
   curl -X POST 'nsqd:4151/topic/create?topic=files';  \
   curl -X POST 'nsqd:4151/topic/create?topic=ffmpeg';"

printf "be sure to set the .env with the correct shared media locations"
docker-compose -f "docker/docker-compose-stage.yml" up -d

docker-compose -f "docker/docker-compose-stage.yml" logs -f
