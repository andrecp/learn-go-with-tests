package integers

import "testing"

func TestAdder(t *testing.T) {

	t.Run("Saying hello to people", func(t *testing.T) {
		got := add(2, 2)
		want := 4
		if got != want {
			t.Errorf("%d is not %d", got, want)
		}
	})
}
