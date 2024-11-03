package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/2019/common"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"io"
)

type World struct {
	Program []int
}

type TileType byte

const (
	Empty            TileType = iota
	Wall                      = iota
	Block                     = iota
	HorizontalPaddle          = iota
	Ball                      = iota
)

type PosAndType struct {
	Position utils.Vector2i
	Type     TileType
}

func DoWithInputPart01(world World) int {
	input := make(chan int, 1)
	output := make(chan int)
	halt := make(chan int)

	defer close(halt)

	computer := common.NewIntCodeComputer("Unknown", world.Program, input, output, halt)

	go common.Run(computer)

	formattedOutput := make(chan PosAndType)

	// read 3 values from output and send them to formattedOutput channel
	go func() {
		for x := range output {
			y := <-output
			ttype := <-output

			formattedOutput <- PosAndType{
				Position: utils.Vector2i{X: x, Y: y},
				Type:     TileType(ttype),
			}
		}

		close(formattedOutput)
	}()

	tilesCh := make(chan map[utils.Vector2i]TileType)

	go func() {
		tiles := make(map[utils.Vector2i]TileType)

		for posAndType := range formattedOutput {

			if posAndType.Position.X < 0 || posAndType.Position.Y < 0 {
				fmt.Println("Unseen")
			}

			if _, ok := tiles[posAndType.Position]; ok {
				fmt.Println("Overwrite")
			}

			switch posAndType.Type {
			case Empty:
				delete(tiles, posAndType.Position)
			default:
				tiles[posAndType.Position] = posAndType.Type
			}
		}

		tilesCh <- tiles
	}()

	<-halt
	close(output)
	close(input)

	tiles := <-tilesCh

	tilesMatrix := matrix.NewMatrixFromMap(tiles)

	fmt.Println(matrix.StringFmt(tilesMatrix, func(value TileType) string {
		switch value {
		case Empty:
			return " "
		case Wall:
			return "+"
		case Block:
			return "#"
		case HorizontalPaddle:
			return "-"
		case Ball:
			return "O"
		default:
			panic("Unknown")
		}
	}))

	blockTilesCount := 0
	for _, tileType := range tiles {
		if tileType == Block {
			blockTilesCount++
		}
	}

	return blockTilesCount
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	return World{Program: utils.ExtractInts(scanner.Text(), true)}
}
