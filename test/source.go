package test

import "strings"

func Split(s, sep string) []string {
	var result []string
	i := strings.Index(s, sep)
	for i > -1 {
		if len(s[:i]) > 0 {
			result = append(result, s[:i])
		}
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	if len(s) > 0 {
		result = append(result, s)
	}
	return result
}
