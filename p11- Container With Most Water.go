package main


func maxArea(height []int) int {
	hLen := len(height)

	i := 0
	j := hLen - 1
	max := 0

	for {

		if i >= j {
			break
		}

		area := minHeight(height[i], height[j]) * (j - i)
		if max < area {
			max = area
		}

		if height[i] < height[j] {
			i++
		}else {
			j--
		}
	}

	return max
}

func minHeight (left int, right int) int{
	if left > right {
		return right
	}

	return left
}