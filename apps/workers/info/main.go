package main

import (
	"github.com/hydronica/task_ffmpeg"
	"github.com/pcelvng/task-tools/bootstrap"
	"github.com/pcelvng/task-tools/file"
	"github.com/pcelvng/task/bus"
)

const (
	taskType    = "mediainfo"
	description = `media info will use the mediainfo application to get a json string of the
	file media information.
Example task:
{"type":"mediainfo", "info":"/path/to/media/file.mp4"}`
)

type options struct {
	DestFolder string `toml:"dest_folder" comment:"destination folder to store mediainfo json files"`
	FileTopic  string `toml:"file_topic" comment:"topic where account file stats are sent"`
	StatusPort int    `toml:"status_port" comment:"port for HTTP status request"`

	fileOptions *file.Options
	producer    bus.Producer
}

func main() {
	o := &options{
		FileTopic:   "files",
		DestFolder:  "nop://",
		fileOptions: file.NewOptions(),
	}

	app := bootstrap.NewWorkerApp(taskType, o.newWorker, o).
		Version(task_ffmpeg.GetVersion()).
		Description(description).
		FileOpts()

	app.Initialize()

	o.fileOptions = app.GetFileOpts()
	if o.FileTopic != "-" {
		o.producer = app.NewProducer()
	}

	app.Run()
}

func (o *options) Validate() error {

	return nil
}
