package palindromenumber

import "strconv"

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	r := len(s) - 1

	for i := 0; i < r; i++ {
		if s[i] != s[r] {
			return false
		}
		r--
	}

	return true

}
