package fnboot

import (
	"encoding/json"
)

func ForeachArray[T FnInterface](arr []T, f func(i int, v T) T) []T {
	for i, v := range arr {
		arr[i] = f(i, v)
	}
	return arr
}

func ForeachObj[T FnInterface](obj FnInterface, f func(k string, v T) T) map[string]T {
	var m map[string]T
	bs, _ := json.Marshal(obj)
	_ = json.Unmarshal(bs, &m)
	for k, v := range m {
		m[k] = f(k, v)
	}
	return m
}
