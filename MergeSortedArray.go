package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
	}

	m--
	n--

	for i := len(nums1) - 1; i >= 0 && n >= 0; i-- {
		if m >= 0 && nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
			continue
		}
		nums1[i] = nums2[n]
		n--
	}
}

func main() {
	merge([]int{2, 0}, 1, []int{1}, 1)
}
