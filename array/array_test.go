package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("A collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	// t.Run("Sum of 2 collections", func(t *testing.T) {
	// 	numbers1 := []int{1, 2}
	// 	numbers2 := []int{4, 5, 6}

	// 	got := SumAll(numbers1, numbers2)
	// 	want := []int{3, 15}

	// 	if len(got) != len(want) {
	// 		t.Errorf("got %d want %d, given %v and %v", len(got), len(want), numbers1, numbers2)
	// 	}
	// 	for i, _ := range got {
	// 		if got[i] != want[i] {
	// 			t.Errorf("got %v want %v, given %v and %v", got, want, numbers1, numbers2)
	// 		}
	// 	}
	// })

	t.Run("Sum of 2 collections", func(t *testing.T) {

		got := SumAll([]int{1, 2}, []int{4, 5, 6})
		want := []int{3, 15}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Sum of more than 2 collections", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{4, 5, 6}, []int{7, 8, 9, 10})
		want := []int{3, 15, 34}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
