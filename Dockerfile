# this is the main build for task_ffmpeg binaries
# tag = task_ffmpeg/build:stage

FROM golang:1.13 as build
RUN go get -u -d github.com/pcelvng/task
RUN go get -u -d github.com/pcelvng/task-tools/file
RUN go get -u -d github.com/pcelvng/task/bus
RUN go get -u -d github.com/go-sql-driver/mysql
RUN go get -u -d github.com/lib/pq
RUN go get -u -d github.com/pelletier/go-toml
RUN go get -u -d gopkg.in/BurntSushi/toml.v0
RUN go get -u -d github.com/json-iterator/go
COPY / /go/src/github.com/hydronica/task_ffmpeg
WORKDIR /go/src/github.com/hydronica/task_ffmpeg
RUN CGO_ENABLED=0 GOOS=linux make

FROM alpine:3.8
RUN apk add curl
COPY --from=build /go/src/github.com/hydronica/task_ffmpeg/build/ /usr/bin/
RUN chmod +x /usr/bin/


