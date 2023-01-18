package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
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

type Edges = [4]uint16

type Tiles = map[int]*Tile

type Tile struct {
	Id            int
	Data          utils.Matrix[bool]
	EdgesVariants [8]Edges
}

//func (t Tile) GetEdges(rotation int) [4]uint16 {
//	return t.EdgesVariants[rotation%2]
//}

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

		for orientation, edgesC := range candidate.EdgesVariants {
			for edgeIndexC, edgeC := range edgesC {
				for edgeIndex, edge := range tile.EdgesVariants[0] {
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

func searchRight(tile *Tile, rightTiles collections.Stack[*Tile], width, expectedSize int, mainTile *Tile, availableTiles Tiles) {
	delete(availableTiles, tile.Id)
	rightTiles.Push(tile)

	rightEdge := tile.EdgesVariants[0][utils.Right]

	// find neighbours on right side
	for _, candidate := range availableTiles {
		for _, edges := range candidate.EdgesVariants {
			if rightEdge == edges[utils.Left] {
				// try recursive
				searchRight(candidate, rightTiles, width+1, expectedSize, mainTile, availableTiles)
			}
		}
	}

	fmt.Printf("  Searching left @ %v (%v)\n", width, tile.Id)
	searchLeft(mainTile, collections.Stack[*Tile]{}, rightTiles, width, expectedSize, availableTiles)

	// remove main again
	delete(availableTiles, mainTile.Id)

	// return the tile back to searchable tiles
	availableTiles[tile.Id] = tile
	rightTiles.Pop()
}

func searchLeft(tile *Tile, leftTiles, rightTiles collections.Stack[*Tile], width, expectedSize int, availableTiles Tiles) {
	delete(availableTiles, tile.Id)
	leftTiles.Push(tile)

	leftEdge := tile.EdgesVariants[0][utils.Left]

	// find neighbours on right side
	for _, candidate := range availableTiles {
		for _, edges := range candidate.EdgesVariants {
			if leftEdge == edges[utils.Right] {
				// try recursive
				searchLeft(candidate, leftTiles, rightTiles, width+1, expectedSize, availableTiles)
			}
		}
	}

	if width >= expectedSize {
		// reverse left side
		row := slices.Reverse(leftTiles.PeekAll())
		// remove main tile
		row = row[:len(row)-1]
		// add right side
		row = append(row, rightTiles.PeekAll()...)

		ids := slices.Map(row, func(t *Tile) int { return t.Id })
		fmt.Printf("  * found row of %v tiles - %v\n", width, ids)

		//searchRowAbove(collections.Stack[*Tile]{}, 0, row, availableTiles)
	}
	//fmt.Printf("Not found\n")

	// return the tile back to searchable tiles
	availableTiles[tile.Id] = tile
	leftTiles.Pop()
}

func DoWithInputPart01(world World) int {
	tiles := world.Tiles

	//id, tile := maps.FirstEntry(tiles)
	//delete(tiles, id)

	//for id := range tiles {
	//	connections := findCandidates(id, tiles)
	//	fmt.Printf("%4v: %3v candidates\n", id, len(connections))
	//	for _, connection := range connections {
	//		fmt.Printf("  * %4v @ %v @ %v -> %4v @ %v @ %v\n", connection.TileA.Id, connection.EdgeA, 0, connection.TileB.Id, connection.EdgeB, connection.Flipper)
	//	}
	//}

	for _, tile := range tiles {
		fmt.Printf("#%v\n", tile.Id)
		searchRight(tile, collections.Stack[*Tile]{}, 1, 3, tile, tiles)
	}

	//tile := tiles[2311]
	//fmt.Printf("#%v\n", tile.Id)
	//searchRight(tile, collections.Stack[*Tile]{}, 1, 3, tile, tiles)

	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func extractEdges(data utils.Matrix[bool]) (Edges, Edges) {
	var edges, flippedEdges Edges

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

	edges[0] = utils.ParseBinaryBool16(top)
	edges[1] = utils.ParseBinaryBool16(right)
	edges[2] = utils.ParseBinaryBool16(bottom)
	edges[3] = utils.ParseBinaryBool16(left)

	flippedEdges[0] = utils.ParseBinaryBool16(slices.Reverse(top))
	flippedEdges[1] = utils.ParseBinaryBool16(slices.Reverse(left))
	flippedEdges[2] = utils.ParseBinaryBool16(slices.Reverse(bottom))
	flippedEdges[3] = utils.ParseBinaryBool16(slices.Reverse(right))

	return edges, flippedEdges
}

// rotate rotates tile edges counter-clockwise by simply shifting the indexes
func rotate(edges Edges, amount int) Edges {
	rotated := edges

	for i, edge := range edges {
		rotated[(i+amount)%4] = edge
	}

	return rotated
}

func rotateAndFlipEdges(edges, flippedEdges Edges) [8]Edges {
	variants := [8]Edges{}
	for i := 0; i < 4; i++ {
		variants[i] = rotate(edges, i)
	}

	for i := 0; i < 4; i++ {
		variants[i+4] = rotate(flippedEdges, i)
	}

	return variants
}

func ParseInput(r io.Reader) World {
	tiles := make(Tiles)

	parseTile := func(lines []string, i int) Tile {
		id := utils.ExtractInts(lines[0], false)[0]

		dataString := strings.Join(lines[1:], "\n")
		reader := strings.NewReader(dataString)

		data := parsers.ParseToMatrix(reader, parsers.MapperBoolean('#', '.'))
		edges := rotateAndFlipEdges(extractEdges(data))

		tile := Tile{
			Id:            id,
			Data:          data,
			EdgesVariants: edges,
		}

		tiles[id] = &tile

		return tile
	}

	parsers.ParseToGroups(r, parseTile)
	return World{Tiles: tiles}
}