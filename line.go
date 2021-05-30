package log

import "taylz.io/types"

// Line is log data
type Line struct {
	Level  Level
	Fields Fields
	Source types.Source
	Time   types.Time
	Args   []interface{}
}

func NewLine(lvl Level, fields Fields, src types.Source, time types.Time, args []interface{}) Line {
	return Line{
		Level:  lvl,
		Fields: fields,
		Source: src,
		Time:   time,
		Args:   args,
	}
}
