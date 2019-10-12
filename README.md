## task_ffmpeg

a task system for re-encoding files based on media-info rules (such as h264 to h265)

### Building and running the system

1. First run the deploy/build.sh or deploy/build.bat (for windows)
   - you must have docker (desktop) running for these
2. You'll need to update the .env file with the location of your media files
   - the HOST_MEDIA value is the local path where your media files will be watched, and re-encoded
   - the HOST_INFO is where the JSON info about the media will be stored and read for tffmpeg to re-encode
3. Next you should run the stop.sh or stop.bat files to stop any running instances
4. Run the start.sh or start.bat to start up the docker containers.

<hr>

### This is a work in progress, it's not perfect yet some issues I have found so far

- The filewatcher will pickup all files, might be good to filter what's watched and what isn't.
- The encoder will encode any file type in the HOST_MEDIA folder, as long as it has a video stream
  - it would be nice to have better rules around how things are encoded
  - maybe if a wav file is placed in the folder it should be encoded to an mp3
- The tffmpeg app doesn't give output or status updates as files are encoded, this leaves the user in the dark about how much time it's been running, or any errors the encoder might be getting as it's running.
  - There is a setting in ffmpeg to send progress to a http endpoint that could be read to return status info about the encoding progress. Calculations could be done on this info to determine how fast the encoding is running, (what frame is being processed) and how far into the file the process has reached.
  - There isn't a way to stop the current encoding process
    - there is a context in tffmpeg when running ffmpeg, but the cancel isn't being used, so right now the only way to stop the encoding is to stop the tffmpeg app.