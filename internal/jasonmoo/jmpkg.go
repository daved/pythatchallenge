package jmpkg

import "strings"

func Parse(input []string) []string {
	var in bool
	return strings.FieldsFunc(strings.Join(input, " "), func(c rune) bool {
		switch {
		case in && c == '"':
			in = false
			return true
		case c == '"':
			in = true
			return true
		case !in && c == ' ':
			return true
		}
		return false
	})
}
