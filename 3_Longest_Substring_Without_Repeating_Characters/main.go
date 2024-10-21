package main

import "fmt"

func lengthOfLongestSubstring1(s string) int {

	if len(s) < 2 { // case: "a", ""
		return len(s)
	}
	res := 0
	for i := 0; i < len(s); i++ {
		hmap := make(map[byte]bool)
		for j := i; j <= len(s); j++ {
			if j == len(s) { // case: "aab"
				if j-i > res {
					res = j - i
				}
			} else if hmap[s[j]] {
				if j-i > res {
					res = j - i
				}
				break
			} else {
				hmap[s[j]] = true
			}
		}
	}

	return res
}

func lengthOfLongestSubstring2(s string) int {
	// store the elements in the sliding window
	hmap := make(map[byte]int)
	// l and r refer to the left and right margin of the sliding window
	res, l, r := 0, 0, 0
	for ; r < len(s); r++ {
		inx, ok := hmap[s[r]]
		// the char hasn't been covered by SW or once been shrinked out
		if !ok || inx == -1 {
			hmap[s[r]] = r
			continue
		}

		if r-l > res {
			res = r - l
		}

		// update the index of current repeated char
		hmap[s[r]] = r
		// and shrink SW by moving the left margin, during which all the chars are removed out
		for ; l < inx; l++ {
			hmap[s[l]] = -1
		}
		l++
	}

	// the last window. case: "acacb"
	if r-l > res {
		res = r - l
	}

	return res
}

func lengthOfLongestSubstring3(s string) int {
	// store the elements in the sliding window
	hmap := [128]int{}
	// l and r refer to the left and right margin of the sliding window
	res, l, r := 0, 0, 0
	for ; r < len(s); r++ {
		inx := hmap[s[r]]
		hmap[s[r]] = r + 1
		// the char hasn't been covered by SW or once been shrinked out
		if inx == 0 {
			continue
		}

		if r-l > res {
			res = r - l
		}

		// and shrink SW by moving the left margin, during which all the chars are removed out
		for ; l < inx-1; l++ {
			hmap[s[l]] = 0
		}
		l++
	}

	// the last window. case: "acacb"
	if r-l > res {
		res = r - l
	}

	return res
}

func main() {
	str := "abcdbae"
	fmt.Println(lengthOfLongestSubstring1(str))
	fmt.Println(lengthOfLongestSubstring2(str))
	fmt.Println(lengthOfLongestSubstring3(str))
}
