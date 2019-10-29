package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/pcelvng/task"
	"github.com/pcelvng/task-tools/file"
	"github.com/pcelvng/task/bus"
)

var json = jsoniter.ConfigFastest

type worker struct {
	inputFile string
	procPath  string
	encoding  string
	audio     string
	crf       string
	destPath  string
	fileTopic string
	writer    file.Writer // writer that writes the raw records
	producer  bus.Producer
}

func (o *options) newWorker(info string) task.Worker {
	var err error

	// parse info file (should be mediainfo json data)
	reader, err := file.NewReader(info, o.fileOptions)
	if err != nil {
		return task.InvalidWorker("could not create reader %s: %s", info, err.Error())
	}

	b, err := reader.ReadLine()
	if err != nil {
		return task.InvalidWorker("could not read info file %s", err.Error())
	}

	m := &MediaInfo{}
	err = json.Unmarshal(b, m)
	if err != nil {
		return task.InvalidWorker("could not unmarshal mediainfo %s: %s", info, err.Error())
	}

	videoTrack := false

	for _, t := range m.Media.Track {
		if t.Type == "Video" {
			videoTrack = true
			if t.EncodedLibraryName == "x265" && o.EncodeType == "libx265" {
				return task.InvalidWorker("encoding format is already x265")
			}
			if t.EncodedLibraryName == "x264" && o.EncodeType == "libx264" {
				return task.InvalidWorker("encoding format is already x264")
			}
		}
	}
	if !videoTrack && o.EncodeType != "none" {
		return task.InvalidWorker("no video track for format %s", o.EncodeType)
	}

	w := &worker{
		inputFile: info,
		procPath:  m.Media.Ref,
		encoding:  o.EncodeType,
		audio:     o.AudioType,
		crf:       o.CRF,
		producer:  o.producer,
		fileTopic: o.FileTopic,
		destPath:  fmt.Sprintf("%s/%s.%s.%s", o.DestFolder, filepath.Base(m.Media.Ref), o.EncodeType, o.NewFileType),
	}

	w.writer, err = file.NewWriter(w.destPath, o.fileOptions)
	if err != nil {
		return task.InvalidWorker("could not create writer %s", err.Error())
	}

	return w
}

func (w *worker) DoTask(ctx context.Context) (result task.Result, msg string) {
	fmt.Println("processing ffmpeg", w.inputFile)
	ctx, cancel := context.WithCancel(ctx)

	encodeTime := time.Now().Format("2006-01-02T00:00:00")
	defer cancel()
	cmdStr := fmt.Sprintf(`ffmpeg -hide_banner -i "%s" -map 0 -c:v %s -c:a %s -crf %s -map_metadata -1 -map_chapters -1 -metadata creation_time=%s "%s"`,
		w.procPath, w.encoding, w.audio, w.crf, encodeTime, w.destPath)
	fmt.Println("running command", cmdStr)
	cmd := exec.CommandContext(ctx, "sh", "-c", cmdStr)
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error running ffmpeg", err, "b:", string(b))
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
