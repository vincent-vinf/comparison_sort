package impl

type InsertSort struct{}

// O(n^2) 稳定
func (i InsertSort) Sort(input []int) []int {
	l := len(input)
	for i := 1; i < l; i++ {
		j := i - 1
		// 取出即将插入的值
		key := input[i]
		// 从有序的最后一个开始，如果>key则将其赋值给后一个位置(值本身已经被缓存)
		for j >= 0 && input[j] > key {
			input[j+1] = input[j]
			j--
		}
		input[j+1] = key
	}
	return input
}
