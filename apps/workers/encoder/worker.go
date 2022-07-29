package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"
	"time"

	"github.com/pcelvng/task"
	"github.com/pcelvng/task-tools/file"
)

type worker struct {
	FilePath string `uri:"origin"`
	Encoding string `uri:"Encoding"`
	Audio    string `uri:"audio"`
	CRF      string `uri:"crf"`
	options

	procFile string
}

func (o *options) newWorker(info string) task.Worker {
	w := &worker{
		Encoding: "libx265",
		Audio:    "copy",
		CRF:      "27",
		options:  *o,
	}
	return w

}

func (w *worker) DoTask(ctx context.Context) (result task.Result, msg string) {
	// copy file to local drive
	r, err := file.NewReader(w.FilePath, w.File)
	if err != nil {
		return task.Failf("failed to copy %s: %s", w.FilePath, err)
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return task.Failf("fail to read %s: %s", w.FilePath, err)
	}
	_, f := path.Split(w.FilePath)
	w.procFile = w.WorkingDir + "/" + f
	writer, err := file.NewWriter(w.procFile, w.File)
	if err != nil {
		return task.Failf("writer init failed for %s: %s", w.procFile, err)
	}
	writer.Write(b)
	writer.Close()

	encodeTime := time.Now().Format("2006-01-02T00:00:00")
	cmdStr := fmt.Sprintf(`-hide_banner -i "%s" -map 0 -c:v %s -c:a %s -crf %s -map_metadata -1 -map_chapters -1 -metadata creation_time=%s "%s"`,
		w.procFile, w.Encoding, w.Audio, w.CRF, encodeTime, w.procFile+".tmp")
	fmt.Println("running command", cmdStr)
	cmd := exec.CommandContext(ctx, "ffmpeg", cmdStr)
	b, err = cmd.CombinedOutput()
	fmt.Println(b)
	if err != nil {
		fmt.Println("error running ffmpeg", err)
		return task.Failed(err)
	}

	return task.Completed("media info saved %s", w.procFile+".tmp")
}
