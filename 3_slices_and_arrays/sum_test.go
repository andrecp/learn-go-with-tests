package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Test adding five integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Got %d, wanted %d, summed %v", got, want, numbers)
		}
	})

	t.Run("Test adding 3 items", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("Got %d, wanted %d, summed %v", got, want, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {
	t.Run("Test suming multiple slices of integers", func(t *testing.T) {
		numbers1 := []int{1, 2, 3, 4, 5}
		numbers2 := []int{1, 2, 3}

		got := SumAll(numbers1, numbers2)
		want := []int{15, 6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

// func ExampleSum() {
// 	repatedLetter := Sum("a", 5)
// 	fmt.Println(repatedLetter)
// 	// Output: aaaaa
// }

// func BenchmarkSum(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Sum("a", 5)
// 	}
// }
