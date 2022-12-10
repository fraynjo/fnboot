package internal

import "testing"

func TestInRange(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		result := InRange(2.2, 1, 4.5)
		t.Log(result)
		if result != true {
			t.Fatal("fail")
		}
	})
}

func TestRandom(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		result := Random(0, 10)
		t.Log(result)
		if result < 0 || result > 10 {
			t.Fatal("fail")
		}
	})
}
