package lib

import (
	"bytes"
	"fmt"
	"strings"
)

// IntArrayDeepCopy ...
func IntArrayDeepCopy(nums []int) []int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	return temp
}

// StringArrayDeepCopy ...
func StringArrayDeepCopy(nums []string) []string {
	temp := make([]string, len(nums))
	copy(temp, nums)
	return temp
}

func SplitWithToken(s string, pre, suf rune) ([]string, error) {
	preS := string(pre)
	sufS := string(suf)
	if !strings.HasPrefix(s, preS) || !strings.HasSuffix(s, sufS) {
		fmt.Printf("%s %s %s %s %s\n", s, preS, sufS, string(s[0]), string(s[len(s)-1]))
		return nil, fmt.Errorf("invalid string(%s): pre(%v) and suf(%v)", s, pre, suf)
	}

	s = strings.Replace(s, " ", "", -1)
	s = s[1 : len(s)-1]
	if strings.Count(s, preS) == 0 && strings.Count(s, sufS) == 0 {
		return strings.Split(s, ","), nil
	} else if strings.Count(s, ",") == 0 {
		return []string{s}, nil
	}

	var result []string
	var tokenCount int
	var buf bytes.Buffer
	for k, v := range s {
		// TODO: 优化
		if k == 0 {
			buf.WriteRune(v)
		} else if k == len(s)-1 {
			buf.WriteRune(v)
			result = append(result, buf.String())
			buf.Reset()
		} else if tokenCount == 0 {
			if v == ',' {
				result = append(result, buf.String())
				buf.Reset()
			} else {
				buf.WriteRune(v)
			}
		} else {
			buf.WriteRune(v)
		}

		switch v {
		case pre:
			tokenCount++
		case suf:
			tokenCount--
		}
	}

	return result, nil
}
