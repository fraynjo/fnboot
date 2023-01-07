package collection

import (
	"encoding/json"
	"github.com/fraynjo/fnboot"
	"github.com/fraynjo/fnboot/array"
	"github.com/fraynjo/fnboot/lang"
)

func ToMap[T fnboot.FnT](obj interface{}) map[T]interface{} {
	var m map[T]interface{}
	bs, _ := json.Marshal(obj)
	_ = json.Unmarshal(bs, &m)
	return m
}

func ForeachArray[T fnboot.FnInterface, T1 fnboot.FnInterface](arr []T, f func(index int, item T) T1) []T1 {
	tempArr := make([]T1, len(arr))
	for i, v := range arr {
		tempArr[i] = f(i, v)
	}
	return tempArr
}

func ForeachObj[T fnboot.FnT](obj interface{}, f func(key interface{}, value interface{}) interface{}) map[T]interface{} {
	m := ToMap[T](obj)
	for k, v := range m {
		m[k] = f(k, v)
	}
	return m
}

func ForeachRightArray[T fnboot.FnInterface](arr []T, f func(index int, item T) T) []T {
	tempArr := make([]T, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		tempArr[len(arr)-i-1] = f(i, arr[i])
	}
	return tempArr
}

func Every(obj interface{}, f func(key interface{}, item interface{}) bool) bool {
	if lang.IsArray(obj) {
		arr := array.InterfaceToArray(obj)
		return everyArray(arr, f)
	}
	if lang.IsMap(obj) {
		return everyObj(obj, f)
	}
	return false
}

func everyArray(arr []interface{}, f func(key interface{}, item interface{}) bool) bool {
	for i, v := range arr {
		if !f(i, v) {
			return false
		}
	}
	return true
}

func everyObj(obj interface{}, f func(key interface{}, value interface{}) bool) bool {
	var m map[string]interface{}
	bs, _ := json.Marshal(obj)
	_ = json.Unmarshal(bs, &m)
	for k, v := range m {
		if f(k, v) {
			return true
		}
	}
	return false
}

func Filter(obj interface{}, f func(key interface{}, item interface{}) bool) interface{} {
	if lang.IsArray(obj) {
		arr := array.InterfaceToArray(obj)
		return everyArray(arr, f)
	}
	if lang.IsMap(obj) {
		return everyObj(obj, f)
	}
	return nil
}

func filterArray(arr []interface{}, f func(key interface{}, item interface{}) bool) interface{} {
	for i, v := range arr {
		if f(i, v) {
			return v
		}
	}
	return nil
}

func filterObj(obj interface{}, f func(key interface{}, value interface{}) bool) interface{} {
	var m map[string]interface{}
	bs, _ := json.Marshal(obj)
	_ = json.Unmarshal(bs, &m)
	for k, v := range m {
		if f(k, v) {
			return v
		}
	}
	return nil
}
