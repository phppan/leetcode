package main

import (
	"fmt"
)



func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	hashMap := make(map[int]int)

	i := 0
	for j:= 0; j < n; j++ {
		key := int(s[j])
		if k, ok := hashMap[key]; ok {
			if k > i {
				i = k
			}
		}

		if j - i + 1 > ans {
			ans = j - i + 1
		}

		hashMap[key] = j + 1
	}


	return ans
}

func lengthOfLongestSubstring6(s string) int {
	n := len(s)
	ans := 0
	hashMap := make(map[int]int)

	i := 0
	for j:= 0; j < n; j++ {
		key := int(s[j])
		if k, ok := hashMap[key]; ok {
			if k > i {
				i = k
			}
		}

		if j - i + 1 > ans {
			ans = j - i + 1
		}

		hashMap[key] = j + 1
		fmt.Print(ans, " ")
	}

	sLen := len(s)
	leftCount := make([]int, sLen)
	leftMap := make(map[int]int)
	t := 0
	for  i := 0; i < sLen; i++ {
		key := int(s[i])
		if v, ok := leftMap[key]; ok {
			if v > t {
				t = v
			}
		}

		leftCount[i] = i  - t

		leftMap[key] = i
		fmt.Println("\r\n", leftCount)
	}
	fmt.Println("\r\n", leftCount)

	return ans
}

func lengthOfLongestSubstring4(s string) int {
	sLen := len(s)
	if sLen <= 0 {
		return 0
	}
	//leftCount := make([]int, sLen)
	leftMap := make(map[int]int)
	flag := make([]int, sLen)
	for  i := 0; i < sLen; i++ {
		flag[i] = 0
	}
	for  i := 0; i < sLen; i++ {
		key := int(s[i])
		if v, ok := leftMap[key]; !ok {
			leftMap2 := make(map[int]int)
			for j := i; j >= 0; j-- {
				flag[j] += 1

				key2 := int(s[j])
				if _, ok2 := leftMap2[key2]; ok2 {
					break
				}
				leftMap2[key2] = 1
			}
		}else{
			leftMap2 := make(map[int]int)
			for j := i; j > v; j-- {
				flag[j] += 1

				key2 := int(s[j])
				if _, ok2 := leftMap2[key2]; ok2 {
					break
				}
				leftMap2[key2] = 1
			}
		}
		leftMap[key] = i
		//fmt.Println(flag)
	}


	//fmt.Println(s)
	//fmt.Println(flag)

	max := 0
	for i := 0; i < sLen; i ++ {
		if max < flag[i] {
			max = flag[i]
		}
	}

	return max
}

func lengthOfLongestSubstring3(s string) int {
	sLen := len(s)
	if sLen <= 0 {
		return 0
	}

	leftCount := make([]int, sLen)
	leftMap := make(map[int]int)
	leftCount[0] = 1
	leftMap[int(s[0])] = 0
	for  i := 1; i < sLen; i++ {
		key := int(s[i])
		if v, ok := leftMap[key]; ok {
			leftMap = make(map[int]int)
			leftCount[i] = i - v
			fmt.Println(i)
		}else{
			if s[i] == s[i - 1] {
				leftCount[i] = 1
			}else{
				leftCount[i] = leftCount[i - 1] + 1
			}
		}
		leftMap[key] = i
	}

	flag := make([]int, sLen)
	flagMap := make(map[int]int)
	flag[0] = 1
	flagMap[int(s[0])] = 0
	for i := 1; i < sLen; i++ {
		key := int(s[i])
		if _, ok := flagMap[key]; ok {
			flag[i] = flag[i - 1]
		}else{
			if leftCount[i] > flag[i - 1] {
				flag[i] = leftCount[i] + 1
			}else{
				flag[i] = flag[i - 1] + 1
			}
			flagMap[key] = i
		}
	}


	fmt.Println(s)
	fmt.Println(leftCount)
	fmt.Println(flag)

	max := 0
	for i := 0; i < sLen; i ++ {
		if max < flag[i] {
			max = flag[i]
		}
	}

	return max
}
func lengthOfLongestSubstring2(s string) int {
	sLen := len(s)
	leftCount := make([]int, sLen)
	leftMap := make(map[int]int)
	for  i := 0; i < sLen; i++ {
		key := int(s[i])
		if v, ok := leftMap[key]; !ok {
			leftMap[key] = i
			leftCount[i] = 0
		}else{
			if i > 0 && leftCount[i - 1] + 1 < i - v {
				leftCount[i] = leftCount[i  - 1] + 1
			}else{
				leftCount[i] = i  - v
			}
			leftMap[key] = i
		}
	}

	rightCount := make([]int, sLen)
	rightMap := make(map[int]int)
	for  i := 0; i < sLen; i++ {
		key := int(s[sLen - i - 1])
		if v, ok := rightMap[key]; !ok {
			rightMap[key] = i
			rightCount[i] = 0
		}else{
			if i > 0 && rightCount[i - 1] + 1 < i - v {
				rightCount[i] = rightCount[i  - 1] + 1
			}else{
				rightCount[i] = i  - v
			}
			rightMap[key] = i
		}
	}

	fmt.Println(leftCount)
	fmt.Println(leftMap)
	fmt.Println(rightCount)
	fmt.Println(rightMap)

	max := 0
	for i := 0; i < sLen; i ++ {
		if max < leftCount[i] {
			max = leftCount[i]
		}
	}

	zeroCount := 0;
	for i := 0; i < sLen; i++ {
		if leftCount[i] == 0{
			zeroCount++
			if zeroCount > max {
				max = zeroCount
			}
		}else{
			zeroCount = 0
		}
	}

	for i := 0; i < sLen; i ++ {
		if max < rightCount[i] {
			max = rightCount[i]
		}
	}


	zeroCount = 0;
	for i := 0; i < sLen; i++ {
		if rightCount[i] == 0{
			zeroCount++
			if zeroCount > max {
				max = zeroCount
			}
		}else{
			zeroCount = 0
		}
	}


	return max
}