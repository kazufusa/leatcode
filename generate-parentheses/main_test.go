package main_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_1(t *testing.T) {
	var tests = []struct {
		expected []string
		given    int
	}{
		{[]string{"()"}, 1},
		{[]string{"(())", "()()"}, 2},
		{[]string{"((()))", "(()())", "(())()", "()(())", "()()()"}, 3},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.expected[0], func(t *testing.T) {
			actual := generateParenthesis(tt.given)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func generateParenthesis(n int) []string {
	var patterns [][]int
	pattern := make([]int, n)
	for {
		patterns = append(patterns, pattern)
	}
	ret := []string{}
	sort.Strings(ret)
	return ret
}
