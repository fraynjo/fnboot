package math

import "testing"

func TestAdd(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		result := Add(2.2, 1)
		t.Log(result)
		if result != 3.2 {
			t.Fatal("fail")
		}
	})
}

func TestCeil(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		result := Ceil(3.3)
		t.Log(result)
		if result != 4 {
			t.Fatal("fail")
		}
	})
}

func TestDivide(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		result := Divide(3, 2)
		t.Log(result)
		if result != 1.5 {
			t.Fatal("fail")
		}
	})
}

func TestFloor(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		result := Floor(3.3)
		t.Log(result)
		if result != 3 {
			t.Fatal("fail")
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		result := Max([]float64{1.1, 2.2, 3.3, 5.5})
		t.Log(result)
		if result != 5.5 {
			t.Fatal("fail")
		}
	})
}

func TestMean(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		result := Mean([]int{1, 2, 3})
		t.Log(result)
		if result != 2 {
			t.Fatal("fail")
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		result := Min([]float64{1.1, 2.2, 3.3, 5.5})
		t.Log(result)
		if result != 1.1 {
			t.Fatal("fail")
		}
	})
}

func TestMultiply(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		result := Multiply(2, 2.2)
		t.Log(result)
		if result != 4.4 {
			t.Fatal("fail")
		}
	})
}

func TestRound(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		result := Round(9.84521, 2)
		t.Log(result)
		if result != 9.85 {
			t.Fatal("fail")
		}
	})
}

func TestSubtract(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		result := Subtract(2, 3)
		t.Log(result)
		if result != -1 {
			t.Fatal("fail")
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		result := Sum([]int{1, 2, 3})
		t.Log(result)
		if result != 6 {
			t.Fatal("fail")
		}
	})
}
