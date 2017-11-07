package main


func longestCommonPrefix(strs []string) string {
	sLen := len(strs)

	if sLen == 0 {
		return ""
	}

	if sLen == 1 {
		return strs[0]
	}

	minLen := 2147483647
	for i := 0; i < sLen; i++ {
		lenTmp := len(strs[i])
		if lenTmp < minLen {
			minLen = lenTmp
		}
	}

	i := 0
	for {
		if i >= minLen {
			break
		}

		flag := strs[0][i]
		for j := 1; j < sLen; j++ {
			if strs[j][i] != flag {
				flag = strs[j][i]
				break
			}
		}

		if flag != strs[0][i] {
			break
		}
		i++
	}


	return strs[0][0:i]
}
