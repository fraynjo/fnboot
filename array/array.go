package array

import (
	"github.com/fraynjo/fnboot"
	"github.com/fraynjo/fnboot/lang"
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
