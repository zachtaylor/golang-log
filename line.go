package log

import "time"

// Line is log data
type Line struct {
	Level  Level
	Fields Fields
	Source Source
	Time   time.Time
	Args   []interface{}
}

func NewLine(lvl Level, fields Fields, src Source, time time.Time, args []interface{}) Line {
	return Line{
		Level:  lvl,
		Fields: fields,
		Source: src,
		Time:   time,
		Args:   args,
	}
}
