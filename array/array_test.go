package array

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

func TestDrop(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4}
	result := Drop(arr, 1)
	t.Log(result)
	if len(result) != 3 {
		t.Fatal("fail")
	}
}

func TestDropRight(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4}
	result := DropRight(arr, 2)
	t.Log(result)
	if len(result) != 2 {
		t.Fatal("fail")
	}
}

func TestFill(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4}
	result := Fill(arr, 0.0)
	t.Log(result)
	if result[0] != 0.0 {
		t.Fatal("fail")
	}
}

func TestFindIndex(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4}
	result := FindIndex(arr, 3.3)
	t.Log(result)
	if result != 2 {
		t.Fatal("fail")
	}
}

func TestFindLastIndex(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	result := FindLastIndex(arr, 3.3)
	t.Log(result)
	if result != 4 {
		t.Fatal("fail")
	}
}

func TestFirst(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	result := First(arr)
	t.Log(result)
	if result != 2.2 {
		t.Fatal("fail")
	}
}

func TestIndexOf(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	result := IndexOf(arr, 3.3)
	t.Log(result)
	if result != 2 {
		t.Fatal("fail")
	}
}

func TestInitial(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	result := Initial(arr)
	t.Log(result)
	if len(result) != 4 {
		t.Fatal("fail")
	}
}

func TestIntersection(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	arr1 := []float64{4.4}
	arr2 := []float64{1.4, 4.4}
	arr3 := []float64{1.4, 4.4}
	result := Intersection(arr, arr1, arr2, arr3)
	t.Log(result)
	if len(result) != 1 {
		t.Fatal("fail")
	}
}

func TestIntersectionBy(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	arr1 := []float64{4.4}
	arr2 := []float64{4.3, 1.5}
	result := IntersectionBy(math.Floor, arr, arr1, arr2)
	t.Log(result)
	if len(result) != 1 {
		t.Fatal("fail")
	}
}

func TestInterfaceToArray(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3}
	result := InterfaceToArray(arr)
	t.Log(result)
	if len(result) != 3 {
		t.Fatal("fail")
	}
}

func TestJoin(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3}
	result := Join(arr, "-")
	t.Log(result)
	if result != "2.2-1.4-3.3" {
		t.Fatal("fail")
	}
}

func TestLast(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	result := Last(arr)
	t.Log(result)
	if result != 3.3 {
		t.Fatal("fail")
	}
}

func TestLastIndexOf(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	result := LastIndexOf(arr, 3.3)
	t.Log(result)
	if result != 4 {
		t.Fatal("fail")
	}
}

func TestNth(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	result := Nth(arr, -1)
	t.Log(result)
	if float64(result) != 4.4 {
		t.Fatal("fail")
	}
}

func TestPull(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	result := Pull(arr, 3.3)
	t.Log(result)
	if len(result) != 3 {
		t.Fatal("fail")
	}
}

func TestPullAll(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	result := PullAll(arr, []float64{2.2, 3.3})
	t.Log(result)
	if len(result) != 2 {
		t.Fatal("fail")
	}
}

func TestPullAt(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	result := PullAt(arr, []int{1, 3})
	t.Log(result)
	if len(result) != 3 {
		t.Fatal("fail")
	}
}

func TestRemove(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	result := Remove(arr, func(v float64) bool {
		return v == 3.3
	})
	t.Log(result)
	if len(result) != 3 {
		t.Fatal("fail")
	}
}

func TestReverse(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4, 3.3}
	result := Reverse(arr)
	t.Log(result)
	if result[0] != 3.3 || len(result) != 5 {
		t.Fatal("fail")
	}
}

func TestSortedIndex(t *testing.T) {
	arr := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	result := SortedIndex(arr, 2.3)
	t.Log(result)
	if result != 2 {
		t.Fatal("fail")
	}
}

func TestSortedLastIndex(t *testing.T) {
	arr := []float64{1.1, 2.2, 3.3, 3.3, 5.5}
	result := SortedLastIndex(arr, 3.3)
	t.Log(result)
	if result != 4 {
		t.Fatal("fail")
	}
}

func TestSortedUniq(t *testing.T) {
	arr := []int{3, 2, 3, 3, 5}
	result := SortedUniq(arr)
	t.Log(result)
	if len(result) != 3 || result[0] != 2 {
		t.Fatal("fail")
	}
}

func TestSortedUniqBy(t *testing.T) {
	arr := []float64{3.3, 2.3, 3.3, 3.3, 5.3}
	result := SortedUniqBy(arr, math.Floor)
	t.Log(result)
	if len(result) != 3 || result[0] != 2.3 {
		t.Fatal("fail")
	}
}

func TestTail(t *testing.T) {
	arr := []float64{3.3, 2.3, 3.3, 3.3, 5.3}
	result := Tail(arr)
	t.Log(result)
	if len(result) != 4 || result[0] != 2.3 {
		t.Fatal("fail")
	}
}

func TestTake(t *testing.T) {
	arr := []float64{3.3, 2.3, 3.3, 3.3, 5.3}
	result := Take(arr, 2)
	t.Log(result)
	if len(result) != 2 || result[0] != 3.3 {
		t.Fatal("fail")
	}
}

func TestTakeRight(t *testing.T) {
	arr := []float64{3.3, 2.3, 3.3, 3.3, 5.3}
	result := TakeRight(arr, 2)
	t.Log(result)
	if len(result) != 2 || result[0] != 3.3 {
		t.Fatal("fail")
	}
}

func TestUnion(t *testing.T) {
	arr := []float64{3.3, 2.3, 3.3, 3.3, 5.3}
	arr1 := []float64{3.3, 2.3, 3.3, 3.3, 5.5}
	arr2 := []float64{1.1, 1.2, 6.6}
	result := Union(arr, arr1, arr2)
	t.Log(result)
	if len(result) != 7 || result[0] != 3.3 {
		t.Fatal("fail")
	}
}

func TestUniq(t *testing.T) {
	arr := []int{3, 2, 3, 3, 5}
	result := Uniq(arr)
	t.Log(result)
	if result[0] != 3 || len(result) != 3 {
		t.Fatal("fail")
	}
}

func TestUnionBy(t *testing.T) {
	arr := []float64{3.3, 2.3, 3.4, 3.5, 5.3}
	arr1 := []float64{1.1, 1.2, 3.4}
	arr2 := []float64{6.1, 6.2, 3.4}
	result := UnionBy(math.Floor, arr, arr1, arr2)
	t.Log(result)
	if result[0] != 3.3 || len(result) != 5 {
		t.Fatal("fail")
	}
}

func TestUniqBy(t *testing.T) {
	arr := []float64{3.3, 2.3, 3.4, 3.5, 5.3}
	result := UniqBy(arr, math.Floor)
	t.Log(result)
	if result[0] != 3.3 || len(result) != 3 {
		t.Fatal("fail")
	}
}

func TestUniqConvertBy(t *testing.T) {
	arr := []float64{3.3, 2.3, 3.4, 3.5, 5.3}
	result := UniqConvertBy(arr, math.Floor)
	t.Log(result)
	if result[0] != 3 || len(result) != 3 {
		t.Fatal("fail")
	}
}

func TestWithout(t *testing.T) {
	arr := []float64{3.3, 2.3, 3.4, 3.5, 5.3}
	excludes := []float64{3.3, 2.3}
	result := Without(arr, excludes)
	t.Log(result)
	if result[0] != 3.4 || len(result) != 3 {
		t.Fatal("fail")
	}
}

func TestXor(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4}
	arr1 := []float64{4.4}
	arr2 := []float64{4.3, 1.5}
	result := Xor(arr, arr1, arr2)
	t.Log(result)
	if len(result) != 5 {
		t.Fatal("fail")
	}
}

func TestXorBy(t *testing.T) {
	arr := []float64{2.2, 1.4, 3.3, 4.4}
	arr1 := []float64{4.4}
	arr2 := []float64{4.3, 1.5}
	result := XorBy(math.Floor, arr, arr1, arr2)
	t.Log(result)
	if len(result) != 2 {
		t.Fatal("fail")
	}
}
