package fnboot

import (
	"math"
	"testing"
)

func TestChunk(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7}
		result := Chunk(arr, 2)
		if len(result) != 4 {
			t.Fatal("fail")
		}
	})
	t.Run("string", func(t *testing.T) {
		arr := []string{"a", "b", "c", "d"}
		result := Chunk(arr, 3)
		if len(result) != 2 {
			t.Fatal("fail")
		}
	})
	t.Run("interface{}", func(t *testing.T) {
		arr := []interface{}{[]int{1, 2}, []string{"a", "b"}, []bool{true, false}}
		result := Chunk(arr, 2)
		if len(result) != 2 {
			t.Fatal("fail")
		}
	})
}

func TestCompact(t *testing.T) {
	arr := []interface{}{1, false, true, 0, "hello", "", nil}
	result := Compact(arr)
	if len(result) != 3 {
		t.Fatal("fail")
	}
}

func TestConcat(t *testing.T) {
	arr := []interface{}{1}
	result := Concat[interface{}](arr, 2, 3, []int{4})
	t.Log(result)
	if len(result) != 4 {
		t.Fatal("fail")
	}
}

func TestDifference(t *testing.T) {
	arr := []int{2, 1}
	arr2 := []int{2, 3}
	result := Difference[int](arr, arr2)
	t.Log(result)
	if result[0] != 1 {
		t.Fatal("fail")
	}
}

func TestDifferenceBy(t *testing.T) {
	arr := []float64{2.2, 1.4}
	arr2 := []float64{2.3, 3.6}
	result := DifferenceBy[float64, int](arr, arr2, func(x float64) int {
		return int(math.Floor(x))
	})
	t.Log(result)
	if result[0] != 1 {
		t.Fatal("fail")
	}
}
