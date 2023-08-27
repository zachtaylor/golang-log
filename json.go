package log

import "encoding/json"

func EncodingJsonFormatter(srcfmt SourceFormatter) Formatter {
	return FormatterFunc(func(l Line) []byte {
		obj := map[string]interface{}{}
		for k, v := range l.Fields {
			obj[k] = v
		}
		obj["lvl"] = l.Level.Lowercase()
		obj["time"] = l.Time
		if len := len(l.Args); len > 1 {
			obj["msg"] = l.Args
		} else if len > 0 {
			obj["msg"] = l.Args[0]
		}
		if srcfmt != nil {
			obj["src"] = srcfmt.FormatSource(l.Source)
		}
		data, _ := json.Marshal(obj)
		return data
	})
}

func JsonRotatingFileLiner(srcfmt SourceFormatter, path string) Liner {
	return RotatingFileLiner(EncodingJsonFormatter(srcfmt), path)
}
