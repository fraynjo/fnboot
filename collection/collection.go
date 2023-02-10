package collection

import (
	"encoding/json"
	"github.com/fraynjo/fnboot"
	"github.com/fraynjo/fnboot/array"
	"github.com/fraynjo/fnboot/lang"
	"github.com/fraynjo/fnboot/number"
	"reflect"
	"sort"
	"strings"
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
		arr := array.ToArray(obj)
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
	m := ToMap[string](obj)
	for k, v := range m {
		if !f(k, v) {
			return false
		}
	}
	return true
}

func Filter(obj interface{}, f func(key interface{}, item interface{}) bool) []interface{} {
	if lang.IsArray(obj) {
		arr := array.ToArray(obj)
		return filterArray(arr, f)
	}
	if lang.IsMap(obj) {
		return filterObj(obj, f)
	}
	return nil
}

func filterArray(arr interface{}, f func(key interface{}, item interface{}) bool) []interface{} {
	out := make([]interface{}, 0)
	newArr := array.ToArray(arr)
	for i, v := range newArr {
		if f(i, v) {
			out = append(out, v)
		}
	}
	return out
}

func filterObj(obj interface{}, f func(key interface{}, value interface{}) bool) []interface{} {
	out := make([]interface{}, 0)
	m := ToMap[string](obj)
	for k, v := range m {
		if f(k, v) {
			out = append(out, v)
		}
	}
	return out
}

func Find(obj interface{}, f func(key interface{}, item interface{}) bool) interface{} {
	if lang.IsArray(obj) {
		arr := array.ToArray(obj)
		return findArray(arr, f)
	}
	if lang.IsMap(obj) {
		return findObj(obj, f)
	}
	return nil
}

func findArray(arr interface{}, f func(key interface{}, item interface{}) bool) interface{} {
	newArr := array.ToArray(arr)
	for i, v := range newArr {
		if f(i, v) {
			return v
		}
	}
	return nil
}

func findObj(obj interface{}, f func(key interface{}, value interface{}) bool) interface{} {
	m := ToMap[string](obj)
	for k, v := range m {
		if f(k, v) {
			return v
		}
	}
	return nil
}

func FindLast(obj interface{}, f func(key interface{}, item interface{}) bool) interface{} {
	if lang.IsArray(obj) {
		arr := array.ToArray(obj)
		return findLastArray(arr, f)
	}
	if lang.IsMap(obj) {
		return findLastObj(obj, f)
	}
	return nil
}

func findLastArray(arr interface{}, f func(key interface{}, item interface{}) bool) interface{} {
	newArr := array.ToArray(arr)
	for i := len(newArr) - 1; i >= 0; i-- {
		if f(i, newArr[i]) {
			return newArr[i]
		}
	}
	return nil
}

func findLastObj(obj interface{}, f func(key interface{}, value interface{}) bool) interface{} {
	filterObjs := filterObj(obj, f)
	if len(filterObjs) >= 0 {
		return array.Last(filterObjs)
	}
	return nil
}

func Includes(obj interface{}, f func(key interface{}, item interface{}) bool) bool {
	if lang.IsArray(obj) {
		arr := array.ToArray(obj)
		return includeArray(arr, f)
	}
	if lang.IsMap(obj) {
		return includeObj(obj, f)
	}
	return false
}

func includeArray(arr []interface{}, f func(key interface{}, item interface{}) bool) bool {
	for i, v := range arr {
		if f(i, v) {
			return true
		}
	}
	return false
}

func includeObj(obj interface{}, f func(key interface{}, value interface{}) bool) bool {
	m := ToMap[string](obj)
	for k, v := range m {
		if f(k, v) {
			return true
		}
	}
	return false
}


func KeyBy(obj interface{}, by interface{}) (m map[string]interface{}) {
	if lang.IsArray(obj) {
		arr := array.ToArray(obj)
		return keyByArray(arr, by)
	}
	if lang.IsMap(obj) {
		return keyByObj(obj, by)
	}
	return nil
}

func keyByArray(arr []interface{}, by interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	var f func(key interface{}, v interface{}) string
	if reflect.TypeOf(by) == reflect.TypeOf(f) {
		f := by.(func(key interface{}, item interface{}) string)
		for i, v := range arr {
			m[f(i, v)] = v
		}
	} else if lang.IsString(by) {
		keyStr := by.(string)
		for _, v := range m {
			vMap := ToMap[string](v)
			if key, ok := vMap[keyStr]; ok {
				m[key.(string)] = v
			}
		}
	}
	return m
}

func keyByObj(obj interface{}, by interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	objMap := ToMap[string](obj)
	var f func(key interface{}, v interface{}) string
	if reflect.TypeOf(by) == reflect.TypeOf(f) {
		f := by.(func(key interface{}, item interface{}) string)
		for k, v := range objMap {
			m[f(k, v)] = v
		}
	} else if lang.IsString(by) {
		keyStr := by.(string)
		for _, v := range objMap {
			vMap := ToMap[string](v)
			if key, ok := vMap[keyStr]; ok {
				m[key.(string)] = v
			}
		}
	}
	return m
}

func OrderBy(obj interface{}, field string, direction string) []interface{} {
	arr := array.ToArray(obj)
	m := make(map[interface{}]interface{})
	var ktype string
	strArr := make([]string, 0)
	numArr := make([]float64, 0)
	for i, v := range arr {
		vm := ToMap[string](v)
		if lang.IsString(vm[field]) {
			ktype = "string"
			strArr = append(strArr, vm[field].(string))
		} else if lang.IsNumber(vm[field]) {
			ktype = "number"
			numArr = append(numArr, vm[field].(float64))
		} else {
			break
		}
		m[vm[field]] = i
	}
	newArr := make([]interface{}, 0)
	if ktype == "string" {
		sort.Strings(strArr)
		if strings.ToLower(direction) == "desc" {
			strArr = array.Reverse[string](strArr)
		}
		for _, v := range strArr {
			newArr = append(newArr, m[v])
		}
	}
	if ktype == "number" {
		sort.Float64s(numArr)
		if strings.ToLower(direction) == "desc" {
			numArr = array.Reverse[float64](numArr)
		}
		for _, v := range numArr {
			newArr = append(newArr, arr[m[v].(int)])
		}
	}
	return newArr
}

func ReduceArray[T fnboot.FnAny, T1 fnboot.FnAny](arr []T1, f func(result T, item T1) T) interface{} {
	var result T
	for _, v := range arr {
		result = f(result, v)
	}
	return result
}

func ReduceArrayRight[T fnboot.FnAny, T1 fnboot.FnAny](arr []T1, f func(result T, item T1) T) interface{} {
	var result T
	for i:=len(arr)-1; i>=0; i-- {
		result = f(result, arr[i])
	}
	return result
}

func Some(obj interface{}, f func(key interface{}, item interface{}) bool) bool {
	if lang.IsArray(obj) {
		arr := array.ToArray(obj)
		return someArray(arr, f)
	}
	if lang.IsMap(obj) {
		return someObj(obj, f)
	}
	return false
}

func someArray(arr []interface{}, f func(key interface{}, item interface{}) bool) bool {
	for i, v := range arr {
		if f(i, v) {
			return true
		}
	}
	return false
}

func someObj(obj interface{}, f func(key interface{}, value interface{}) bool) bool {
	m := ToMap[string](obj)
	for k, v := range m {
		if f(k, v) {
			return true
		}
	}
	return false
}

func Sample[T fnboot.FnAny](arr []T) T {
	randInt := number.Random(0, len(arr))
	return arr[randInt]
}

func SampleSize[T fnboot.FnAny](arr []T, size int) []T {
	result := make([]T, 0)
	for i:=1; i<=size; i++ {
		index := number.Random(0, len(arr))
		result = append(result, arr[index])
		arr = array.RemoveBy(arr, index)
	}
	return result
}
