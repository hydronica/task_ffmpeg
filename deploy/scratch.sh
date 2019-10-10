#!/usr/bin/env bash

# used to drain topics to stdout so they can be reviewed
# probably need task logging for this
nsq_tail --topic=done --lookupd-http-address=nsqlookupd:4161

# This command will create an audio wave file, and then a video file merged with the audio file for testing
ffmpeg -f lavfi -i sine=frequency=1000:duration=10 test.wav
ffmpeg -f lavfi -i testsrc=duration=10:size=1920x1080:rate=60 -i test.wav testsrc.mkv
rm test.wav

# this will convert the file to a mp4 x264 file just for testing
ffmpeg -hide_banner -i testsrc.mkv -c:v libx265 -c:a copy -metadata title="testing" -metadata creation_time=2019-10-08T00:00:00 testsrc.265.mp4

# mediainfo --Output=JSON testsrc.mkv | jq

#nsq admin
http://localhost:4171


# for setting environment vars in windows on the command line this page might be helpful
https://github.com/docker/compose/issues/5089
https://user-images.githubusercontent.com/13176884/29216420-4455c7e0-7e74-11e7-9252-5311ee86d7e4.png

https://staxmanade.com/2016/05/how-to-get-environment-variables-passed-through-docker-compose-to-the-containers/