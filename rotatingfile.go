package log

import (
	"io"
	"os"
	"time"
)

// RotatingFile creates a io.WriteCloser that rotates backing file, named for each day, for a directory
func RotatingFile(path string) io.WriteCloser {
	r := &rotatingFile{
		path: path,
		out:  make(chan []byte),
		done: make(chan bool),
	}
	go r.start()
	return r
}

func RotatingFileLiner(fmt Formatter, path string) Liner {
	return IOLiner(fmt, RotatingFile(path))
}

type rotatingFile struct {
	path string
	out  chan []byte
	done chan bool
	file io.WriteCloser
}

func (rf *rotatingFile) Write(bytes []byte) (int, error) {
	go rf.write(bytes)
	return len(bytes), nil
}

func (rf *rotatingFile) write(bytes []byte) { rf.out <- bytes }

func (rf *rotatingFile) Close() error {
	go rf.close()
	return nil
}

func (rf *rotatingFile) close() { close(rf.done) }

func (rf *rotatingFile) start() {
	t := time.Now()
	timer := time.NewTimer(eod(t))
	for {
		fileTitle := rf.path + t.Format("2006_01_02")
		fileName := fileTitle + ".log"
		file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		rf.file = file
		err := rf.wait(timer)
		file.Close()
		if err == io.EOF {
			return
		}
		t = time.Now()
		timer.Reset(eod(t))
	}
}

// wait ends with the timer (nil) or when w.done is closed (io.EOF)
func (rf *rotatingFile) wait(timer *time.Timer) error {
	for {
		select {
		case <-timer.C:
			return nil
		case msg := <-rf.out:
			if _, err := rf.file.Write(msg); err != nil {
				return err
			}
		case <-rf.done:
			return io.EOF
		}
	}
}

func eod(t time.Time) time.Duration {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1).Sub(t)
}
