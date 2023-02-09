package collection

import (
	"github.com/fraynjo/fnboot/array"
	"testing"
)

func TestForeachArray(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		arr := []int{1, 2, 3}
		result := ForeachArray(arr, func(i int, v int) float64 {
			return float64(v) + float64(1.1)
		})
		t.Log(result)
		if result[0] != 2.1 {
			t.Fatal("fail")
		}
	})
}

func TestForeachObj(t *testing.T) {
	t.Run("map<string, string>", func(t *testing.T) {
		obj := map[string]string{
			"name1": "foo1",
			"name2": "foo2",
		}
		result := ForeachObj[string](obj, func(k interface{}, v interface{}) interface{} {
			return v.(string) + "_" + "a"
		})
		t.Log(result)
		if result["name1"] != "foo1_a" {
			t.Fatal("fail")
		}
	})
}

func TestForeachRightArray(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		arr := []int{1, 2, 3}
		result := ForeachRightArray(arr, func(i int, v int) int {
			return v + 1
		})
		t.Log(result)
		if result[0] != 4 {
			t.Fatal("fail")
		}
	})
}

func TestEvery(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"user":   "barney",
			"age":    30,
			"active": false,
		}
		obj2 := map[string]interface{}{
			"user":   "fred",
			"age":    36,
			"active": false,
		}
		arr := []map[string]interface{}{obj1, obj2}
		result := Every(arr, func(index interface{}, item interface{}) bool {
			obj := ToMap[string](item)
			return !obj["active"].(bool)
		})
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
	t.Run("map", func(t *testing.T) {
		obj := map[string]interface{}{
			"user":   "barney",
			"age":    30,
			"active": false,
		}
		result := Every(obj, func(key interface{}, value interface{}) bool {
			if key == "age" {
				return value.(float64) == 30
			}
			return false
		})
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"user":   "barney",
			"age":    30,
			"active": true,
		}
		obj2 := map[string]interface{}{
			"user":   "fred",
			"age":    36,
			"active": true,
		}
		arr := map[string]interface{}{"obj1": obj1, "obj2": obj2}
		result := Filter(arr, func(index interface{}, item interface{}) bool {
			obj := ToMap[string](item)
			return obj["active"].(bool)
		})
		resultArr := array.ToArray(result)
		t.Log(resultArr)
		if len(resultArr) != 2 {
			t.Fatal("fail")
		}
	})
	t.Run("array", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"user":   "barney",
			"age":    30,
			"active": true,
		}
		obj2 := map[string]interface{}{
			"user":   "fred",
			"age":    36,
			"active": true,
		}
		arr := []map[string]interface{}{obj1, obj2}
		result := Filter(arr, func(index interface{}, item interface{}) bool {
			obj := ToMap[string](item)
			return obj["active"].(bool)
		})
		resultArr := array.ToArray(result)
		t.Log(resultArr)
		if len(resultArr) != 2 {
			t.Fatal("fail")
		}
	})
}

func TestFind(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"user":   "barney",
			"age":    30,
			"active": true,
		}
		obj2 := map[string]interface{}{
			"user":   "fred",
			"age":    36,
			"active": true,
		}
		arr := map[string]interface{}{"obj1": obj1, "obj2": obj2}
		result := Find(arr, func(index interface{}, item interface{}) bool {
			obj := ToMap[string](item)
			return obj["active"].(bool)
		})
		m := ToMap[string](result)
		t.Log(m)
		if m["user"] != "barney" {
			t.Fatal("fail")
		}
	})
	t.Run("array", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"user":   "barney",
			"age":    30,
			"active": true,
		}
		obj2 := map[string]interface{}{
			"user":   "fred",
			"age":    36,
			"active": true,
		}
		arr := []map[string]interface{}{obj1, obj2}
		result := Find(arr, func(index interface{}, item interface{}) bool {
			obj := ToMap[string](item)
			return obj["active"].(bool)
		})
		m := ToMap[string](result)
		t.Log(m)
		if m["user"] != "barney" {
			t.Fatal("fail")
		}
	})
}

func TestFindLast(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"user":   "barney",
			"age":    30,
			"active": true,
		}
		obj2 := map[string]interface{}{
			"user":   "fred",
			"age":    36,
			"active": true,
		}
		arr := map[string]interface{}{"obj1": obj1, "obj2": obj2}
		result := FindLast(arr, func(index interface{}, item interface{}) bool {
			obj := ToMap[string](item)
			return obj["active"].(bool)
		})
		m := ToMap[string](result)
		t.Log(m)
		if m["user"] != "fred" {
			t.Fatal("fail")
		}
	})
	t.Run("array", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"user":   "barney",
			"age":    30,
			"active": true,
		}
		obj2 := map[string]interface{}{
			"user":   "fred",
			"age":    36,
			"active": true,
		}
		arr := []map[string]interface{}{obj1, obj2}
		result := FindLast(arr, func(index interface{}, item interface{}) bool {
			obj := ToMap[string](item)
			return obj["active"].(bool)
		})
		m := ToMap[string](result)
		t.Log(m)
		if m["user"] != "fred" {
			t.Fatal("fail")
		}
	})
}

func TestIncludes(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"user":   "barney",
			"age":    30,
			"active": false,
		}
		obj2 := map[string]interface{}{
			"user":   "fred",
			"age":    36,
			"active": false,
		}
		arr := []map[string]interface{}{obj1, obj2}
		result := Includes(arr, func(index interface{}, item interface{}) bool {
			obj := ToMap[string](item)
			return !obj["active"].(bool)
		})
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
	t.Run("map", func(t *testing.T) {
		obj := map[string]interface{}{
			"user":   "barney",
			"age":    30,
			"active": false,
		}
		result := Includes(obj, func(key interface{}, value interface{}) bool {
			if key == "age" {
				return value.(float64) == 30
			}
			return false
		})
		t.Log(result)
		if !result {
			t.Fatal("fail")
		}
	})
}

func TestKeyBy(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"user":   "barney",
			"age":    30,
			"active": false,
		}
		obj2 := map[string]interface{}{
			"user":   "fred",
			"age":    36,
			"active": false,
		}
		arr := []map[string]interface{}{obj1, obj2}
		result := KeyBy(arr, func(index interface{}, item interface{}) string {
			m := ToMap[string](item)
			return m["user"].(string)
		})
		t.Log(result)
		if len(result) != 2 {
			t.Fatal("fail")
		}
	})
	t.Run("map", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"user":   "barney",
			"age":    30,
			"active": false,
		}
		obj2 := map[string]interface{}{
			"user":   "fred",
			"age":    36,
			"active": false,
		}
		obj := map[string]interface{}{
			"obj1": obj1,
			"obj2": obj2,
		}
		result := KeyBy(obj, func(key interface{}, value interface{}) string {
			m := ToMap[string](value)
			return m["user"].(string)
		})
		t.Log(result)
		if len(result) != 2 {
			t.Fatal("fail")
		}
	})
	t.Run("string", func(t *testing.T) {
		obj1 := map[string]interface{}{
			"user":   "barney",
			"age":    30,
			"active": false,
		}
		obj2 := map[string]interface{}{
			"user":   "fred",
			"age":    36,
			"active": false,
		}
		obj := map[string]interface{}{
			"obj1": obj1,
			"obj2": obj2,
		}
		result := KeyBy(obj, "user")
		t.Log(result)
		if len(result) != 2 {
			t.Fatal("fail")
		}
	})
}
