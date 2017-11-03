package main


func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for k, v := range nums {
		numsMap[v] = k
	}

	for k, v := range nums {
		if  k1, ok := numsMap[target - v]; ok && k1 != k {
			return []int{k, k1}
		}
	}

	return []int{0, 1}
}
