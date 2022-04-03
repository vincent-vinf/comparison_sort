package impl

type SelectSort struct{}

// O(n^2) 数组实现不稳定
func (i SelectSort) Sort(input []int) []int {
	l := len(input)
	for i := l - 1; i > 0; i-- {
		maxIndex := i
		for j := 0; j < i; j++ {
			if input[j] > input[maxIndex] {
				maxIndex = j
			}
		}
		input[maxIndex], input[i] = input[i], input[maxIndex]
	}
	return input
}
