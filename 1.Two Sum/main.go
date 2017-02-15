package main

import (
	"fmt"
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
	m:=make(map[int]int)
	for i,n:=range nums{
		want:=target-n
		if wantIndex,ok:=m[want];ok{
			return []int{wantIndex,i}
		}
		m[n]=i
	}
	return nil
}
func main() {
	s:=[]int{3,2,4}
	fmt.Println(twoSum3(s, 6))
}
