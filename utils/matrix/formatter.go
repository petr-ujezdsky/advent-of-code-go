package matrix

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
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

func StringFmtSeparator[T any](view View[T], separator string, formatter ValueFormatter[T]) string {
	adapter := func(value T, x, y int) string {
		return formatter(value)
	}

	return StringFmtSeparatorIndexed(view, false, separator, adapter)
}

func StringFmtSeparatorIndexed[T any](view View[T], rulers bool, separator string, formatter ValueFormatterIndexed[T]) string {
	return StringFmtSeparatorIndexedOrigin(view, rulers, utils.Vector2i{}, separator, formatter)
}
func StringFmtSeparatorIndexedOrigin[T any](view View[T], rulers bool, origin utils.Vector2i, separator string, formatter ValueFormatterIndexed[T]) string {
	var sb strings.Builder

	if rulers {
		for i := 0; i < 3; i++ {
			sb.WriteString("    ")

			for x := 0; x < view.GetWidth(); x++ {
				sb.WriteRune(rune(fmt.Sprintf("%3d ", x+origin.X)[i]))
			}
			sb.WriteString("\n")
		}

		sb.WriteString("\n")
	}

	for y := 0; y < view.GetHeight(); y++ {
		for x := 0; x < view.GetWidth(); x++ {
			val := view.Get(x, y)

			if rulers && x == 0 {
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
