package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	l := m + n
	auxPointer := 0

	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}

	for i := 0; i < l; i++ {
		if i < m-1 {
			if nums1[i] >= nums2[auxPointer] {
				for j := i + 1; j < l; j++ {
					nums1[j] = nums1[j-1]
				}
				nums1[i] = nums2[auxPointer]
			} else if nums1[i] <= nums2[auxPointer] {
				for j := i + 1; j < l; j++ {
					nums1[j+1] = nums1[j]
				}
				nums1[i+1] = nums2[auxPointer]
			}
			auxPointer++
		}
	}
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
