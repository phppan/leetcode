package main

func reverse(x int) int {
	flag := 1
	if x < 0{
		flag = -1
		x = 0 - x

	}
	rs := 0
	for {
		if x <= 0 {
			break
		}

		rs = rs*10 + x%10
		x = x / 10
	}

	rs = rs * flag

	if rs > 2147483647 {
		return 0
	}

	if rs < -2147483648 {
		return 0
	}

	return rs
}
