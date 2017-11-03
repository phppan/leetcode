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

	var tl, tr uint64
	var j int

	i := x
	tl = 0
	tr = 0
	j = 1


	for {
		if i <= 0 {
			break
		}

		tl = tl * 10 + uint64(i % 10)
		j = j * 10
		i = i / 10
	}

	i = x
	for {
		if i <= 0 {
			break
		}

		tr = tr * 10 + uint64(i / j)
		i = i % j
		j = j / 10
	}

	if tl == tr {
		return true
	}

	return false
}