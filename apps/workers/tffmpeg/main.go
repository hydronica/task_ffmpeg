package main

import (
	"github.com/pcelvng/task-tools/bootstrap"
	"github.com/pcelvng/task-tools/file"
	"github.com/pcelvng/task/bus"
	"github.com/zJeremiah/task_ffmpeg"
)

const (
	taskType    = "task_ffmpeg"
	description = `app will use the mediainfo json to determine how to process an encoding with ffmpeg
Example task:
{"type":"ffmpeg", "info":"/path/to/mediainfo/file.json"}`
)

type options struct {
	DestFolder     string `toml:"dest_folder" comment:"destination folder to store mediainfo json files"`
	FileTopic      string `toml:"file_topic" comment:"topic where account file stats are sent"`
	StatusPort     int    `toml:"status_port" comment:"port for HTTP status request"`
	RemoveOriginal bool   `toml:"rm_original" comment:"remove the origial media file after re-encoding"`
	// ffmpeg encoding options (this can be expanded to allow for hundreds of options)
	NewFileType string `toml:"new_file_type" comment:"new file extension, default is mp4"`
	EncodeType  string `toml:"encode_type" comment:"default is libx265"`
	AudioType   string `toml:"audio_type" comment:"default is copy (from source)"`
	Title       string `toml:"title" comment:"this can be set in the query params, default is just the file name"`
	CRF         string `toml:"crf" comment:"Constant Rate Factor default is 27"`

	fileOptions *file.Options
	producer    bus.Producer
}

func main() {
	o := &options{
		FileTopic:      "files",
		DestFolder:     "nop://",
		fileOptions:    file.NewOptions(),
		NewFileType:    "mp4",
		EncodeType:     "libx265",
		AudioType:      "copy",
		CRF:            "27",
		RemoveOriginal: false,
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
