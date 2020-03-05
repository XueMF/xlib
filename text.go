package xlib

import (
	"sort"
	"strconv"
)

// PadZero pad strings with char to indicated length, like 1 -> 0001
func PadZero(number int, digitLength int, padChar string) (numberStr string) {
	if padChar == "" {
		padChar = " "
	}
	numberStr = strconv.Itoa(number)
	for {
		if len(numberStr) >= digitLength {
			break
		}
		numberStr = padChar + numberStr
	}
	return numberStr
}

// ReverseString do reversing of string
func ReverseString(s string) string {
	chars := []rune(s)
	length := len(chars)
	for i := 0; i <= length/2; i++ {
		chars[i], chars[length-1-i] = chars[length-1-i], chars[i]
	}
	return string(chars)
}

// ReverseStringSlice do reversing of string slice
func ReverseStringSlice(s []string) {
	length := len(s)
	for i := 0; i <= length/2; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}

// GetSortedKey do reversing of the string slice
func GetSortedKey(m map[string]int, increase bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	if increase == false {
		ReverseStringSlice(keys)
	}
	return keys
}

// Levenshtein calculate
// costIns: Defines the cost of insertion.
// costRep: Defines the cost of replacement.
// costDel: Defines the cost of deletion.
func Levenshtein(str1, str2 string, costIns, costRep, costDel int) int {
	var maxLen = 255
	l1 := len(str1)
	l2 := len(str2)
	if l1 == 0 {
		return l2 * costIns
	}
	if l2 == 0 {
		return l1 * costDel
	}
	if l1 > maxLen || l2 > maxLen {
		return -1
	}

	tmp := make([]int, l2+1)
	p1 := make([]int, l2+1)
	p2 := make([]int, l2+1)
	var c0, c1, c2 int
	var i1, i2 int
	for i2 := 0; i2 <= l2; i2++ {
		p1[i2] = i2 * costIns
	}
	for i1 = 0; i1 < l1; i1++ {
		p2[0] = p1[0] + costDel
		for i2 = 0; i2 < l2; i2++ {
			if str1[i1] == str2[i2] {
				c0 = p1[i2]
			} else {
				c0 = p1[i2] + costRep
			}
			c1 = p1[i2+1] + costDel
			if c1 < c0 {
				c0 = c1
			}
			c2 = p2[i2] + costIns
			if c2 < c0 {
				c0 = c2
			}
			p2[i2+1] = c0
		}
		tmp = p1
		p1 = p2
		p2 = tmp
	}
	c0 = p1[l2]

	return c0
}

// SimilarText find similarity
func SimilarText(first, second string, percent *float64) int {
	var similarText func(string, string, int, int) int
	similarText = func(str1, str2 string, len1, len2 int) int {
		var sum, max int
		pos1, pos2 := 0, 0

		// Find the longest segment of the same section in two strings
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
					if l+1 > max {
						max = l + 1
						pos1 = i
						pos2 = j
					}
				}
			}
		}

		if sum = max; sum > 0 {
			if pos1 > 0 && pos2 > 0 {
				sum += similarText(str1, str2, pos1, pos2)
			}
			if (pos1+max < len1) && (pos2+max < len2) {
				s1 := []byte(str1)
				s2 := []byte(str2)
				sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
			}
		}

		return sum
	}

	l1, l2 := len(first), len(second)
	if l1+l2 == 0 {
		return 0
	}
	sim := similarText(first, second, l1, l2)
	if percent != nil {
		*percent = float64(sim*200) / float64(l1+l2)
	}
	return sim
}
