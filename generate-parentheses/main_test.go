package main_test

import (
	"sort"
	"strings"
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
	if n == 1 {
		return []string{"()"}
	}
	ret := []string{}
	graph := make([]int, n)

LOOP:
	for {
		for i := n - 1; i >= 0; i-- {
			if graph[i] > i {
				if i == 1 {
					break LOOP
				}
				graph[i-1]++
				for j := i; j < n; j++ {
					graph[j] = graph[i-1]
				}
			}
		}
		ps := make([]string, n)

		for i := n - 1; i >= 0; i-- {
			if ps[i] == "" {
				ps[i] = "()"
			} else {
				ps[i] = "(" + ps[i] + ")"
			}
			if graph[i] != 0 {
				ps[graph[i]-1] += ps[i]
				ps[i] = ""
			}
		}
		ret = append(ret, strings.Join(ps, ""))

		graph[n-1]++
	}

	sort.Strings(ret)
	return ret
}
