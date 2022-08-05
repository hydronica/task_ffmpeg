package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/jbsmith7741/uri"
	"github.com/pcelvng/task"
	"github.com/pcelvng/task-tools/file"
)

const (
	B  int64 = 1
	KB int64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

type worker struct {
	task.Meta
	FilePath string `uri:"origin"`
	Encoding string `uri:"Encoding"`
	Audio    string `uri:"audio"`
	CRF      string `uri:"crf"`
	options

	procFile string
}

func (o *options) newWorker(info string) task.Worker {
	w := &worker{
		Meta:     task.NewMeta(),
		Encoding: "libx265",
		Audio:    "copy",
		CRF:      "27",
		options:  *o,
	}
	uri.Unmarshal(info, w)
	return w

}

func (w *worker) DoTask(ctx context.Context) (result task.Result, msg string) {
	// copy file to local drive
	r, err := file.NewReader(w.FilePath, w.File)
	if err != nil {
		return task.Failf("failed to copy %s: %s", w.FilePath, err)
	}
	sts, err := file.Stat(w.FilePath, w.File)
	if err != nil {
		return task.Failed(err)
	}

	_, f := path.Split(w.FilePath)
	ext := path.Ext(f)
	w.procFile = strings.ReplaceAll(strings.TrimRight(w.WorkingDir, "/")+"/"+f, " ", "_")
	outFile := strings.ReplaceAll(w.procFile, ext, ".tmp"+ext)

	writer, err := file.NewWriter(w.procFile, w.File)
	if err != nil {
		return task.Failf("writer init failed for %s: %s", w.procFile, err)
	}

	start := time.Now()
	if sts.Size > GB {
		// write file piecewise
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			writer.Write(scanner.Bytes())
		}
	} else {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			return task.Failf("fail to read %s: %s", w.FilePath, err)
		}
		writer.Write(b)
	}
	writer.Close()

	w.SetMeta("copyTime", time.Since(start).String())

	encodeTime := time.Now()
	cmdStr := []string{"-hide_banner",
		"-i", w.procFile,
		"-map", "0",
		"-c:v", w.Encoding,
		"-c:a", w.Audio,
		"-crf", w.CRF,
		"-map_metadata", "-1",
		"-map_chapters", "-1",
		"-metadata", "creation_time=" + encodeTime.Format("2006-01-02T15:04:05"),
		outFile}
	fmt.Println("cmd: ffmpeg", strings.Join(cmdStr, " "))
	cmd := exec.CommandContext(ctx, "ffmpeg", cmdStr...)
	b, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(b))
		fmt.Println("error running ffmpeg", err)
		return task.Failed(err)
	}
	w.SetMeta("procTime", time.Since(encodeTime).String())

	return task.Completed("media info saved %s", w.procFile+".tmp")
}
