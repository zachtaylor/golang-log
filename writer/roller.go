package writer

import (
	"io"
	"os"
	"time"
)

// NewRoller creates a io.WriteCloser that rotates backing file, named for each day, for a directory
func NewRoller(path string) io.WriteCloser {
	r := &roller{
		Path: path,
		out:  make(chan []byte),
		done: make(chan bool),
	}
	go r.start()
	return r
}

type roller struct {
	Path string
	out  chan []byte
	done chan bool
	file io.WriteCloser
}

func (r *roller) Write(bytes []byte) (int, error) {
	go r.write(bytes)
	return len(bytes), nil
}

func (r *roller) write(bytes []byte) { r.out <- bytes }

func (r *roller) Close() error {
	go r.close()
	return nil
}

func (r *roller) close() { close(r.done) }

func (r *roller) start() {
	t := time.Now()
	timer := time.NewTimer(eod(t))
	for {
		fileTitle := r.Path + t.Format("2006_01_02")
		fileName := fileTitle + ".log"
		file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		r.file = file
		err := r.wait(timer)
		file.Close()
		if err == io.EOF {
			return
		}
		t = time.Now()
		timer.Reset(eod(t))
	}
}

// wait ends with the timer (nil) or when w.done is closed (io.EOF)
func (r *roller) wait(ticker *time.Timer) error {
	for {
		select {
		case <-ticker.C:
			return nil
		case msg := <-r.out:
			if _, err := r.file.Write(msg); err != nil {
				return err
			}
		case <-r.done:
			return io.EOF
		}
	}
}

func eod(t time.Time) time.Duration {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1).Sub(t)
}
