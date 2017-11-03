package main

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	lenCount := len1  + len2

	rs := 0
	if lenCount % 2 == 0 {
		flag := int(math.Ceil(float64(lenCount / 2))) - 1
		flag2 := flag + 1
		i := 0
		j := 0
		for {
			if i + j == flag || i + j == flag2 {
				// 如果第一个数组用完了
				if i >= len1 {
					rs += nums2[j]
				}else if j >= len2 {
					rs += nums1[i]
				}else if nums1[i] <= nums2[j] {
					rs += nums1[i]
				}else {
					rs += nums2[j]
				}
			}

			if i + j == flag2 {
				break
			}

			if i >= len1 {
				j++
			}else if j >= len2 {
				i++
			}else if nums1[i] <= nums2[j] {
				i++
			}else {
				j++
			}
		}

		return float64(rs)  / 2
	}else{
		flag := int(math.Ceil(float64((lenCount - 1)/ 2)))
		i := 0
		j := 0
		for {
			if i+j == flag {
				if i >= len1 {
					rs += nums2[j]
				} else if j >= len2 {
					rs += nums1[i]
				} else if nums1[i] <= nums2[j] {
					rs += nums1[i]
				} else {
					rs += nums2[j]
				}
				break
			}

			if i >= len1 {
				j++
			} else if j >= len2 {
				i++
			} else if nums1[i] <= nums2[j] {
				i++
			} else {
				j++
			}
		}
		return float64(rs)
	}
}
