package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		if got != want {
			t.Errorf("%d is not %d", got, want)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
