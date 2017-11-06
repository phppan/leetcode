package main

func isMatch2(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	if pLen <= 0 {
		return sLen <= 0
	}

	firstMatch := sLen > 0 && (p[0] == s[0] || p[0] == 46)

	if pLen >= 2 && p[1] == 42 {
		return isMatch(s, p[2:]) || firstMatch && isMatch(s[1:], p)
	}else{
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)

	memo := make(map[int]map[int]bool, sLen + 1)

	for i := 0; i < sLen + 1; i++ {
		memo[i] = make(map[int]bool, pLen + 1)
	}

	memo[sLen][pLen] = true

	for i := sLen; i >= 0; i-- {
		for j := pLen - 1; j >= 0; j-- {
			firstMatch := i < sLen && (p[j] == s[i] || p[j] == 46)

			if j + 1 < pLen && p[j + 1] == 42 {
				memo[i][j] = memo[i][j + 2] || firstMatch && memo[i + 1][j]
			}else{
				memo[i][j] = firstMatch && memo[i + 1][j + 1]
			}
		}
	}

	return memo[0][0]
}

