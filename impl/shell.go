package impl

type ShellSort struct{}

// O(n^2)-O(nlog2n) 在小数据量的情况下甚至优于快排和堆排 不稳定
func (i ShellSort) Sort(input []int) []int {
	l := len(input)
	for gap := l / 2; gap > 0; gap /= 2 {
		for j := gap; j < l; j++ {
			tmp := input[j]
			k := j - gap
			for ; k >= 0 && input[k] > tmp; k -= gap {
				input[k+gap] = input[k]
			}
			input[k+gap] = tmp
		}
	}
	return input
}
