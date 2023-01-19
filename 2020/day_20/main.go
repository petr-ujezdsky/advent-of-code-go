package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"math"
	"strings"
)

type Connection struct {
	TileA, TileB *Tile
	EdgeA, EdgeB int
	Flipper      int
}

type Edge struct {
	Hash     uint16
	Booleans []bool
	Reversed *Edge
}

func NewEdge(booleans []bool) Edge {
	reversed := slices.Reverse(booleans)
	return Edge{
		Hash:     utils.ParseBinaryBool16(booleans),
		Booleans: booleans,
		Reversed: &Edge{
			Hash:     utils.ParseBinaryBool16(reversed),
			Booleans: reversed,
			Reversed: nil,
		},
	}
}

type Edges = [4]Edge

type Tiles = map[int]*Tile

type OrientedTile struct {
	Id    int
	Edges Edges
	Tile  *Tile
}

type Tile struct {
	Id            int
	Data          utils.Matrix[bool]
	OrientedTiles [8]OrientedTile
}

//func (t Tile) GetEdges(rotation int) [4]uint16 {
//	return t.EdgesVariants[rotation%2]
//}

type World struct {
	Tiles Tiles
}

//func findCandidates(id int, tiles Tiles) []Connection {
//	tile, ok := tiles[id]
//	if !ok {
//		panic("Not found")
//	}
//
//	var connections []Connection
//	for _, candidate := range tiles {
//		// skip same tile
//		if candidate.Id == id {
//			continue
//		}
//
//		for orientation, edgesC := range candidate.OrientedTiles {
//			for edgeIndexC, edgeC := range edgesC {
//				for edgeIndex, edge := range tile.EdgesVariants[0] {
//					if edgeC == edge {
//						connections = append(connections, Connection{
//							TileA:   tile,
//							TileB:   candidate,
//							EdgeA:   edgeIndex,
//							EdgeB:   edgeIndexC,
//							Flipper: orientation,
//						})
//					}
//				}
//			}
//		}
//	}
//
//	return connections
//}

func searchRight(tile *OrientedTile, rightTiles collections.Stack[*OrientedTile], width, expectedSize int, mainTile *OrientedTile, availableTiles Tiles) {
	delete(availableTiles, tile.Id)
	rightTiles.Push(tile)

	rightEdge := tile.Edges[utils.Right]

	// find neighbours on right side
	for _, candidate := range availableTiles {
		for _, orientedTile := range candidate.OrientedTiles {
			if rightEdge.Hash == orientedTile.Edges[utils.Left].Reversed.Hash {
				// try recursive
				searchRight(&orientedTile, rightTiles, width+1, expectedSize, mainTile, availableTiles)
			}
		}
	}

	fmt.Printf("  Searching left @ %v (%v)\n", width, tile.Id)
	searchLeft(mainTile, collections.Stack[*OrientedTile]{}, rightTiles, width, expectedSize, availableTiles)

	// remove main again
	delete(availableTiles, mainTile.Id)

	// return the tile back to searchable tiles
	availableTiles[tile.Id] = tile.Tile
	rightTiles.Pop()
}

func searchLeft(tile *OrientedTile, leftTiles, rightTiles collections.Stack[*OrientedTile], width, expectedSize int, availableTiles Tiles) {
	delete(availableTiles, tile.Id)
	leftTiles.Push(tile)

	leftEdge := tile.Edges[utils.Left]

	// find neighbours on right side
	for _, candidate := range availableTiles {
		for _, orientedTile := range candidate.OrientedTiles {
			if leftEdge.Hash == orientedTile.Edges[utils.Right].Reversed.Hash {
				// try recursive
				searchLeft(&orientedTile, leftTiles, rightTiles, width+1, expectedSize, availableTiles)
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

		ids := slices.Map(row, func(t *OrientedTile) int { return t.Id })
		fmt.Printf("  * found row of %v tiles - %v\n", width, ids)

		searchRowAbove(nil, collections.Stack[*OrientedTile]{}, collections.Stack[[]*OrientedTile]{}, 0, row, row, availableTiles)
	}
	//fmt.Printf("Not found\n")

	// return the tile back to available tiles
	availableTiles[tile.Id] = tile.Tile
	leftTiles.Pop()
}

func searchRowAbove(tile *OrientedTile, rightTiles collections.Stack[*OrientedTile], aboveRows collections.Stack[[]*OrientedTile], i int, rowBelow, mainRow []*OrientedTile, availableTiles Tiles) {
	if tile != nil {
		delete(availableTiles, tile.Id)
		rightTiles.Push(tile)
	}

	if i < len(rowBelow) {
		below := rowBelow[i]
		// find neighbours on right side matching neighbours below
		for _, candidate := range availableTiles {
			for _, orientedTile := range candidate.OrientedTiles {
				if matches(orientedTile, nil, nil, below, tile) {
					searchRowAbove(&orientedTile, rightTiles, aboveRows, i+1, rowBelow, mainRow, availableTiles)
				}
			}
		}
	} else {
		row := rightTiles.PeekAll()
		aboveRows.Push(row)

		ids := slices.Map(row, func(t *OrientedTile) int { return t.Id })
		fmt.Printf("    * found row above %v\n", ids)
		// find another row above
		searchRowAbove(nil, collections.Stack[*OrientedTile]{}, aboveRows, 0, row, mainRow, availableTiles)

		// find rows below main row
		searchRowBelow(nil, collections.Stack[*OrientedTile]{}, collections.Stack[[]*OrientedTile]{}, aboveRows, 0, mainRow, availableTiles)
		aboveRows.Pop()
	}

	// return the tile back to searchable tiles
	if tile != nil {
		availableTiles[tile.Id] = tile.Tile
		rightTiles.Pop()
	}
}

func searchRowBelow(tile *OrientedTile, rightTiles collections.Stack[*OrientedTile], belowRows, aboveRows collections.Stack[[]*OrientedTile], i int, rowAbove []*OrientedTile, availableTiles Tiles) {
	if tile != nil {
		delete(availableTiles, tile.Id)
		rightTiles.Push(tile)
	}

	if len(availableTiles) == 0 {
		// found it!
		// reverse above rows
		rows := slices.Reverse(aboveRows.PeekAll())
		// remove main row
		rows = rows[:len(rows)-1]
		// add below rows
		rows = append(rows, belowRows.PeekAll()...)

		idRows := slices.Map(rows, func(r []*OrientedTile) []int {
			return slices.Map(r, func(t *OrientedTile) int {
				return t.Id
			})
		})
		fmt.Printf("    * found it:\n")
		for _, ids := range idRows {
			fmt.Printf("    *           %v\n", ids)
		}

		//idRows :=
		//idRows := make([][]int, len(rows))
		//for j, row := range rows {
		//	ids := slices.Map(row, func(t *OrientedTile) int { return t.Id })
		//	idRows[j] = ids
		//	fmt.Printf("    * found row below %v\n", ids)
		//}
	}

	if i < len(rowAbove) {
		above := rowAbove[i]
		// find neighbours on right side matching neighbours above
		for _, candidate := range availableTiles {
			for _, orientedTile := range candidate.OrientedTiles {
				if matches(orientedTile, above, nil, nil, tile) {
					searchRowBelow(&orientedTile, rightTiles, belowRows, aboveRows, i+1, rowAbove, availableTiles)
				}
			}
		}
	} else {
		row := rightTiles.PeekAll()
		belowRows.Push(row)
		// find another row above
		ids := slices.Map(row, func(t *OrientedTile) int { return t.Id })
		fmt.Printf("    * found row below %v\n", ids)

		// find another row below
		searchRowBelow(nil, collections.Stack[*OrientedTile]{}, belowRows, aboveRows, 0, row, availableTiles)
		belowRows.Pop()
	}

	// return the tile back to searchable tiles
	if tile != nil {
		availableTiles[tile.Id] = tile.Tile
		rightTiles.Pop()
	}
}

func matches(tile OrientedTile, above, right, below, left *OrientedTile) bool {
	if above != nil && above.Edges[utils.Down].Reversed.Hash != tile.Edges[utils.Up].Hash {
		return false
	}

	if right != nil && right.Edges[utils.Left].Reversed.Hash != tile.Edges[utils.Right].Hash {
		return false
	}

	if below != nil && below.Edges[utils.Up].Reversed.Hash != tile.Edges[utils.Down].Hash {
		return false
	}

	if left != nil && left.Edges[utils.Right].Reversed.Hash != tile.Edges[utils.Left].Hash {
		return false
	}

	return true
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

	expectedSize := int(math.Sqrt(float64(len(world.Tiles))))

	//for _, tile := range tiles {
	//	orientedTile := &tile.OrientedTiles[0]
	//	fmt.Printf("#%v\n", tile.Id)
	//	searchRight(orientedTile, collections.Stack[*OrientedTile]{}, 1, expectedSize, orientedTile, tiles)
	//}

	//tile := tiles[2311]
	//tile := tiles[1427]

	//tile := tiles[3079] // top right
	//orientedTile := &tile.OrientedTiles[0]

	tile := tiles[1427] // middle middle
	orientedTile := &tile.OrientedTiles[6]
	fmt.Printf("#%v\n", tile.Id)
	searchRight(orientedTile, collections.Stack[*OrientedTile]{}, 1, expectedSize, orientedTile, tiles)

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

	edges[0] = NewEdge(top)
	edges[1] = NewEdge(right)
	edges[2] = NewEdge(bottom)
	edges[3] = NewEdge(left)

	flippedEdges[0] = NewEdge(slices.Reverse(top))
	flippedEdges[1] = NewEdge(slices.Reverse(left))
	flippedEdges[2] = NewEdge(slices.Reverse(bottom))
	flippedEdges[3] = NewEdge(slices.Reverse(right))

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

	parseTile := func(lines []string, i int) *Tile {
		id := utils.ExtractInts(lines[0], false)[0]

		dataString := strings.Join(lines[1:], "\n")
		reader := strings.NewReader(dataString)

		data := parsers.ParseToMatrix(reader, parsers.MapperBoolean('#', '.'))
		edges := rotateAndFlipEdges(extractEdges(data))

		tile := &Tile{
			Id:            id,
			Data:          data,
			OrientedTiles: [8]OrientedTile{},
		}

		for j, edge := range edges {
			tile.OrientedTiles[j] = OrientedTile{
				Id:    id,
				Edges: edge,
				Tile:  tile,
			}
		}

		tiles[id] = tile

		return tile
	}

	parsers.ParseToGroups(r, parseTile)
	return World{Tiles: tiles}
}
