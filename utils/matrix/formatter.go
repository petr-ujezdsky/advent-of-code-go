package matrix

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"strconv"
	"strings"
)

type ViewFormatter[T any] struct {
	View View[T]
}

type ValueFormatter[T any] func(value T) string

type ValueFormatterIndexed[T any] func(value T, x, y int) string

func StringFmt[T any](view View[T], formatter ValueFormatter[T]) string {
	return StringFmtSeparator(view, " ", formatter)
}

func NonIndexedAdapter[T any](formatter ValueFormatter[T]) ValueFormatterIndexed[T] {
	return func(value T, x, y int) string {
		return formatter(value)
	}
}

func StringFmtSeparator[T any](view View[T], separator string, formatter ValueFormatter[T]) string {
	return StringFmtSeparatorIndexed(view, false, separator, NonIndexedAdapter(formatter))
}

func StringFmtSeparatorIndexed[T any](view View[T], rulers bool, separator string, formatter ValueFormatterIndexed[T]) string {
	rulerWidth := 0
	if rulers {
		rulerWidth = 1
	}
	return StringFmtSeparatorIndexedOrigin(view, rulerWidth, utils.Vector2i{}, separator, formatter)
}
func StringFmtSeparatorIndexedOrigin[T any](view View[T], rulerWidth int, origin utils.Vector2i, separator string, formatter ValueFormatterIndexed[T]) string {
	var sb strings.Builder

	if rulerWidth > 0 {
		for i := 0; i < 3; i++ {
			sb.WriteString("    ")

			for x := 0; x < view.GetWidth(); x++ {
				if x > 0 {
					sb.WriteString(separator)
				}

				rulerDigit := string(fmt.Sprintf("%3d ", x+origin.X)[i])
				sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(rulerWidth)+"v", rulerDigit))
			}
			sb.WriteString("\n")
		}

		sb.WriteString("\n")
	}

	for y := 0; y < view.GetHeight(); y++ {
		for x := 0; x < view.GetWidth(); x++ {
			val := view.Get(x, y)

			if rulerWidth > 0 && x == 0 {
				sb.WriteString(fmt.Sprintf("%3d ", y+origin.Y))
			}

			if x > 0 {
				sb.WriteString(separator)
			}
			sb.WriteString(formatter(val, x, y))
		}
		if y < view.GetHeight()-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func FmtNative[T any](value T) string {
	return fmt.Sprint(value)
}

func FmtFmt[T any](format string) func(v T) string {
	return func(val T) string {
		return fmt.Sprintf(format, val)
	}
}

func FmtConstant[T any](value string) func(v T) string {
	return func(val T) string {
		return value
	}
}

func FmtMap[T comparable](mapper map[T]string) func(v T) string {
	return func(val T) string {
		return mapper[val]
	}
}

func FmtBoolean[T comparable](val T) string {
	return FmtBooleanConst[T](".", "#")(val)
}

func FmtBooleanConst[T comparable](falseVal, trueVal string) ValueFormatter[T] {
	return FmtBooleanCustom[T](FmtConstant[T](falseVal), FmtConstant[T](trueVal))
}

func FmtBooleanCustom[T comparable](formatterFalse, formatterTrue ValueFormatter[T]) func(v T) string {
	return func(val T) string {
		var empty T
		if val == empty {
			return formatterFalse(empty)
		} else {
			return formatterTrue(val)
		}
	}
}
