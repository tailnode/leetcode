package main

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

var l1 *ListNode
var l2 *ListNode

const L1_LEN = 1000
const L2_LEN = 899

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addTwoNumbers1(l1, l2)
	}
}
func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addTwoNumbers2(l1, l2)
	}
}
func TestMain(m *testing.M) {
	initTestData()
	os.Exit(m.Run())
}
func initTestData() {
	rand.Seed(time.Now().UnixNano())
	l1Nums := make([]int, L1_LEN)
	l2Nums := make([]int, L2_LEN)
	for i := 0; i < len(l1Nums); i++ {
		l1Nums[i] = int(rand.Int31n(10))
	}
	for i := 0; i < len(l2Nums); i++ {
		l2Nums[i] = int(rand.Int31n(10))
	}
	l1, l2 = genList(l1Nums), genList(l2Nums)
}
