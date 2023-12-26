package leetcode

import "strings"

func reverseWords1(s string) string {
	s = strings.Trim(s, " ")
	runes := []rune(s)
	result := ""
	i, n := 0, len(runes)
	for i < n {
		for ; i < n && runes[i] == ' '; i++ {
		}
		j := i
		for ; j < n && runes[j] != ' '; j++ {
		}

		result = string(runes[i:j]) + " " + result
		i = j + 1
	}

	return strings.TrimRight(result, " ")
}

func reverseWords2(s string) string {
	s = strings.Trim(s, " ")
	words := strings.Fields(s)

	for start, end := 0, len(words)-1; start < end; start, end = start+1, end-1 {
		words[start], words[end] = words[end], words[start]
	}

	return strings.Join(words, " ")
}
