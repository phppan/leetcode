package main

import (
	"strings"
)

func myAtoi(str string) int {
	inputStr := strings.TrimSpace(str)
	rs := 0

	i := 0
	sLen := len(inputStr)
	flag := 0
	for {
		if i >= sLen {
			break
		}


		chNum := int(inputStr[i])

		if chNum > 57 || chNum < 48 {
			if flag == 0 {
				if chNum == 45 {
					flag = -1
				}else if chNum == 43 {
					flag = 1
				}else{
					break
				}
			}else{
				break
			}
		}else{
			if flag == 0 {
				flag = 1
			}
		}

		if chNum >= 48 && chNum <= 57 {
			rs = rs * 10 + (chNum - 48)
		}
		i++

		if flag == 1 && rs >= 2147483647 {
			return 2147483647
		}

		if flag == -1 && rs >= 2147483648 {
			return -2147483648
		}
	}



	rs = rs * flag

	if rs > 2147483647 {
		return 2147483647
	}

	if rs < -2147483648 {
		return -2147483648
	}

	return rs
}