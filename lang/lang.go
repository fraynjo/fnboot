package lang

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fraynjo/fnboot"
	"reflect"
	"strconv"
	"time"
)

func CastArray[T fnboot.FnAny](value T) []T {
	return []T{value}
}

func ConformsTo[T fnboot.FnInterface, T1, T2 fnboot.FnT](object T, source map[T1]func(T2) bool) bool {
	var m map[T1]T2
	b, _ := json.Marshal(object)
	_ = json.Unmarshal(b, &m)
	result := true
	for k, f := range source {
		if v, ok := m[k]; ok {
			if !f(v) {
				result = false
				break
			}
		} else {
			result = false
			break
		}
	}
	return result
}

func Eq(value any, other any) bool {
	return reflect.DeepEqual(value, other)
}

func Gt[T fnboot.FnNumber, T1 fnboot.FnNumber](value T, other T1) bool {
	return float64(value) > float64(other)
}

func Gte[T fnboot.FnNumber, T1 fnboot.FnNumber](value T, other T1) bool {
	return float64(value) >= float64(other)
}

func IsArray(value interface{}) bool {
	t := reflect.TypeOf(value).Kind()
	return t == reflect.Array || t == reflect.Slice
}

func IsBool[T fnboot.FnInterface](value T) bool {
	t := reflect.TypeOf(value).Kind()
	return t == reflect.Bool
}

func IsBuffer[T fnboot.FnInterface](value T) bool {
	return reflect.TypeOf(value) == reflect.TypeOf(&bytes.Buffer{})
}

func IsTime[T fnboot.FnInterface](value T) bool {
	return reflect.TypeOf(value) == reflect.TypeOf(time.Now())
}

func IsEmpty[T fnboot.FnInterface](value T) bool {
	val := reflect.ValueOf(value)
	tKind := reflect.TypeOf(value).Kind()
	if tKind == reflect.Slice || tKind == reflect.Map || tKind == reflect.Array {
		return val.Len() == 0
	}
	return val.IsZero()
}

func IsError[T fnboot.FnInterface](value T) bool {
	return reflect.TypeOf(value) == reflect.TypeOf(errors.New(""))
}

func IsInteger[T fnboot.FnInterface](value T) bool {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64, reflect.Uint8, reflect.Uint16,
		reflect.Uint, reflect.Uint32, reflect.Uint64:
		return true
	}
	return false
}

func IsMap(value interface{}) bool {
	t := reflect.TypeOf(value)
	return t.Kind() == reflect.Map
}

//func IsMatch[T1 FnInterface, T2 FnInterface](object T1, source T2) bool {
//	var m map[interface{}]interface{}
//	b, _ := json.Marshal(object)
//	_ = json.Unmarshal(b, &m)
//	var sm map[FnInterface]FnInterface
//	sb, _ := json.Marshal(source)
//	_ = json.Unmarshal(sb, &sm)
//	for k, v := range sm {
//		if vi, ok := m[k]; ok && vi == v {
//			return true
//		}
//	}
//	return false
//}

func IsNil(object fnboot.FnInterface) bool {
	return object == nil
}

func IsNumber(value fnboot.FnInterface) bool {
	return IsInteger(value) || IsFloat(value)
}

func IsString(value fnboot.FnInterface) bool {
	t := reflect.TypeOf(value).Kind()
	return t == reflect.String
}

func IsFloat(value fnboot.FnInterface) bool {
	t := reflect.TypeOf(value).Kind()
	return t == reflect.Float32 || t == reflect.Float64
}

func Lt[T fnboot.FnNumber, T1 fnboot.FnNumber](value T, other T1) bool {
	return float64(value) < float64(other)
}

func Lte[T fnboot.FnNumber, T1 fnboot.FnNumber](value T, other T1) bool {
	return float64(value) <= float64(other)
}

func ToInteger(value fnboot.FnInterface) int64 {
	if IsInteger(value) {
		return int64(value.(int))
	}
	if IsFloat(value) {
		si := fmt.Sprintf("%f", value)
		f64, _ := strconv.ParseFloat(si, 64)
		return int64(f64)
	}
	if IsString(value) {
		f64, _ := strconv.ParseFloat(value.(string), 64)
		return int64(f64)
	}
	return int64(0)
}

func ToFloat(value fnboot.FnInterface) float64 {
	if IsInteger(value) {
		return float64(value.(int))
	}
	if IsFloat(value) {
		return value.(float64)
	}
	if IsString(value) {
		f64, _ := strconv.ParseFloat(value.(string), 64)
		return f64
	}
	return float64(0)
}

func ToString(value fnboot.FnInterface) string {
	if IsString(value) {
		return value.(string)
	}
	if IsInteger(value) {
		return strconv.Itoa(value.(int))
	}
	if IsFloat(value) {
		return strconv.FormatFloat(value.(float64), 'g', -1, 64)
	}

	return ""
}
