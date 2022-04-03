package impl

type QuickSort struct{}

// O(nlogn) 不稳定
func (s QuickSort) Sort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	// 选择第一个元素为枢纽
	pivot := input[0]
	i, j := 1, len(input)-1
	for i < j {
		//从左向右，直到找到input[i]大于枢纽停止
		for i < len(input) && input[i] < pivot {
			i++
		}
		//从右向左，直到找到input[j]小于枢纽停止
		for j > 0 && input[j] > pivot {
			j--
		}
		//两者并未交叉，则交换
		if i < j {
			input[i], input[j] = input[j], input[i]
		}
	}
	if input[j] < input[0] {
		input[j], input[0] = input[0], input[j]
	}
	//递归对左右两侧进行快排
	s.Sort(input[:j])
	s.Sort(input[j+1:])
	return input
}
