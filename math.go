package fnboot

import (
	"fmt"
	"math"
	"strconv"
)

func Add[T FnNumber, T1 FnNumber](augend T, addend T1) float64 {
	return float64(augend) + float64(addend)
}

func Ceil[T FnNumber](value T) float64 {
	return math.Ceil(float64(value))
}

func Divide[T FnNumber, T1 FnNumber](dividend T, divisor T1) float64 {
	return float64(dividend) / float64(divisor)
}

func Floor[T FnNumber](value T) float64 {
	return math.Floor(float64(value))
}

func Max[T FnNumber](arr []T) float64 {
	max := float64(0)
	for i, v := range arr {
		if i == 0 {
			max = float64(v)
		} else {
			max = math.Max(max, float64(v))
		}
	}
	return max
}

func Mean[T FnNumber](arr []T) float64 {
	sum := float64(0)
	for _, v := range arr {
		sum = Add(sum, float64(v))
	}
	return Divide(sum, len(arr))
}

func Min[T FnNumber](arr []T) float64 {
	min := float64(0)
	for i, v := range arr {
		if i == 0 {
			min = float64(v)
		} else {
			min = math.Min(min, float64(v))
		}
	}
	return min
}

func Multiply[T FnNumber, T1 FnNumber](augend T, addend T1) float64 {
	return float64(augend) * float64(addend)
}

func Round[T FnNumber](value T, precision int) float64 {
	formater := "%." + strconv.Itoa(precision) + "f"
	v, _ := strconv.ParseFloat(fmt.Sprintf(formater, float64(value)), 64)
	return v
}

func Subtract[T FnNumber, T1 FnNumber](minuend T, subtrahend T1) float64 {
	return float64(minuend) - float64(subtrahend)
}

func Sum[T FnNumber](arr []T) float64 {
	sum := float64(0)
	for _, v := range arr {
		sum = Add(sum, float64(v))
	}
	return sum
}
