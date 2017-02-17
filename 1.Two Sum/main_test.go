package main

import (
	"os"
	"testing"
)

var cases = make([]testCase, 100)

func BenchmarkTwoSum3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(cases); j++ {
			twoSum3(cases[j].nums, cases[j].result)
		}
	}
}
func BenchmarkTwoSum4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(cases); j++ {
			twoSum4(cases[j].nums, cases[j].result)
		}
	}
}
func TestMain(m *testing.M) {
	initTestData()
	os.Exit(m.Run())
}
func initTestData() {
	for i := range cases {
		cases[i].result, cases[i].nums = generate(1000)
	}
}
