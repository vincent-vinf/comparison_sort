package impl

type QuickSort struct{}

func (s QuickSort) Sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivotNewIndex := partition(nums)
	s.Sort(nums[:pivotNewIndex])
	s.Sort(nums[pivotNewIndex+1:])
	return nums
}

func partition(nums []int) int {
	n := len(nums)
	pivotValue := nums[n-1]
	storeIndex := 0
	for i := 0; i < n; i++ {
		if nums[i] < pivotValue {
			nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
			storeIndex++
		}
	}
	nums[n-1], nums[storeIndex] = nums[storeIndex], nums[n-1]
	return storeIndex
}
