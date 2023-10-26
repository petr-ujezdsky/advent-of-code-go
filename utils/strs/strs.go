package strs

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"strconv"
	"unsafe"
)

func ReverseString(str string) string {
	return string(slices.Reverse([]rune(str)))
}

func Substring(str string, from, to int) string {
	return str[from:to]
}

func ToBinary[T utils.AnyNumber](value T) string {
	bitsCount := int(unsafe.Sizeof(value)) * 8
	format := "%." + strconv.Itoa(bitsCount) + "b"
	return fmt.Sprintf(format, value)
}
