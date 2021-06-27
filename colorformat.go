package log

import (
	"fmt"
	"sort"
	"strings"
)

// ColorFormat is a Formatter designed for humans
type ColorFormat struct {
	Colors     ColorMap
	ColorTime  bool
	ColorSrc   bool
	ColorMsg   bool
	ColorField bool
	SrcFmt     SourceFormatter
	TimeFmt    TimeFormatter
}

func (fmt *ColorFormat) Format(line Line) []byte {
	var colorCode string
	if fmt.Colors != nil && (fmt.ColorTime || fmt.ColorSrc || fmt.ColorMsg) {
		colorCode = fmt.Colors[line.Level]
	}

	time := fmt.TimeFmt.FormatTime(line.Time)
	src := fmt.SrcFmt.FormatSource(line.Source)
	var sb strings.Builder
	var dirty bool

	if colorCode == "" {
		if len(time) > 0 {
			sb.WriteString(time)
			sb.WriteByte(' ')
		}
		sb.WriteByte(line.Level.ByteCode())
		dirty = true // leave it on now
		if len(src) > 0 {
			sb.WriteByte(' ')
			sb.WriteString(src)
		}
		for _, arg := range line.Args {
			sb.WriteByte(' ')
			fmt.writeValue(&sb, arg)
		}
	} else {
		var incolor bool

		if len(time) > 0 {
			if fmt.ColorTime {
				sb.WriteString(colorCode)
				incolor = true
			}
			sb.WriteString(time)
			dirty = true
		}

		if len(src) > 0 {
			if fmt.ColorSrc != incolor {
				if fmt.ColorSrc {
					sb.WriteString(colorCode)
					incolor = true
				} else {
					sb.WriteString(ColorOff)
					incolor = false
				}
			}
			if dirty {
				sb.WriteByte(' ')
			}
			sb.WriteString(src)
			dirty = true
		}

		if fmt.ColorMsg != incolor {
			if fmt.ColorMsg {
				sb.WriteString(colorCode)
				incolor = true
			} else {
				sb.WriteString(ColorOff)
				incolor = false
			}
		}

		for _, arg := range line.Args {
			if dirty {
				sb.WriteByte(' ')
			}
			fmt.writeValue(&sb, arg)
			dirty = true
		}

		if incolor {
			sb.WriteString(ColorOff)
			incolor = false
		}
	}

	if dirty {
		sb.WriteString("\t\t\t")
		dirty = false
	}

	for _, key := range fmt.fieldKeys(line.Fields) {
		if dirty {
			sb.WriteByte(' ')
		}
		if fmt.ColorField {
			sb.WriteString(colorCode)
		}
		sb.WriteString(key)
		if fmt.ColorField {
			sb.WriteString(ColorOff)
		}
		sb.WriteByte('=')
		fmt.writeValue(&sb, line.Fields[key])
		dirty = true
	}

	sb.WriteByte('\n')

	return []byte(sb.String())
}

func (*ColorFormat) fieldKeys(fields Fields) []string {
	keys, i := make([]string, len(fields)), 0
	for k := range fields {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys

}

// writeValue records a value representing any type
func (*ColorFormat) writeValue(sb *strings.Builder, arg interface{}) {
	if arg == nil {
		sb.WriteString("<nil>")
	} else if err, _ := arg.(error); err != nil {
		sb.WriteString(err.Error())
	} else if str, _ := arg.(string); len(str) > 0 {
		sb.WriteString(str)
	} else if stringer, _ := arg.(fmt.Stringer); stringer != nil {
		sb.WriteString(stringer.String())
	} else {
		fmt.Fprintf(sb, "%v", arg)
	}
}
