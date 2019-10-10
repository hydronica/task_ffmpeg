package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/pcelvng/task"
	"github.com/pcelvng/task-tools/file"
	"github.com/pcelvng/task/bus"
)

type worker struct {
	filePath  string
	destPath  string
	fileTopic string
	writer    file.Writer // writer that writes the raw records
	producer  bus.Producer
}

func (o *options) newWorker(info string) task.Worker {
	var err error

	w := &worker{
		filePath:  info,
		producer:  o.producer,
		fileTopic: o.FileTopic,
		destPath:  fmt.Sprintf("%s/%s.mediainfo.json", o.DestFolder, filepath.Base(info)),
	}

	w.writer, err = file.NewWriter(w.destPath, o.fileOptions)
	if err != nil {
		return task.InvalidWorker("could not create writer %s", err.Error())
	}

	return w
}

func (w *worker) DoTask(ctx context.Context) (result task.Result, msg string) {
	fmt.Println("processing task", w.filePath)
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	cmdStr := fmt.Sprintf("mediainfo --Output=JSON %s | jq -c '.' > %s", w.filePath, w.destPath)
	err := exec.CommandContext(ctx, "sh", "-c", cmdStr).Run()
	if err != nil {
		fmt.Println("error running mediainfo", err)
		return task.Failed(err)
	}

	// send stats file info
	stats := w.writer.Stats()
	err = w.producer.Send(w.fileTopic, stats.JSONBytes())
	if err != nil {
		log.Println(err.Error())
	}

	return task.Completed("media info saved %s", w.destPath)
}
