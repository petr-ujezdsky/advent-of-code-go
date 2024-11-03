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

func initWorld(program []int) map[utils.Vector2i]TileType {
	input := make(chan int, 1)
	output := make(chan int)
	halt := make(chan int)

	defer close(halt)

	computer := common.NewIntCodeComputer("Unknown", program, input, output, halt)

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

	return tiles
}

func DoWithInputPart01(world World) int {
	tiles := initWorld(world.Program)
	tilesMatrix := matrix.NewMatrixFromMap[TileType](tiles)

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
	//tiles := matrix.NewMatrixFromMap[TileType](initWorld(slices.Clone(world.Program)))
	tiles := matrix.NewMatrix[TileType](37, 24)

	input := make(chan int)
	readyInput := make(chan int)
	output := make(chan int)
	halt := make(chan int)

	defer close(halt)

	// insert coins
	world.Program[0] = 2
	computer := common.NewIntCodeComputerReadyInput("Unknown", world.Program, input, readyInput, output, halt)

	go common.Run(computer)

	outputBuffered := make(chan [3]int)

	// read 3 values from output and send them to formattedOutput channel
	go func() {
		for a := range output {
			b := <-output
			c := <-output

			outputBuffered <- [3]int{a, b, c}
		}

		close(outputBuffered)
	}()

	score := 0
	end := make(chan int)
	positionBallCh := make(chan utils.Vector2i, 2)
	//positionPaddleCh := make(chan utils.Vector2i, 2)

	//positionBall := utils.Vector2i{}
	//positionPaddle := utils.Vector2i{}

	go func() {
		for buffer := range outputBuffered {
			position := utils.Vector2i{X: buffer[0], Y: buffer[1]}

			if position == (utils.Vector2i{X: -1, Y: 0}) {
				score = buffer[2]
				continue
			}

			tileType := TileType(buffer[2])
			tiles.SetV(position, tileType)

			switch tileType {
			case Ball:
				positionBallCh <- position
				//positionBall = position
			case HorizontalPaddle:
				//positionPaddleCh <- position
				//positionPaddle = position
			}
		}

		end <- 0
	}()

	go func() {
		positionPaddle := utils.Vector2i{X: 18, Y: 22}

		for {
			readyInput <- 1

			positionBall := <-positionBallCh
			//positionPaddle := <-positionPaddleCh

			printTiles(tiles)
			println()
			println()
			println()

			dir := utils.Signum(positionBall.X - positionPaddle.X)
			positionPaddle.X += dir

			input <- dir
		}
	}()

	<-halt
	close(output)
	close(input)

	<-end

	printTiles(tiles)

	return score
}

func printTiles(m matrix.Matrix[TileType]) {
	fmt.Println(matrix.StringFmt(m, func(value TileType) string {
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
			return "?"
		}
	}))
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	return World{Program: utils.ExtractInts(scanner.Text(), true)}
}
