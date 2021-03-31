package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubleSortWorstCase(t *testing.T) {
	els := []int{9, 8, 7, 6, 5}
	BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, []int{5, 6, 7, 8, 9}, els)
}

func TestBubleSortBestCase(t *testing.T) {
	els := []int{5, 6, 7, 8, 9}
	BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, []int{5, 6, 7, 8, 9}, els)
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	els := getElements(5)
	BubbleSort(els)

	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, []int{0, 1, 2, 3, 4}, els)
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}
