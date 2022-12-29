package math

import (
	"fmt"
	"github.com/fraynjo/fnboot"
	"math"
	"strconv"
)

func Add[T fnboot.FnNumber, T1 fnboot.FnNumber](augend T, addend T1) float64 {
	return float64(augend) + float64(addend)
}

func Ceil[T fnboot.FnNumber](value T) float64 {
	return math.Ceil(float64(value))
}

func Divide[T fnboot.FnNumber, T1 fnboot.FnNumber](dividend T, divisor T1) float64 {
	return float64(dividend) / float64(divisor)
}

func Floor[T fnboot.FnNumber](value T) float64 {
	return math.Floor(float64(value))
}

func Max[T fnboot.FnNumber](arr []T) float64 {
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

func Mean[T fnboot.FnNumber](arr []T) float64 {
	sum := float64(0)
	for _, v := range arr {
		sum = Add(sum, float64(v))
	}
	return Divide(sum, len(arr))
}

func Min[T fnboot.FnNumber](arr []T) float64 {
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

func Multiply[T fnboot.FnNumber, T1 fnboot.FnNumber](augend T, addend T1) float64 {
	return float64(augend) * float64(addend)
}

func Round[T fnboot.FnNumber](value T, precision int) float64 {
	formater := "%." + strconv.Itoa(precision) + "f"
	v, _ := strconv.ParseFloat(fmt.Sprintf(formater, float64(value)), 64)
	return v
}

func Subtract[T fnboot.FnNumber, T1 fnboot.FnNumber](minuend T, subtrahend T1) float64 {
	return float64(minuend) - float64(subtrahend)
}

func Sum[T fnboot.FnNumber](arr []T) float64 {
	sum := float64(0)
	for _, v := range arr {
		sum = Add(sum, float64(v))
	}
	return sum
}
