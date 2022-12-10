package fnboot

func Chunk[T FnAny](array []T, size int) [][]T {
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

func Concat[T FnAny](array []T, values ...T) []T {
	newArr := array
	newArr = append(newArr, values...)
	return newArr
}

func Difference[T FnAny](array []T, values []T) []T {
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

func DifferenceBy[T FnAny, E FnAny](array []T, values []T, f func(x T) E) []E {
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
