package fnboot

import "testing"

func TestForeachArray(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		arr := []int{1, 2, 3}
		result := ForeachArray(arr, func(i int, v int) int {
			return v + 1
		})
		t.Log(result)
		if result[0] != 2 {
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
		result := ForeachObj(obj, func(k string, v string) string {
			return v + "_" + "a"
		})
		t.Log(result)
		if result["name1"] != "foo1_a" {
			t.Fatal("fail")
		}
	})
}
