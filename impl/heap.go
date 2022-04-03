package impl

type HeapSort struct{}

// O(nlogn) 不稳定
func (s HeapSort) Sort(input []int) []int {
	//建堆(最大堆)
	s.createHeap(input)
	for i := len(input) - 1; i > 0; i-- {
		//将堆顶数据(最大值)移动到最终位置，并用堆最后一个元素填补
		input[0], input[i] = input[i], input[0]
		//调整堆
		s.shiftDown(input[:i], 0)
	}
	return input
}

// 建堆操作
func (s HeapSort) createHeap(input []int) {
	l := len(input)
	for i := l/2 - 1; i >= 0; i-- {
		s.shiftDown(input, i)
	}
}

// 使得堆重新有序
func (s HeapSort) shiftDown(heap []int, index int) {
	l := len(heap)
	// father
	f := index
	// child
	c := index*2 + 1
	for c < l {
		//如果右儿子存在，且大于左儿子
		if c+1 < l && heap[c+1] > heap[c] {
			c++
		}
		//如果孩子比父亲大，当前值下移
		if heap[f] < heap[c] {
			heap[f], heap[c] = heap[c], heap[f]
			f = c
			c = c*2 + 1
		} else {
			return
		}
	}
	//	可以实现一个不交换的版本
	//  可以理解为，为一个元素寻找一个正确的位置(空穴)，算法模拟了空穴下移的过程
}
