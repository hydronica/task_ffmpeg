FROM task_ffmpeg/build:stage as build


FROM jeremiahz/ffmpeg:v4.2.1
RUN  mkdir /configs
COPY deploy/configs/tffmpeg.toml /configs/tffmpeg.toml
COPY --from=build /usr/bin/tffmpeg /usr/bin/tffmpeg

CMD [ "sh","-c","tffmpeg -config /configs/tffmpeg.toml" ]

