package main

import (
	"sort"
)

func threeSum(nums[]int) [][]int {
	sort.Ints(nums)
	rs := make([][]int, 0)

	numsLen := len(nums)

	for k := 0; k < numsLen; k++ {
		if nums[k] > 0 {
			break
		}

		if k > 0 && nums[k] == nums[k - 1] {
			continue
		}

		target := 0 - nums[k]
		i := k + 1
		j := numsLen - 1

		for {
			if i >= j {
				break
			}

			if nums[i] + nums[j] == target {
				rs = append(rs, []int{nums[k], nums[i], nums[j]})
				for {
					if i >= j || nums[i] != nums[i + 1] {
						break
					}
					i++
				}
				for {
					if i >= j || nums[j] != nums[j - 1] {
						break
					}
					j--
				}

				i++
				j--
			}else if nums[i] + nums[j] < target {
				i++
			}else {
				j--
			}
		}
	}

	return rs
}

/*

import (
	"strconv"
	"bytes"
)
func threeSum(nums []int) [][]int {

	rs := make([][]int, 0)

	numsLen := len(nums)
	numsMap := make(map[int]int, numsLen)
	flagMap := make(map[string]bool, numsLen)

	for  _, value := range nums {
		if val, ok := numsMap[value]; ok {
			numsMap[value] =  val + 1
		}else{
			numsMap[value] = 1
		}
	}

	negativeNums := make([]int, 0) // 负数
	positiveNums := make([]int, 0) // 正数
	for  _, value := range nums {
		if value < 0 {
			negativeNums = append(negativeNums, value)
		}else if value > 0 {
			positiveNums = append(positiveNums, value)
		}
	}

	zeroCnt := 0
	for  _, value := range nums {
		if value == 0 {
			zeroCnt++
		}
	}

	// 三种情况

	// 3. 没有0的情况
	//3.1 两个正数，一个负数

	negativeNumsLen := len(negativeNums)

	// 3.2 两个负数，一个正数
	if negativeNumsLen > 1 {
		for i := 0; i < negativeNumsLen - 1; i++ {
			for j := i + 1; j < negativeNumsLen; j++ {
				key := 0 - negativeNums[i] - negativeNums[j]
				flagKey := createFlagKey(key, negativeNums[i], negativeNums[j])
				_, flagOk := flagMap[flagKey]
				if _, ok := numsMap[key]; ok && !flagOk {
					rs = append(rs, []int{negativeNums[i], negativeNums[j], key})
					flagMap[flagKey] = true
				}
			}
		}
	}

	// 2. 一个0的情况
	if zeroCnt >= 1 {
		for  _, value := range negativeNums {
			key := 0 - value
			flagKey := strconv.Itoa(value) + strconv.Itoa(0) +  strconv.Itoa(key)
			_, flagOk := flagMap[flagKey]
			if _, ok := numsMap[key]; ok && !flagOk {
				rs = append(rs, []int{value, 0, key})
				flagMap[flagKey] = true
			}
		}
	}

	positiveNumsLen := len(positiveNums)
	if positiveNumsLen > 1 {
		for i := 0; i < positiveNumsLen - 1; i++ {
			for j := i + 1; j < positiveNumsLen; j++ {
				key := 0 - positiveNums[i] - positiveNums[j]
				flagKey := createFlagKey(key, positiveNums[i], positiveNums[j])
				_, flagOk := flagMap[flagKey]
				if _, ok := numsMap[key]; ok && !flagOk {
					rs = append(rs, []int{key, positiveNums[i], positiveNums[j]})
					flagMap[flagKey] = true
				}
			}
		}
	}
	// 1. 三个0的情况
	if zeroCnt >= 3 {
		rs = append(rs, []int{0, 0, 0})
	}


	return rs
}

func createFlagKey(a int, b int, b2 int) string {
	var stringBuffer bytes.Buffer

	if b > b2 {
		stringBuffer.WriteString(strconv.Itoa(a))
		stringBuffer.WriteString("|")
		stringBuffer.WriteString(strconv.Itoa(b))
		stringBuffer.WriteString("|")
		stringBuffer.WriteString(strconv.Itoa(b2))
	}else {
		stringBuffer.WriteString(strconv.Itoa(a))
		stringBuffer.WriteString("|")
		stringBuffer.WriteString(strconv.Itoa(b2))
		stringBuffer.WriteString("|")
		stringBuffer.WriteString(strconv.Itoa(b))
	}

	return stringBuffer.String()
}
*/