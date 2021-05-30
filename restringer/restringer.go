package restringer

import "taylz.io/types"

// RestringerFunc is a func type that implements Restringer
type Func func(string) string

func (f Func) Restring(str string) string { return f(str) }

// RestringerMiddleware returns a Restringer from a sequence of Restringers, each result fed to the next
func Middleware(args ...types.Restringer) types.Restringer {
	lenargs := len(args)
	if lenargs < 1 {
		return nil
	} else if lenargs == 1 {
		return args[0]
	}
	return Func(func(str string) string {
		for i := 0; i < lenargs && str != ""; i++ {
			if next := args[i]; next != nil {
				str = next.Restring(str)
			}
		}
		return str
	})
}
