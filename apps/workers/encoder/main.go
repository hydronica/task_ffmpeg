package main

import (
	"log"
	"os/exec"

	"github.com/hydronica/task_ffmpeg"
	"github.com/jbsmith7741/go-tools/appenderr"
	"github.com/pcelvng/task-tools/bootstrap"
	"github.com/pcelvng/task-tools/file"
)

const (
	taskType    = "task_ffmpeg"
	description = `app will use the mediainfo json to determine how to process an encoding with ffmpeg
info string options 
- encoding: ffmpeg encoding; default: libx265
- audio: ffmpeg audio; default: copy 
- crf: default: 23. see https://trac.ffmpeg.org/wiki/Encode/H.264
- dest: destination for the resulting file. If the filename is blank if will override the existing file. 

Example task:
{"type":"ffmpeg", "info":"/path/to/mediainfo/file.mp4?dest=/path/to/media/file.done.mp4"}`
)

type options struct {
	WorkingDir string `toml:"working_dir" comment:"directory to store files for processing"`

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
	errs := appenderr.New()
	errs.Add(err)
	if o.WorkingDir == "" {
		errs.Addf("working_dir is required")
	}
	return errs.ErrOrNil()
}
