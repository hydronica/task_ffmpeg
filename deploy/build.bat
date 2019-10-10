docker network create nsq_network_default

docker build -t task_ffmpeg/build:stage .

docker build -f apps\utils\filewatcher\Dockerfile -t task_ffmpeg/filewatcher:stage .
docker build -f apps\taskmasters\files\Dockerfile -t task_ffmpeg/files:stage .
docker build -f apps\workers\info\Dockerfile -t task_ffmpeg/info:stage .