package log

// Restringer is a header for string manipulation
type Restringer interface {
	// Restring returns another string
	Restring(string) string
}

// RestringerFunc is a func type that implements Restringer
type RestringerFunc func(string) string

func (f RestringerFunc) Restring(str string) string { return f(str) }

// RestringerLenExact returns a Restringer which cuts message prefix (...) or right-pads as necessary
func RestringerLenExact(size int) Restringer {
	return RestringerFunc(func(str string) string { return FormatStringLenExact(str, size) })
}

// RestringerLenMin returns a Restringer which right-pads short messages
func RestringerLenMin(size int) Restringer {
	return RestringerFunc(func(str string) string { return FormatStringLenMin(str, size) })
}

// RestringerMiddleware returns a Restringer from a sequence of Restringers, each result fed to the next
func RestringerMiddleware(args ...Restringer) Restringer {
	lenargs := len(args)
	if lenargs < 1 {
		return nil
	} else if lenargs == 1 {
		return args[0]
	}
	return RestringerFunc(func(str string) string {
		for i := 0; i < lenargs && str != ""; i++ {
			if next := args[i]; next != nil {
				str = next.Restring(str)
			}
		}
		return str
	})
}

// RestringerCutPrefixes is Restringer that cuts prefixes from its list
type RestringerCutPrefixes []string

// NewRestringerCutPrefixes creates a RestringerCutPrefixes
func NewRestringerCutPrefixes() RestringerCutPrefixes { return RestringerCutPrefixes([]string{}) }

// Restring implements Restringer by removing string prefixes
func (list RestringerCutPrefixes) Restring(str string) string {
	strlen := len(str)
	for _, pre := range list {
		if len := len(pre); pre[len-1] != '/' {
			pre = pre + "/"
		}
		if len := len(pre); len >= strlen {
		} else if str[:len] != pre {
		} else {
			return str[len:]
		}
	}
	return str
}
