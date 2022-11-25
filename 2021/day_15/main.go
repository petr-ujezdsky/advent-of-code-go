package day_15

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

func ParseInput(r io.Reader) (utils.Matrix2i, error) {
	return utils.ParseToMatrix(r)
}
