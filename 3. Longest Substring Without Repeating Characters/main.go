// https://leetcode.com/problems/longest-substring-without-repeating-characters/?tab=Description
package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var left, right int
	var l, r int

	length := len(s)
	charMap := make(map[byte]int)
	for i := 0; i < length; i++ {
		pos, find := charMap[s[i]]
		r = i
		charMap[s[i]] = i
		if find {
			if r-l > right-left {
				left, right = l, r
			}
			l = pos + 1
			charMap[s[i]] = i
			for j := range charMap {
				if charMap[j] <= pos {
					delete(charMap, j)
				}
			}
		}
	}
	if r+1-l > right-left {
		return r + 1 - l
	} else {
		return right - left
	}
}
func main() {
	lengthOfLongestSubstring("ababca")
}
