package log

import (
	"io"
	"os"
)

type Liner interface {
	Line(Line)
	io.Closer
}

// // LinerFunc is a func type that implements Liner
// type LinerFunc func(Line)

// func (f LinerFunc) Line(line Line) { f(line) }

type routineLiner struct {
	in    chan Line
	liner Liner
}

func RoutineLiner(buf int, l Liner) Liner {
	rl := &routineLiner{
		in:    make(chan Line, buf),
		liner: l,
	}
	go rl.start()
	return rl
}
func (rl *routineLiner) start() {
	for {
		line, ok := <-rl.in
		if !ok {
			return
		}
		rl.liner.Line(line)
	}
}

func (rw *routineLiner) Line(line Line) {
	if rw.in != nil {
		rw.in <- line
	}
}

func (rw *routineLiner) Close() error {
	close(rw.in)
	rw.in = nil
	return rw.liner.Close()
}

type ioLiner struct {
	fmt Formatter
	out io.WriteCloser
}

// IOLiner returns a Liner that is backed by Formater and io.Writer
func IOLiner(f Formatter, w io.WriteCloser) Liner {
	return ioLiner{f, w}
}

func (w ioLiner) Line(line Line) { w.out.Write(w.fmt.Format(line)) }

func (w ioLiner) Close() error { return w.out.Close() }

type levelLiner struct {
	level Level
	liner Liner
}

// LevelLiner wraps a Liner, and limits the minimum level of lines
func LevelLiner(level Level, liner Liner) Liner {
	return levelLiner{level, liner}
}

func (l levelLiner) Line(line Line) {
	if line.Level >= l.level {
		l.liner.Line(line)
	}
}

func (l levelLiner) Close() error { return l.liner.Close() }

// MultiLiner is a []Liner that implements Liner, calling each internal Liner, for fan-out
type MultiLiner []Liner

func (m MultiLiner) Line(line Line) {
	for _, w := range m {
		w.Line(line)
	}
}

func (m MultiLiner) Close() (error error) {
	for _, w := range m {
		if err := w.Close(); err != nil {
			error = err
		}
	}
	return
}

// StdOutLiner returns a Liner that writes to os.Stdout
func StdOutLiner(fmt Formatter) Liner {
	return stdOutLiner{
		Formatter: fmt,
	}
}

// StdOutLiner is a Liner that Writes lines to os.Stdout synchronously
type stdOutLiner struct{ Formatter Formatter }

func (w stdOutLiner) Line(line Line) { os.Stdout.Write(w.Formatter.Format(line)) }

func (w stdOutLiner) Close() error { return os.Stdout.Close() }
