docker-compose -f "docker/docker-compose-nsq.yml" up -d

echo Be sure to set the .env file with the correct shared volume paths
docker-compose -f "docker/docker-compose-stage.yml" up -d

docker run -it --network nsq_network_default task_ffmpeg/build:stage /bin/sh -c "curl -X POST 'nsqd:4151/topic/create?topic=mediainfo';"
docker run -it --network nsq_network_default task_ffmpeg/build:stage /bin/sh -c "curl -X POST 'nsqd:4151/topic/create?topic=done';"
docker run -it --network nsq_network_default task_ffmpeg/build:stage /bin/sh -c "curl -X POST 'nsqd:4151/topic/create?topic=files';"
docker run -it --network nsq_network_default task_ffmpeg/build:stage /bin/sh -c "curl -X POST 'nsqd:4151/topic/create?topic=ffmpeg';"

docker-compose -f "docker/docker-compose-stage.yml" logs