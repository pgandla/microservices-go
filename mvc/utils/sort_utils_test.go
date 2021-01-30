package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	els := []int{5, 4, 3, 2, 1}
	BubbleSort(els)
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, els[4])
}

func getElements(n int) []int {
	res := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		res[i] = j
		i++
	}
	return res
}

func BenchmarkBubbleSort(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}
