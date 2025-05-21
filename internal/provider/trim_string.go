package provider

import (
	"sort"
	"strings"
)

func TrimString(input string, length int, substrings []string, preserveSuffix bool) string {
	if length <= 0 {
		return ""
	}

	sort.Strings(substrings)

	substringsFound := true
	for substringsFound {
		if len(input) == 0 || len(input) < length {
			break
		}
		substringsFound = false
		for _, substring := range substrings {
			if len(input) == 0 || len(input) < length {
				break
			}
			if len(substring) == 0 {
				continue
			}
			if strings.HasPrefix(input, substring+"-") {
				substringsFound = true
				input = input[len(substring)+1:]
			}
			if strings.HasSuffix(input, "-"+substring) {
				substringsFound = true
				input = input[:len(input)-len(substring)-1]
			}
			if strings.Contains(input, "-"+substring+"-") {
				substringsFound = true
				input = strings.ReplaceAll(input, "-"+substring, "")
			}
		}
	}

	if len(input) > length {
		if preserveSuffix {
			lastDash := strings.LastIndex(input, "-")
			if lastDash == -1 {
				return input[:length]
			}
			suffix := input[lastDash:]
			trimmedInput := input[:(length - len(suffix))]
			trimmedInput = strings.TrimSuffix(trimmedInput, "-")

			return trimmedInput + suffix
		} else {
			return strings.TrimSuffix(input[:length], "-")
		}
	}

	return input
}
