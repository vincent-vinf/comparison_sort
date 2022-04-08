package impl

import "sync"

type ConcurrentQSort struct{}

func (s ConcurrentQSort) Sort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	mid := s.partition(nums, n/2)
	if n < 512 {
		s.Sort(nums[:mid])
		s.Sort(nums[mid+1:])
	} else {
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			s.Sort(nums[:mid])
			wg.Done()
		}()
		go func() {
			s.Sort(nums[mid+1:])
			wg.Done()
		}()
		wg.Wait()
	}
	return nums
}

func (s ConcurrentQSort) partition(nums []int, pivot int) int {
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
