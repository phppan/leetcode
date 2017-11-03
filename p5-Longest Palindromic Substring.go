package main

func longestPalindrome(s string) string {
	sLen := len(s)

	max := 0
	begin := 0
	end := 0
	for i := 1; i < sLen; i++ {
		subLen := 0
		j := 0
		for j = 0; j < i && i + j < sLen; j++ {
			if i - j - 1 >= 0 && s[i - j - 1] == s[i + j] {
				subLen += 2
				if subLen > max {
					begin = i - j - 1
					end = i + j
					max = subLen
				}
			}else{
				break
			}
		}

		subLen = 1
		for j = 0; j < i && i + j + 1< sLen; j++ {
			if i - j - 1 >= 0 && i + j + 1 < sLen  && s[i - j - 1] == s[i + j + 1] {
				subLen += 2
				if subLen > max {
					max = subLen
					begin = i - j - 1
					end = i + j + 1
				}
			}else{
				break
			}
		}
	}

	return s[begin:end + 1]
}