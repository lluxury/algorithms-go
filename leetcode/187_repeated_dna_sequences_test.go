package leetcode

import (
	"sort"
	"testing"

	"github.com/Chyroc/algorithms-go/test"
)

/*
> https://leetcode.com/problems/repeated-dna-sequences/description/

All DNA is composed of a series of nucleotides abbreviated as A, C, G, and T, for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.

Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.

Example:

Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"

Output: ["AAAAACCCCC", "CCCCCAAAAA"]

*/

func findRepeatedDnaSequences(s string) []string {
	var m = make(map[string]int)
	for i := 0; i+10 <= len(s); i++ {
		m[s[i:i+10]]++
	}

	var max int
	var max_s []string
	for k, v := range m {
		if v > max {
			max = v
			max_s = []string{k}
		} else if v == max {
			max_s = append(max_s, k)
		}
	}

	if max > 1 {
		sort.Strings(max_s) //for test
		return max_s
	}

	return nil
}

func Test_187(t *testing.T) {
	test.Runs(t, findRepeatedDnaSequences, []*test.Case{
		{Input: `AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT`, Output: `[AAAAACCCCC,CCCCCAAAAA]`},
		{Input: `AAAAAAAAAA`, Output: ``},
	})
}
