package lang

import (
	"bytes"
	"errors"
	"testing"
	"time"
)

func TestCastArray(t *testing.T) {
	t.Run("[]int", func(t *testing.T) {
		arr := []int{1}
		result := CastArray(arr)
		t.Log(result)
		if result[0][0] != 1 {
			t.Fatal("fail")
		}
	})
	t.Run("map", func(t *testing.T) {
		arr := map[string]interface{}{"a": 1}
		result := CastArray(arr)
		t.Log(result)
		if len(result) != 1 {
			t.Fatal("fail")
		}
	})
	t.Run("empty", func(t *testing.T) {
		result := CastArray("")
		t.Log(result)
		if len(result) != 1 {
			t.Fatal("fail")
		}
	})
}

func TestConformsTo(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		obj := map[string]int{
			"a": 1,
			"b": 2,
		}
		result := ConformsTo(obj, map[string]func(int) bool{
			"b": func(i int) bool {
				return i > 1
			},
		})
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
	t.Run("struct", func(t *testing.T) {
		type Temp struct {
			A int
			B int
		}
		obj := Temp{
			A: 1,
			B: 2,
		}
		result := ConformsTo(obj, map[string]func(int) bool{
			"A": func(i int) bool {
				return i > 1
			},
		})
		t.Log(result)
		if result {
			t.Fatal("fail")
		}
	})
}

func TestEq(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		result := Eq(1, 1)
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
	t.Run("[]", func(t *testing.T) {
		result := Eq([]int64{1}, []int64{1})
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
	t.Run("map", func(t *testing.T) {
		result := Eq(map[string]interface{}{"a": 1, "b": 2}, map[string]interface{}{"b": 2, "a": 1})
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
}

func TestGt(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		result := Gt(2, 1)
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})

	t.Run("int", func(t *testing.T) {
		result := Gt(1, 3.0)
		t.Log(result)
		if result {
			t.Fatal("fail")
		}
	})
}

func TestGte(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		result := Gte(2, 1)
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})

	t.Run("int", func(t *testing.T) {
		result := Gte(3, 3.0)
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
}

func TestIsArray(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		result := IsArray([]int{1})
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})

	t.Run("map", func(t *testing.T) {
		result := IsArray(map[string]interface{}{"a": 1})
		t.Log(result)
		if result {
			t.Fatal("fail")
		}
	})
}

func TestIsBool(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		result := IsBool(1)
		t.Log(result)
		if result {
			t.Fatal("fail")
		}
	})

	t.Run("bool", func(t *testing.T) {
		result := IsBool(true)
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
}

func TestIsBuffer(t *testing.T) {
	t.Run("bytes", func(t *testing.T) {
		result := IsBuffer(bytes.NewBufferString("ddd"))
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})

	t.Run("other", func(t *testing.T) {
		result := IsBuffer(true)
		t.Log(result)
		if result {
			t.Fatal("fail")
		}
	})
}

func TestIsTime(t *testing.T) {
	t.Run("Date", func(t *testing.T) {
		result := IsTime(time.Now())
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})

	t.Run("other", func(t *testing.T) {
		result := IsTime("2022")
		t.Log(result)
		if result {
			t.Fatal("fail")
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		result := IsEmpty(0)
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})

	t.Run("[]", func(t *testing.T) {
		result := IsEmpty([]int{})
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})

	t.Run("map", func(t *testing.T) {
		result := IsEmpty(map[string]interface{}{})
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
}

func TestIsError(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		result := IsError(errors.New("1"))
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
}

func TestIsInteger(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		result := IsInteger(int8(1))
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
}

func TestIsMap(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		m := map[string]interface{}{}
		result := IsMap(m)
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
	t.Run("mapPtr", func(t *testing.T) {
		m := &map[string]interface{}{}
		result := IsMap(*m)
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
}

//func TestIsMatch(t *testing.T) {
//	t.Run("map", func(t *testing.T) {
//		m1 := map[string]int{"a": 1, "b": 2}
//		result := IsMatch(m1, map[string]int{"b": 2})
//		t.Log(result)
//		if !result {
//			t.Fatal("fail")
//		}
//	})
//}

func TestIsNil(t *testing.T) {
	t.Run("integer", func(t *testing.T) {
		i := 1
		result := IsNil(i)
		t.Log(result)
		if result {
			t.Fatal("fail")
		}
	})
}

func TestIsNumber(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		v := "foo"
		result := IsNumber(v)
		t.Log(result)
		if result {
			t.Fatal("fail")
		}
	})
	t.Run("float", func(t *testing.T) {
		v := 3.3
		result := IsNumber(v)
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
}

func TestIsFloat(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		v := "foo"
		result := IsFloat(v)
		t.Log(result)
		if result {
			t.Fatal("fail")
		}
	})
	t.Run("float", func(t *testing.T) {
		v := 3.3
		result := IsFloat(v)
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
}

func TestIsString(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		v := "foo"
		result := IsString(v)
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
	t.Run("float", func(t *testing.T) {
		v := 3.3
		result := IsString(v)
		t.Log(result)
		if result {
			t.Fatal("fail")
		}
	})
}

func TestLt(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		result := Lt(2, 1)
		t.Log(result)
		if result {
			t.Fatal("fail")
		}
	})

	t.Run("float", func(t *testing.T) {
		result := Lt(1, 3.3)
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
}

func TestLte(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		result := Lte(2, 1)
		t.Log(result)
		if result {
			t.Fatal("fail")
		}
	})

	t.Run("float", func(t *testing.T) {
		result := Lte(3, 3.0)
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
}

func TestToInteger(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		result := ToInteger(2.2)
		t.Log(result)
		if result != int64(2) {
			t.Fatal("fail")
		}
	})

	t.Run("string", func(t *testing.T) {
		result := ToInteger("2")
		t.Log(result)
		if result != int64(2) {
			t.Fatal("fail")
		}
	})
}

func TestToFloat(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		result := ToFloat(2)
		t.Log(result)
		if result != float64(2) {
			t.Fatal("fail")
		}
	})

	t.Run("string", func(t *testing.T) {
		result := ToFloat("2.0")
		t.Log(result)
		if result != float64(2.0) {
			t.Fatal("fail")
		}
	})
}

func TestToString(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		result := ToString(2)
		t.Log(result)
		if result != "2" {
			t.Fatal("fail")
		}
	})

	t.Run("float", func(t *testing.T) {
		result := ToString("2.0")
		t.Log(result)
		if result != "2.0" {
			t.Fatal("fail")
		}
	})
}
