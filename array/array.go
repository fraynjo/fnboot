package array

import (
	"github.com/fraynjo/fnboot"
	"github.com/fraynjo/fnboot/lang"
	"reflect"
	"strconv"
	"strings"
)

func Chunk[T fnboot.FnAny](array []T, size int) [][]T {
	newArr := make([][]T, 0)
	tempArr := make([]T, 0)
	for k, v := range array {
		tempArr = append(tempArr, v)
		if (k+1)%size == 0 {
			newArr = append(newArr, tempArr)
			tempArr = make([]T, 0)
		}
	}
	if len(tempArr) > 0 {
		newArr = append(newArr, tempArr)
	}
	return newArr
}

func Compact(array []interface{}) []interface{} {
	newArr := make([]interface{}, 0)
	for _, v := range array {
		if v == 0 ||
			v == "" ||
			v == false ||
			v == nil {
			continue
		}
		newArr = append(newArr, v)
	}
	return newArr
}

func Concat[T fnboot.FnAny](array []T, values ...T) []T {
	newArr := array
	newArr = append(newArr, values...)
	return newArr
}

func Difference[T fnboot.FnAny](array []T, values []T) []T {
	newArr := make([]T, 0)
	valueMap := make(map[interface{}]bool)
	for _, v := range values {
		valueMap[v] = true
	}
	for _, v := range array {
		if _, ok := valueMap[v]; !ok {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

func DifferenceBy[T fnboot.FnAny, E fnboot.FnAny](array []T, values []T, f func(x T) E) []E {
	newArr := make([]E, 0)
	valueMap := make(map[interface{}]bool)
	for _, v := range values {
		valueMap[f(v)] = true
	}
	for _, v := range array {
		if _, ok := valueMap[f(v)]; !ok {
			newArr = append(newArr, f(v))
		}
	}
	return newArr
}

// func DifferenceWith[T Fn_T](array []T, values []T, comparator)

func Drop[T fnboot.FnAny](array []T, n int) []T {
	if len(array) < n {
		return make([]T, 0)
	}
	return array[n:]
}

func DropRight[T fnboot.FnAny](array []T, n int) []T {
	if len(array) < n {
		return make([]T, 0)
	}
	return array[:(len(array) - n)]
}

func Fill[T fnboot.FnAny](array []T, value T) []T {
	newArr := make([]T, len(array))
	for i := 0; i < len(array); i++ {
		newArr[i] = value
	}
	return newArr
}

func FindIndex[T fnboot.FnAny](array []T, value T) int {
	index := -1
	for i, v := range array {
		if lang.Eq(v, value) {
			index = i
			break
		}
	}
	return index
}

func FindLastIndex[T fnboot.FnAny](array []T, value T) int {
	index := -1
	for i := len(array) - 1; i >= 0; i-- {
		if lang.Eq(array[i], value) {
			index = i
			break
		}
	}
	return index
}

func First[T fnboot.FnAny](array []T) T {
	return array[0]
}

func IndexOf[T fnboot.FnAny](array []T, value T) int {
	for i, v := range array {
		if lang.Eq(v, value) {
			return i
		}
	}
	return -1
}

func Initial[T fnboot.FnAny](array []T) []T {
	size := len(array)
	if size <= 1 {
		return make([]T, 0)
	}
	return array[:(size - 1)]
}

func Intersection[T fnboot.FnAny](arraies ...[]T) []T {
	resultArray := make([]T, 0)
	if len(arraies) == 0 {
		return resultArray
	} else if len(arraies) == 1 {
		return Uniq(arraies[0])
	}
	shortIndex := 0
	tempShortArrayLen := len(arraies[0])
	for i, array := range arraies {
		if len(array) < tempShortArrayLen {
			tempShortArrayLen = len(array)
			shortIndex = i
		}
	}
	for _, v := range arraies[shortIndex] {
		count := 0
		for i, array := range arraies {
			if i == shortIndex || IndexOf(array, v) == -1 {
				continue
			}
			count += 1
		}
		if count == len(arraies)-1 {
			resultArray = append(resultArray, v)
		}
	}
	return resultArray
}

func IntersectionBy[T fnboot.FnAny, T1 fnboot.FnAny](f func(v T) T1, arraies ...[]T) []T {
	resultArray := make([]T, 0)
	if len(arraies) == 0 {
		return resultArray
	} else if len(arraies) == 1 {
		return UniqBy(arraies[0], f)
	}
	shortIndex := 0
	tempShortArrayLen := len(arraies[0])
	convertedArraies := make([][]T1, len(arraies))
	for i, array := range arraies {
		convertedArraies[i] = UniqConvertBy(array, f)
		if len(array) < tempShortArrayLen {
			tempShortArrayLen = len(array)
			shortIndex = i
		}
	}
	for _, v := range arraies[shortIndex] {
		count := 0
		for i, array := range convertedArraies {
			if i == shortIndex || IndexOf(array, f(v)) == -1 {
				continue
			}
			count += 1
		}
		if count == len(arraies)-1 {
			resultArray = append(resultArray, v)
		}
	}
	return resultArray
}

func ToArray(array interface{}) []interface{} {
	size := reflect.ValueOf(array).Len()
	if size == 0 {
		return make([]interface{}, 0)
	}
	out := make([]interface{}, 0)
	for i := 0; i < size; i++ {
		out = append(out, reflect.ValueOf(array).Index(i).Interface())
	}
	return out
}

func Join(array interface{}, sep string) string {
	if !lang.IsArray(array) {
		return ""
	}
	arr := ToArray(array)
	if len(arr) == 0 {
		return ""
	}
	if !lang.IsNumber(arr[0]) && !lang.IsString(arr[0]) {
		return ""
	}
	strArr := make([]string, 0)
	for _, v := range arr {
		s := ""
		if lang.IsFloat(v) {
			s = strconv.FormatFloat(v.(float64), 'f', -1, 64)
		} else if lang.IsInteger(v) {
			s = strconv.Itoa(v.(int))
		} else if lang.IsString(v) {
			s = v.(string)
		} else {
			continue
		}
		strArr = append(strArr, s)
	}
	return strings.Join(strArr, sep)
}

func Last[T fnboot.FnAny](array []T) T {
	return array[len(array)-1]
}

func LastIndexOf[T fnboot.FnAny](array []T, value T) int {
	for i := len(array) - 1; i >= 0; i-- {
		if lang.Eq(array[i], value) {
			return i
		}
	}
	return -1
}

func Nth[T fnboot.FnAny](array []T, n int) T {
	if n >= 0 {
		return array[n]
	} else {
		return array[len(array)+n-1]
	}
}

func Pull[T fnboot.FnAny](array []T, value T) []T {
	tempArr := make([]T, 0)
	for _, v := range array {
		if !lang.Eq(v, value) {
			tempArr = append(tempArr, v)
		}
	}
	return tempArr
}

func PullAll[T fnboot.FnAny](array []T, values []T) []T {
	tempArr := make([]T, 0)
	for _, v := range array {
		if IndexOf(values, v) == -1 {
			tempArr = append(tempArr, v)
		}
	}
	return tempArr
}

func PullAt[T fnboot.FnAny](array []T, indexs []int) []T {
	tempArr := make([]T, 0)
	for i, v := range array {
		if IndexOf(indexs, i) == -1 {
			tempArr = append(tempArr, v)
		}
	}
	return tempArr
}

func RemoveBy[T fnboot.FnAny](array []T, by interface{}) []T {
	tempArr := make([]T, 0)
	var ft func(v T) bool
	if reflect.TypeOf(by) == reflect.TypeOf(ft) {
		f := by.(func(v T) bool)
		for _, v := range array {
			if !f(v) {
				tempArr = append(tempArr, v)
			}
		}
	}
	if lang.IsInteger(by) {
		index := by.(int)
		return append(array[:index], array[index+1:]...)
	}
	return tempArr
}

func Reverse[T fnboot.FnAny](array []T) []T {
	tempArr := make([]T, len(array))
	for i, v := range array {
		tempArr[len(array)-i-1] = v
	}
	return tempArr
}

func SortedIndex[T fnboot.FnNumber](array []T, value T) int {
	for i, v := range array {
		if v >= value {
			return i
		}
	}
	return -1
}

func SortedLastIndex[T fnboot.FnNumber](array []T, value T) int {
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] >= value {
			return i
		}
	}
	return -1
}

func SortedUniq[T fnboot.FnNumber](array []T) []T {
	uniqArr := Uniq(array)
	for i := 0; i < len(uniqArr); i++ {
		for j := 0; j < len(uniqArr)-i-1; j++ {
			if float64(uniqArr[j]) > float64(uniqArr[j+1]) {
				uniqArr[j], uniqArr[j+1] = uniqArr[j+1], uniqArr[j]
			}
		}
	}
	return uniqArr
}

func SortedUniqBy[T fnboot.FnNumber](array []T, f func(v T) T) []T {
	uniqArr := Uniq(array)
	for i := 0; i < len(uniqArr); i++ {
		for j := 0; j < len(uniqArr)-i-1; j++ {
			if f(uniqArr[j]) >= f(uniqArr[j+1]) {
				uniqArr[j], uniqArr[j+1] = uniqArr[j+1], uniqArr[j]
			}
		}
	}
	return uniqArr
}

func Tail[T fnboot.FnAny](array []T) []T {
	return array[1:]
}

func Take[T fnboot.FnAny](array []T, n int) []T {
	return array[:n]
}

func TakeRight[T fnboot.FnAny](array []T, n int) []T {
	return array[len(array)-n:]
}

func Union[T fnboot.FnAny](arraies ...[]T) []T {
	tempArr := make([]T, 0)
	for _, array := range arraies {
		for _, v := range array {
			if IndexOf(tempArr, v) == -1 {
				tempArr = append(tempArr, v)
			}
		}
	}
	return tempArr
}

func UnionBy[T fnboot.FnAny](f func(v T) T, arraies ...[]T) []T {
	if len(arraies) == 0 {
		return make([]T, 0)
	} else if len(arraies) <= 1 {
		return arraies[0]
	}
	tempArr := make([]T, 0)
	uniqArr := make([]T, 0)
	for i := 0; i < len(arraies); i++ {
		for _, v := range arraies[i] {
			if IndexOf(uniqArr, f(v)) == -1 {
				uniqArr = append(uniqArr, f(v))
				tempArr = append(tempArr, v)
			}
		}
	}
	return tempArr
}

func Uniq[T fnboot.FnAny](array []T) []T {
	tempArr := make([]T, 0)
	for _, v := range array {
		if IndexOf(tempArr, v) == -1 {
			tempArr = append(tempArr, v)
		}
	}
	return tempArr
}

func UniqBy[T fnboot.FnAny, T1 fnboot.FnAny](array []T, f func(v T) T1) []T {
	tempArr := make([]T, 0)
	uniqArr := make([]T1, 0)
	for _, v := range array {
		if IndexOf(uniqArr, f(v)) == -1 {
			uniqArr = append(uniqArr, f(v))
			tempArr = append(tempArr, v)
		}
	}
	return tempArr
}

func UniqConvertBy[T fnboot.FnAny, T1 fnboot.FnAny](array []T, f func(v T) T1) []T1 {
	uniqArr := make([]T1, 0)
	for _, v := range array {
		if IndexOf(uniqArr, f(v)) == -1 {
			uniqArr = append(uniqArr, f(v))
		}
	}
	return uniqArr
}

func Without[T fnboot.FnAny](array []T, excludes []T) []T {
	tempArr := make([]T, 0)
	for _, v := range array {
		if IndexOf(excludes, v) == -1 {
			tempArr = append(tempArr, v)
		}
	}
	return tempArr
}

func Xor[T fnboot.FnAny](arraies ...[]T) []T {
	resultArray := make([]T, 0)
	if len(arraies) == 0 {
		return resultArray
	} else if len(arraies) == 1 {
		return Uniq(arraies[0])
	}
	for i1, array1 := range arraies {
		for _, v := range array1 {
			count := 0
			for i2, array2 := range arraies {
				if i2 == i1 || IndexOf(array2, v) >= 0 {
					continue
				}
				count += 1
			}
			if count == len(arraies)-1 {
				resultArray = append(resultArray, v)
			}
		}
	}
	return resultArray
}

func XorBy[T fnboot.FnAny, T1 fnboot.FnAny](f func(v T) T1, arraies ...[]T) []T {
	resultArray := make([]T, 0)
	if len(arraies) == 0 {
		return resultArray
	} else if len(arraies) == 1 {
		return Uniq(arraies[0])
	}
	convertedArraies := make([][]T1, len(arraies))
	for i, array := range arraies {
		convertedArraies[i] = UniqConvertBy(array, f)
	}
	for i1, array1 := range arraies {
		for _, v := range array1 {
			count := 0
			for i2, array2 := range convertedArraies {
				if i2 == i1 || IndexOf(array2, f(v)) >= 0 {
					continue
				}
				count += 1
			}
			if count == len(arraies)-1 {
				resultArray = append(resultArray, v)
			}
		}
	}
	return resultArray
}
