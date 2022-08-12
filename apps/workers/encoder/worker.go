package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jbsmith7741/go-tools/appenderr"
	"github.com/jbsmith7741/uri"
	"github.com/pcelvng/task"
	"github.com/pcelvng/task-tools/file"
)

type worker struct {
	task.Meta
	Source      string `uri:"origin"`
	Encoding    string `uri:"encoding"`
	Audio       string `uri:"audio"`
	CRF         string `uri:"crf"`
	Destination string `uri:"dest"`
	options
}

func (o *options) newWorker(info string) task.Worker {
	w := &worker{
		Meta:     task.NewMeta(),
		Encoding: "libx265",
		Audio:    "copy",
		CRF:      "23",
		options:  *o,
	}
	uri.Unmarshal(info, w)

	w.WorkingDir = strings.TrimRight(w.WorkingDir, "/") + "/"
	if w.Destination == "" {
		w.Destination = w.Source
	}
	return w

}

func (w *worker) DoTask(ctx context.Context) (result task.Result, msg string) {
	// copy file to local drive
	reader, err := file.NewReader(w.Source, w.File)
	if err != nil {
		return task.Failf("failed to copy %s: %s", w.Source, err)
	}
	inSts, err := file.Stat(w.Source, w.File)
	if err != nil {
		return task.Failed(err)
	}

	in, out := getLocalFiles(w.Source, w.Destination)
	inFile := w.WorkingDir + in
	outFile := w.WorkingDir + out

	writer, err := file.NewWriter(inFile, w.File)
	if err != nil {
		return task.Failf("writer init failed for %s: %s", inFile, err)
	}

	start := time.Now()
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return task.Failf("fail to read %s: %s", w.Source, err)
	}
	writer.Write(b)

	writer.Close()
	log.Println(time.Since(start).String())
	w.SetMeta("copyTime", time.Since(start).String())

	defer func() { // cleanup temp file
		os.Remove(inFile)
	}()

	// process the file
	encodeTime := time.Now()
	cmdStr := []string{"-hide_banner",
		"-i", inFile,
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
		if len(b) > 1000 {
			b = b[:1000]
		}
		fmt.Println(string(b))
		fmt.Println("error running ffmpeg", err)
		return task.Failed(err)
	}
	w.SetMeta("procTime", time.Since(encodeTime).String())

	defer func() {
		os.Remove(outFile)
	}()

	// copy processed file to final destination
	start = time.Now()
	reader.Close()
	errs := appenderr.New()
	reader, err = file.NewReader(outFile, nil)
	errs.Add(err)
	writer, err = file.NewWriter(w.Destination, w.File)
	errs.Add(err)

	b, err = ioutil.ReadAll(reader)
	errs.Add(err)
	if err := errs.ErrOrNil(); err != nil {
		task.Failf("unable to copy processed file %s: %s", outFile, err)
	}

	errs = appenderr.New()
	writer.Write(b)
	errs.Add(writer.Close())
	copyBack := time.Since(start)
	w.SetMeta("copyBack", copyBack.String())

	outSts := writer.Stats()
	diff := 100 * float64(inSts.Size-outSts.Size) / float64(inSts.Size)

	return task.Completed("media info saved %s Original: %v New: %v Reduction: %.1f%%",
		w.Destination, humanize.Bytes(uint64(inSts.Size)), humanize.Bytes(uint64(outSts.Size)), diff)

}

func getLocalFiles(source, dest string) (input string, output string) {
	_, fSource := path.Split(source)
	_, fDest := path.Split(dest)
	fSource = strings.ReplaceAll(fSource, " ", "_")
	fDest = strings.ReplaceAll(fDest, " ", "_")

	if fSource == fDest {
		ext := path.Ext(fSource)
		fDest = strings.ReplaceAll(fDest, ext, ".tmp"+ext)
	}

	return fSource, fDest
}
