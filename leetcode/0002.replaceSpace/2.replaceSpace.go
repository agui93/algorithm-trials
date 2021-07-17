package leetcode

import "strings"

func replaceSpace(s string) string {
	//return strings.ReplaceAll(s, " ", "%20")

	var res strings.Builder

	for i := range s {
		if s[i] == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteByte(s[i])
		}
	}

	return res.String()
}
