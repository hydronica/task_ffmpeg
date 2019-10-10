FROM task_ffmpeg/build:stage

RUN mkdir /configs
COPY deploy/configs/info.toml /configs/info.toml
RUN apk add mediainfo
RUN apk add jq

CMD [ "sh","-c","info -config /configs/info.toml" ]

