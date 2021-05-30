package restringer

import "taylz.io/types"

// RestringerLenExact returns a Restringer which cuts message prefix (...) or right-pads as necessary
func LenExact(size int) types.Restringer {
	return Func(func(str string) string { return StringLenExact(str, size) })
}

// StringLenExact returns a string of set size, elided (from the left) if longer, or right-padded if shorter
func StringLenExact(str string, size int) string {
	lenstr := len(str)
	if lenstr == size {
		return str
	} else if size < 1 {
		return ""
	}
	lendif := lenstr - size
	buf := make([]byte, size)
	var i, j int
	if lendif > 0 {
		buf[0], buf[1], buf[2] = '.', '.', '.'
		i = 3
		j = lendif + i
	}
	for i < size && j < lenstr {
		buf[i] = str[j]
		i++
		j++
	}
	for ; i < size; i++ {
		buf[i] = ' '
	}
	return string(buf)
}

// RestringerLenMin returns a Restringer which right-pads short messages
func LenMin(size int) types.Restringer {
	return Func(func(str string) string { return StringLenMin(str, size) })
}

// StringLenMin returns strings of minimum size, right-padded if shorter
func StringLenMin(str string, size int) string {
	lenstr := len(str)
	if lenstr >= size {
		return str
	} else if size < 1 {
		return ""
	}
	buf := make([]byte, size)
	for i := 0; i < lenstr; i++ {
		buf[i] = str[i]
	}
	for i := lenstr; i < size; i++ {
		buf[i] = ' '
	}
	return string(buf)
}
