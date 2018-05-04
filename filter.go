package main

import "strings"

type Filter struct {
	Suffix  string
	Contain string
}

func (f Filter) match(key string) bool {
	matched := true
	if f.Suffix != "" {
		if !strings.HasSuffix(key, f.Suffix) {
			matched = false
		}
	}
	if f.Contain != "" {
		idx := strings.Index(key, f.Contain)
		if idx == -1 {
			matched = false
		}
	}
	return matched
}
