package interation

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	// test ซำซ้อน 100% coverage อยู่แล้ว ถึงแม้ comment อันนี้ไป
	// go test -cover

	// t.Run("5 number ", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}

	// 	got := Sum(numbers)
	// 	want := 15

	// 	if got != want {
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// })

	t.Run("any number ", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

// We need a new function called SumAll which will take a varying number of slices,
// returning a new slice containing the totals for each slice passed in.

// For example
// SumAll([]int{1,2}, []int{0,9}) would return []int{3, 9}

// or

// SumAll([]int{1,1,1}) would return []int{3}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
