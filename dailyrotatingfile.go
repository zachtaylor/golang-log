package log

import (
	"io"
	"os"
	"time"
)

// DailyRotatingFile creates a io.WriteCloser that rotates backing file, named for each day, for a drfectory
func DailyRotatingFile(path string) io.WriteCloser {
	r := &drf{
		Path: path,
		out:  make(chan []byte),
		done: make(chan bool),
	}
	go r.start()
	return r
}

type drf struct {
	Path string
	out  chan []byte
	done chan bool
	file io.WriteCloser
}

func (drf *drf) Write(bytes []byte) (int, error) {
	go drf.write(bytes)
	return len(bytes), nil
}

func (drf *drf) write(bytes []byte) { drf.out <- bytes }

func (drf *drf) Close() error {
	go drf.close()
	return nil
}

func (drf *drf) close() { close(drf.done) }

func (drf *drf) start() {
	t := time.Now()
	timer := time.NewTimer(eod(t))
	for {
		fileTitle := drf.Path + t.Format("2006_01_02")
		fileName := fileTitle + ".log"
		file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		drf.file = file
		err := drf.wait(timer)
		file.Close()
		if err == io.EOF {
			return
		}
		t = time.Now()
		timer.Reset(eod(t))
	}
}

// wait ends with the timer (nil) or when w.done is closed (io.EOF)
func (drf *drf) wait(timer *time.Timer) error {
	for {
		select {
		case <-timer.C:
			return nil
		case msg := <-drf.out:
			if _, err := drf.file.Write(msg); err != nil {
				return err
			}
		case <-drf.done:
			return io.EOF
		}
	}
}

func eod(t time.Time) time.Duration {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1).Sub(t)
}
