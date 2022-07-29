package main

import (
	"log"
	"os/exec"

	"github.com/hydronica/task_ffmpeg"
	"github.com/pcelvng/task-tools/bootstrap"
	"github.com/pcelvng/task-tools/file"
)

const (
	taskType    = "task_ffmpeg"
	description = `app will use the mediainfo json to determine how to process an encoding with ffmpeg
Example task:
{"type":"ffmpeg", "info":"/path/to/mediainfo/file.json"}`
)

type options struct {
	WorkingDir string `toml"working_dir" comment:"directory to store files for processing"`

	File *file.Options `toml:"File"`
}

func main() {
	log.SetFlags(log.Llongfile)

	o := &options{
		File: file.NewOptions(),
	}

	app := bootstrap.NewWorkerApp(taskType, o.newWorker, o).
		Version(task_ffmpeg.GetVersion()).
		Description(description)

	app.Initialize()

	app.Run()

}

func (o *options) Validate() error {
	_, err := exec.LookPath("ffmpeg")

	return err
}
