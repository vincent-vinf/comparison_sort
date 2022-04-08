package impl

type QuickSort struct{}

func (s QuickSort) Sort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := s.partition(nums, len(nums)/2)
	s.Sort(nums[:mid])
	s.Sort(nums[mid+1:])
	return nums
}

func (s QuickSort) partition(nums []int, pivot int) int {
	n := len(nums)
	pivotValue := nums[pivot]
	nums[pivot], nums[n-1] = nums[n-1], nums[pivot]
	storeIndex := 0
	for i := 0; i < n; i++ {
		if nums[i] < pivotValue {
			nums[i], nums[storeIndex] = nums[storeIndex], nums[i]
			storeIndex++
		}
	}
	nums[n-1], nums[storeIndex] = nums[storeIndex], nums[n-1]
	return storeIndex
}
