package impl

type BubbleSort struct {
}

// O(n^2) 稳定
func (b *BubbleSort) Sort(input []int) []int {
	l := len(input)
	for i := l - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}
