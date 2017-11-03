package main


func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	if x % 10 == 0 {
		return false
	}

	var tl uint64

	i := x
	tl = 0

	for {
		if i <= 0 {
			break
		}

		tl = tl * 10 + uint64(i % 10)
		i = i / 10
	}

	if tl == uint64(x) {
		return true
	}

	return false
}