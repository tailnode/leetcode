package main

import "testing"

var cases = []struct {
	s    string
	want int
}{
	{"", 0},
	{"a", 1},
	{"ab", 2},
	{"abc", 3},
	{"aa", 1},
	{"aba", 2},
	{"ababc", 3},
	{"ababca", 3},
	{"ababcd", 4},
}

func Test1(t *testing.T) {
	for _, c := range cases {
		get := lengthOfLongestSubstring(c.s)
		if get != c.want {
			t.Errorf("case [%s] want %d, but get %d\n", c.s, c.want, get)
		}
	}
}
