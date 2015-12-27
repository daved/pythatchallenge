package srtpkg

import "strings"

func Parse(in []string) []string {
	var result []string = []string{}
	var quoteStarted bool
	for _, item := range in {
		replaced := strings.Replace(item, `"`, "", -1)
		if quoteStarted {
			result[len(result)-1] += " " + replaced
		} else {
			result = append(result, replaced)
		}
		if strings.Contains(item, `"`) {
			quoteStarted = !quoteStarted
		}
	}
	return result
}
