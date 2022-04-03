package impl

type MergeSort struct{}

// O(nlogn) 稳定
func (s MergeSort) Sort(input []int) []int {
	output := make([]int, len(input))
	s.subSort(input, output)
	return output
}
func (s MergeSort) subSort(input []int, reg []int) {
	l := len(input)
	if l <= 1 {
		return
	}
	mid := l / 2
	//对左右两部分进行归并排序
	s.subSort(input[:mid], reg[:mid])
	s.subSort(input[mid:], reg[mid:])

	i, j, k := 0, mid, 0
	// 如果有一路已经处理完，则退出循环
	for i < mid && j < l {
		//将较小的值放到临时数组中
		if input[i] < input[j] {
			reg[k] = input[i]
			k++
			i++
		} else {
			reg[k] = input[j]
			k++
			j++
		}
	}
	//将未归并完的一路直接连接到数组之后
	for i < mid {
		reg[k] = input[i]
		i++
		k++
	}
	for j < l {
		reg[k] = input[j]
		j++
		k++
	}
	//将临时数组的值拷贝
	copy(input, reg)
}
