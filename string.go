package fnboot

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
)

type FString struct {
}

func (fs FString) CamelCase(value, sep string) string {
	arr := strings.Split(value, sep)
	s := ""
	for i, v := range arr {
		if i == 0 {
			s += strings.ToLower(v)
		} else {
			s += strings.ToUpper(v[:1]) + v[1:]
		}
	}
	return s
}

func (fs FString) Capitalize(value string) string {
	return strings.ToUpper(value[:1]) + strings.ToLower(value[1:])
}

func (fs FString) EndsWith(value, target string) bool {
	if len(target) > len(value) {
		return false
	}
	targetLen := len(target)
	return value[(len(value)-targetLen):] == target
}

func (fs FString) StartWith(value, target string) bool {
	if len(target) > len(value) {
		return false
	}
	targetLen := len(target)
	return value[:targetLen] == target
}

func (fs FString) LowerFirst(value string) string {
	if value == "" {
		return value
	}
	return fs.ToLower(value[:1]) + value[1:]
}

func (fs FString) UpperFirst(value string) string {
	if value == "" {
		return value
	}
	return fs.ToUpper(value[:1]) + value[1:]
}

func (fs FString) PadEnd(value string, length int) string {
	padStr := ""
	for i := 0; i < length; i++ {
		padStr += " "
	}
	return value + fs.Repeat(" ", length)
}

func (fs FString) PadStart(value string, length int) string {
	return fs.Repeat(" ", length) + value
}

func (fs FString) ParseInt(value string) int64 {
	if i, err := strconv.Atoi(value); err != nil {
		return 0
	} else {
		return int64(i)
	}
}

func (fs FString) Repeat(value string, length int) string {
	str := ""
	for i := 0; i < length; i++ {
		str += value
	}
	return str
}

func (fs FString) Replace(value, find, replacement string) string {
	return strings.Replace(value, find, replacement, 1)
}

func (fs FString) ReplaceAll(value, find, replacement string) string {
	return strings.ReplaceAll(value, find, replacement)
}

func (fs FString) SnakeCase(value string) string {
	s := ""
	sep := "_"
	temp := ""
	for i, c := range value {
		if i > 0 && c >= 65 && c <= 90 {
			s += strings.ToLower(temp) + sep
			temp = ""
		}
		temp += string(c)
	}
	s += strings.ToLower(temp)
	return s
}

func (fs FString) Split(value, sep string) []string {
	return strings.Split(value, sep)
}

func (fs FString) ToLower(value string) string {
	return strings.ToLower(value)
}

func (fs FString) ToUpper(value string) string {
	return strings.ToUpper(value)
}

func (fs FString) Trim(value, cutset string) string {
	return strings.Trim(value, cutset)
}

func (fs FString) TrimEnd(value, cutset string) string {
	return strings.TrimRight(value, cutset)
}

func (fs FString) TrimStart(value, cutset string) string {
	return strings.TrimLeft(value, cutset)
}

func (fs FString) Truncate(value string, length int, omission string) string {
	defaultLen := 30
	defaultOmission := "..."
	if length == 0 {
		length = defaultLen
	}
	if omission == "" {
		omission = defaultOmission
	}
	return value[:length] + omission
}

func (fs FString) RandomNumber(length int) string {
	var str string
	b2 := new(big.Int).SetInt64(int64(10))
	for i := 0; i < length; i++ {
		ri, _ := rand.Int(rand.Reader, b2)
		rii := int(ri.Int64())
		str += strconv.Itoa(rii)
	}
	return str
}

func (fs FString) RandString(len int) string {
	bytes := make([]byte, len)

	b2 := new(big.Int).SetInt64(int64(26))
	for i := 0; i < len; i++ {
		ri, _ := rand.Int(rand.Reader, b2)
		bytes[i] = byte(ri.Int64() + 65)
	}
	return string(bytes)
}
