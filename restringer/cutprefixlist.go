package restringer

// CutPrefixList is Restringer that cuts prefixes from its list
type CutPrefixList []string

// NewCutPrefixList creates a CutPrefixList
func NewCutPrefixList() CutPrefixList { return CutPrefixList([]string{}) }

// Restring implements Restringer by removing string prefixes
func (list CutPrefixList) Restring(str string) string {
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
