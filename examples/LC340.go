package examples

import (
	"fmt"
)

type testCase struct {
	s      string
	k      int
	output int
}

// https://github.com/doocs/leetcode/tree/main/solution/0300-0399/0340.Longest%20Substring%20with%20At%20Most%20K%20Distinct%20Characters
func LC430() {
	fmt.Println("\nRunning Longest Substring with At Most K Distinct Characters (LC340) testcase...")
	testcases := []testCase{
		{s: "eceba", k: 2, output: 3},
		{s: "aa", k: 1, output: 2},
		{s: "", k: 0, output: 0},
		{s: "a", k: 0, output: 0},
	}

	sol := func(s string, k int) int {
		ans := 0
		i, j := 0, 0
		set := make(map[string]struct{})
		for j < len(s) {
			set[string(s[j])] = struct{}{}
			if len(set) <= k {
				ans = max(ans, j-i+1)
			} else {
				delete(set, string(s[i]))
				i++
			}
			j++
		}
		return ans
	}

	for i, t := range testcases {
		o := sol(t.s, t.k)
		i += 1
		if t.output != o {
			fmt.Printf("Testcase #%d failed.\n Expected: %d, Output: %d\n", i, t.output, o)
		} else {
			fmt.Printf("Testcase #%d passed.\n", i)
		}
	}
}
