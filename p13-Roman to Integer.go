package main

func romanToInt(s string) int {
	sLen := len(s)
	roman := map[string]int{"M": 1000,"D": 500 ,"C": 100,"L": 50,"X": 10,"V": 5,"I": 1}

	if sLen == 0 {
		return 0
	}

	if sLen == 1 {
		return roman[string(s)]
	}
	rs := 0
	key := ""
	key2 := ""


	for i := 0; i < sLen - 1; i++ {
		key = string(s[i])
		key2 = string(s[i + 1])
		if roman[key] < roman[key2] {
			rs -= roman[key]
		}else{
			rs += roman[key]
		}
	}

	return rs + roman[key2]
}