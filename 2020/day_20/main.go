package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"strings"
)

type Connection struct {
	TileA, TileB *Tile
	EdgeA, EdgeB int
	Flipper      int
}

type Edges = [2][4]uint16

type Tiles = map[int]*Tile

type Tile struct {
	Id    int
	Data  utils.Matrix[bool]
	Edges Edges
}

func (t Tile) GetEdges(rotation int) [4]uint16 {
	return t.Edges[rotation%2]
}

type World struct {
	Tiles Tiles
}

func findCandidates(id int, tiles Tiles) []Connection {
	tile, ok := tiles[id]
	if !ok {
		panic("Not found")
	}

	var connections []Connection
	for _, candidate := range tiles {
		// skip same tile
		if candidate.Id == id {
			continue
		}

		for orientation, edgesC := range candidate.Edges {
			for edgeIndexC, edgeC := range edgesC {
				for edgeIndex, edge := range tile.Edges[0] {
					if edgeC == edge {
						connections = append(connections, Connection{
							TileA:   tile,
							TileB:   candidate,
							EdgeA:   edgeIndex,
							EdgeB:   edgeIndexC,
							Flipper: orientation,
						})
					}
				}
			}
		}
	}

	return connections
}

func DoWithInputPart01(world World) int {
	tiles := world.Tiles

	//id, tile := maps.FirstEntry(tiles)
	//delete(tiles, id)

	for id := range tiles {
		connections := findCandidates(id, tiles)
		fmt.Printf("%4v: %3v candidates\n", id, len(connections))
		for _, connection := range connections {
			fmt.Printf("  * %4v @ %v @ %v -> %4v @ %v @ %v\n", connection.TileA.Id, connection.EdgeA, 0, connection.TileB.Id, connection.EdgeB, connection.Flipper)
		}
	}

	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func extractEdges(data utils.Matrix[bool]) Edges {
	var edges Edges

	top := make([]bool, data.Width)
	bottom := make([]bool, data.Width)
	for x := 0; x < data.Width; x++ {
		top[x] = data.Columns[x][0]
		bottom[data.Width-x-1] = data.Columns[x][data.Height-1]
	}

	right := make([]bool, data.Height)
	left := make([]bool, data.Height)
	for y := 0; y < data.Height; y++ {
		right[y] = data.Columns[data.Width-1][y]
		left[data.Height-y-1] = data.Columns[0][y]
	}

	edges[0][0] = utils.ParseBinaryBool16(top)
	edges[0][1] = utils.ParseBinaryBool16(right)
	edges[0][2] = utils.ParseBinaryBool16(bottom)
	edges[0][3] = utils.ParseBinaryBool16(left)

	edges[1][0] = utils.ParseBinaryBool16(slices.Reverse(top))
	edges[1][1] = utils.ParseBinaryBool16(slices.Reverse(left))
	edges[1][2] = utils.ParseBinaryBool16(slices.Reverse(bottom))
	edges[1][3] = utils.ParseBinaryBool16(slices.Reverse(right))

	return edges
}

func ParseInput(r io.Reader) World {
	tiles := make(Tiles)

	parseTile := func(lines []string, i int) Tile {
		id := utils.ExtractInts(lines[0], false)[0]

		dataString := strings.Join(lines[1:], "\n")
		reader := strings.NewReader(dataString)

		data := parsers.ParseToMatrix(reader, parsers.MapperBoolean('#', '.'))
		edges := extractEdges(data)

		tile := Tile{
			Id:    id,
			Data:  data,
			Edges: edges,
		}

		tiles[id] = &tile

		return tile
	}

	parsers.ParseToGroups(r, parseTile)
	return World{Tiles: tiles}
}
