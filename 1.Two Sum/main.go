package main

import (
	"fmt"
	"runtime"
	"sort"
)

func twoSum2(nums []int, target int) []int {
	// gen map
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; ok {
			if v+v == target {
				return []int{m[v], i}
			}
		}
		m[v] = i
	}
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] > target {
				break
			}
			if nums[i]+nums[j] == target {
				return []int{m[nums[i]], m[nums[j]]}
			}
		}
	}
	return nil
}
func twoSum3(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if wantIndex, ok := m[target-n]; ok {
			return []int{wantIndex, i}
		}
		m[n] = i
	}
	return nil
}
func twoSum4(nums []int, target int) []int {
	twoSumAsc := func(nums []int, target int) []int {
		m := make(map[int]int)
		length := len(nums)
		for i := 0; i < length; i++ {
			n := nums[i]
			if wantIndex, ok := m[target-n]; ok {
				return []int{wantIndex, i}
			}
			m[n] = i
		}
		return nil
	}
	twoSumDes := func(nums []int, target int) []int {
		m := make(map[int]int)
		length := len(nums)
		for i := length - 1; i >= 0; i-- {
			n := nums[i]
			if wantIndex, ok := m[target-n]; ok {
				return []int{i, wantIndex}
			}
			m[n] = i
		}
		return nil
	}
	if runtime.GOMAXPROCS(0) > 1 {
		c := make(chan []int)
		go func() {
			result := twoSumDes(nums, target)
			c <- result
		}()
		go func() {
			result := twoSumAsc(nums, target)
			c <- result
		}()
		return <-c
	}
	return twoSumDes(nums, target)

}
func main() {
	s := []int{3, 2, 4}
	fmt.Println(twoSum4(s, 6))
}
