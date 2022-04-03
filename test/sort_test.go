package test

import (
	"math/rand"
	"sort_go/impl"
	"testing"
	"time"
)

const dataLen = 7000

type Sorter interface {
	Sort([]int) []int
}

func generateData(len int) (re []int) {
	//re = []int{17, 11, 1, 4, 8, 4, 0}
	re = make([]int, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		re[i] = rand.Intn(20)
	}
	return
}

func checkData(in []int) bool {
	if len(in) != dataLen {
		return false
	}
	for i := 0; i < len(in)-1; i++ {
		if in[i] > in[i+1] {
			return false
		}
	}
	return true
}

func testSort(s Sorter) bool {
	return checkData(s.Sort(generateData(dataLen)))
}

func TestBubble(t *testing.T) {
	if !testSort(&impl.BubbleSort{}) {
		t.Error("error")
	}
}

func TestSelect(t *testing.T) {
	if !testSort(&impl.SelectSort{}) {
		t.Error("error")
	}
}

func TestInsert(t *testing.T) {
	if !testSort(&impl.ShellSort{}) {
		t.Error("error")
	}
}

func TestMerge(t *testing.T) {
	if !testSort(&impl.MergeSort{}) {
		t.Error("error")
	}
}

func TestQuick(t *testing.T) {
	if !testSort(&impl.QuickSort{}) {
		t.Error("error")
	}
}

func TestHeap(t *testing.T) {
	if !testSort(&impl.HeapSort{}) {
		t.Error("error")
	}
}

func TestShell(t *testing.T) {
	if !testSort(&impl.ShellSort{}) {
		t.Error("error")
	}
}
