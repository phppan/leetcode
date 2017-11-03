package main


func convert(s string, numRows int) string {
	sLen := len(s)
	if sLen <= 1 || numRows < 2 {
		return s
	}

	rs := ""
	lag := numRows * 2 - 2
	for i := 0; i < numRows; i++ {
		for j := i; j < sLen; j += lag {
			rs = rs + string(s[j])

			index := (j / lag + 1) * lag - i
			if i > 0 && i < numRows - 1 &&  index < sLen {
				rs = rs + string(s[index])
			}
		}
	}

	return rs
}