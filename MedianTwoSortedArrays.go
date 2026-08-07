package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	p1 := 0
	p2 := 0
	i := 0
	var merged []int
	mergedL := len(nums1) + len(nums2)

	if len(nums1) == len(nums2) {
		for i <= len(nums1) {
			if p1 < len(nums1) && p2 < len(nums2) {
				if nums1[p1] < nums2[p2] {
					merged = append(merged, nums1[p1])
					p1++
				} else {
					merged = append(merged, nums2[p2])
					p2++
				}
			} else if p1 < len(nums1) {
				merged = append(merged, nums1[p1])
				p1++
			} else {
				merged = append(merged, nums2[p2])
				p2++
			}
			i++
		}

		return float64(merged[i-1]+merged[i-2]) / 2.0
	} else if len(nums1) == 0 {
		if len(nums2)%2 == 0 {
			return float64(nums2[len(nums2)/2]+nums2[(len(nums2)/2)-1]) / 2.0
		}
		return float64(nums2[len(nums2)/2])
	} else if len(nums2) == 0 {
		if len(nums1)%2 == 0 {
			return float64(nums1[len(nums1)/2]+nums1[(len(nums1)/2)-1]) / 2.0
		}
		return float64(nums1[len(nums1)/2])
	}

	for len(merged) < mergedL {
		if p1 < len(nums1) && p2 < len(nums2) {
			if nums1[p1] < nums2[p2] {
				merged = append(merged, nums1[p1])
				p1++
			} else {
				merged = append(merged, nums2[p2])
				p2++
			}
		} else if p1 < len(nums1) {
			merged = append(merged, nums1[p1])
			p1++
		} else {
			merged = append(merged, nums2[p2])
			p2++
		}

		if len(merged) == (mergedL/2)+1 {
			if mergedL%2 == 0 {
				result := (float64(merged[i] + merged[i-1])) / 2.0
				return result
			}

			return float64(merged[i])
		}

		i++
	}

	return 0
}

func main() {
	findMedianSortedArrays([]int{}, []int{1})
}
