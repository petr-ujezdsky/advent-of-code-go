package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"math"
	"strconv"
	"strings"
)

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
	Id               int
	OrientationIndex int
	Edges            Edges
	Tile             *Tile
}

type Tile struct {
	Id            int
	Data          matrix.Matrix[bool]
	OrientedTiles [8]OrientedTile
}

type World struct {
	Tiles Tiles
}

func searchRight(left *OrientedTile, rightTiles collections.Stack[*OrientedTile], expectedSize int, availableTiles Tiles) (*matrix.Matrix[*OrientedTile], bool) {
	availableTiles = maps.Copy(availableTiles)

	delete(availableTiles, left.Id)
	rightTiles.Push(left)

	// find neighbours on right side
	for _, candidate := range availableTiles {
		for _, orientedTile := range candidate.OrientedTiles {
			neighbours := Neighbours{
				Left: left,
			}

			if matches(orientedTile, neighbours) {
				// try recursive
				if result, ok := searchRight(&orientedTile, rightTiles, expectedSize, availableTiles); ok {
					return result, true
				}
			}
		}
	}

	mainTile := rightTiles.PeekAll()[0]
	if result, ok := searchLeft(mainTile, collections.Stack[*OrientedTile]{}, rightTiles, expectedSize, availableTiles); ok {
		return result, true
	}

	rightTiles.Pop()

	return nil, false
}

func searchLeft(right *OrientedTile, leftTiles, rightTiles collections.Stack[*OrientedTile], expectedSize int, availableTiles Tiles) (*matrix.Matrix[*OrientedTile], bool) {
	availableTiles = maps.Copy(availableTiles)

	delete(availableTiles, right.Id)
	leftTiles.Push(right)

	// find neighbours on right side
	for _, candidate := range availableTiles {
		for _, orientedTile := range candidate.OrientedTiles {
			neighbours := Neighbours{
				Right: right,
			}

			if matches(orientedTile, neighbours) {
				// try recursive
				if result, ok := searchLeft(&orientedTile, leftTiles, rightTiles, expectedSize, availableTiles); ok {
					return result, true
				}
			}
		}
	}

	width := leftTiles.Len() + rightTiles.Len() - 1
	if width >= expectedSize {
		// reverse left side
		row := slices.Reverse(leftTiles.PeekAll())
		// remove main tile
		row = row[:len(row)-1]
		// add right side
		row = append(row, rightTiles.PeekAll()...)

		//ids := slices.Map(row, func(t *OrientedTile) int { return t.Id })
		//fmt.Printf("  * found row of %2v tiles - %v\n", width, ids)

		if result, ok := searchRowAbove(row, collections.Stack[[]*OrientedTile]{}, availableTiles); ok {
			return result, true
		}
	}

	leftTiles.Pop()

	return nil, false
}

func searchRowAbove(row []*OrientedTile, aboveRows collections.Stack[[]*OrientedTile], availableTiles Tiles) (*matrix.Matrix[*OrientedTile], bool) {
	aboveRows.Push(row)
	mainRow := aboveRows.PeekAll()[0]
	// search above
	if result, ok := searchRowAboveRight(nil, collections.Stack[*OrientedTile]{}, aboveRows, 0, mainRow, availableTiles); ok {
		return result, true
	}

	// search below
	if result, ok := searchRowBelow(mainRow, collections.Stack[[]*OrientedTile]{}, aboveRows, availableTiles); ok {
		return result, true
	}

	aboveRows.Pop()

	return nil, false
}

func searchRowAboveRight(left *OrientedTile, rightTiles collections.Stack[*OrientedTile], aboveRows collections.Stack[[]*OrientedTile], i int, mainRow []*OrientedTile, availableTiles Tiles) (*matrix.Matrix[*OrientedTile], bool) {
	availableTiles = maps.Copy(availableTiles)
	if left != nil {
		delete(availableTiles, left.Id)
		rightTiles.Push(left)
	}

	rowBelow := aboveRows.Peek()

	if i < len(rowBelow) {
		below := rowBelow[i]
		// find neighbours on right side matching neighbours below
		for _, candidate := range availableTiles {
			for _, orientedTile := range candidate.OrientedTiles {
				neighbours := Neighbours{
					Below: below,
					Left:  left,
				}

				if matches(orientedTile, neighbours) {
					if result, ok := searchRowAboveRight(&orientedTile, rightTiles, aboveRows, i+1, mainRow, availableTiles); ok {
						return result, true
					}
				}
			}
		}
	}

	if i == len(rowBelow) {
		row := rightTiles.PeekAll()

		//ids := slices.Map(row, func(t *OrientedTile) int { return t.Id })
		//fmt.Printf("    * found row above       %v\n", ids)

		// find another row above
		if result, ok := searchRowAbove(row, aboveRows, availableTiles); ok {
			return result, true
		}
	}

	if left != nil {
		rightTiles.Pop()
	}

	return nil, false
}

func searchRowBelow(row []*OrientedTile, belowRows, aboveRows collections.Stack[[]*OrientedTile], availableTiles Tiles) (*matrix.Matrix[*OrientedTile], bool) {
	belowRows.Push(row)
	if result, ok := searchRowBelowRight(nil, collections.Stack[*OrientedTile]{}, belowRows, aboveRows, 0, availableTiles); ok {
		return result, ok
	}

	if len(availableTiles) == 0 {
		// found it!
		// reverse above rows
		rows := slices.Reverse(aboveRows.PeekAll())
		// remove main row
		rows = rows[:len(rows)-1]
		// add below rows
		rows = append(rows, belowRows.PeekAll()...)

		m := matrix.NewMatrixRowNotation(rows)

		//fmt.Printf("    * found solution:\n")
		//fmt.Println(m.StringFmt(func(tile *OrientedTile) string { return strconv.Itoa(tile.Id) }))

		return &m, true
	}

	belowRows.Pop()
	return nil, false
}

func searchRowBelowRight(left *OrientedTile, rightTiles collections.Stack[*OrientedTile], belowRows, aboveRows collections.Stack[[]*OrientedTile], i int, availableTiles Tiles) (*matrix.Matrix[*OrientedTile], bool) {
	availableTiles = maps.Copy(availableTiles)

	if left != nil {
		delete(availableTiles, left.Id)
		rightTiles.Push(left)
	}

	rowAbove := belowRows.Peek()

	if i < len(rowAbove) {
		above := rowAbove[i]
		// find neighbours on right side matching neighbours above
		for _, candidate := range availableTiles {
			for _, orientedTile := range candidate.OrientedTiles {
				neighbours := Neighbours{
					Above: above,
					Left:  left,
				}

				if matches(orientedTile, neighbours) {
					if result, ok := searchRowBelowRight(&orientedTile, rightTiles, belowRows, aboveRows, i+1, availableTiles); ok {
						return result, true
					}
				}
			}
		}
	}

	if i == len(rowAbove) {
		row := rightTiles.PeekAll()

		//ids := slices.Map(row, func(t *OrientedTile) int { return t.Id })
		//fmt.Printf("    * found row below       %v\n", ids)

		// find another row below
		if result, ok := searchRowBelow(row, belowRows, aboveRows, availableTiles); ok {
			return result, true
		}
	}

	if left != nil {
		rightTiles.Pop()
	}

	return nil, false
}

type Neighbours struct {
	Above, Right, Below, Left *OrientedTile
}

func matches(tile OrientedTile, neighbours Neighbours) bool {
	if above := neighbours.Above; above != nil && above.Edges[utils.Down].Reversed.Hash != tile.Edges[utils.Up].Hash {
		return false
	}

	if right := neighbours.Right; right != nil && right.Edges[utils.Left].Reversed.Hash != tile.Edges[utils.Right].Hash {
		return false
	}

	if below := neighbours.Below; below != nil && below.Edges[utils.Up].Reversed.Hash != tile.Edges[utils.Down].Hash {
		return false
	}

	if left := neighbours.Left; left != nil && left.Edges[utils.Right].Reversed.Hash != tile.Edges[utils.Left].Hash {
		return false
	}

	return true
}

func multiplyCorners(picture *matrix.Matrix[*OrientedTile]) int {
	return picture.Columns[0][0].Id *
		picture.Columns[picture.Width-1][0].Id *
		picture.Columns[picture.Width-1][picture.Height-1].Id *
		picture.Columns[0][picture.Height-1].Id
}

func ConnectTilesUsing(tile *Tile, availableTiles Tiles) *matrix.Matrix[*OrientedTile] {
	expectedSize := int(math.Sqrt(float64(len(availableTiles))))

	orientedTile := &tile.OrientedTiles[0]
	picture, ok := searchRight(orientedTile, collections.Stack[*OrientedTile]{}, expectedSize, availableTiles)
	if !ok {
		panic("No solution")
	}

	return picture
}

func ConnectTiles(availableTiles Tiles) *matrix.Matrix[*OrientedTile] {
	tile := maps.FirstValue(availableTiles)
	return ConnectTilesUsing(tile, availableTiles)
}

func DoWithInputPart01(world World) int {
	picture := ConnectTiles(world.Tiles)
	fmt.Println(picture.StringFmt(func(tile *OrientedTile) string { return strconv.Itoa(tile.Id) }))
	return multiplyCorners(picture)
}

func rotateMatrix(orientationIndex int, m matrix.Matrix[bool]) matrix.Matrix[bool] {
	rotated := m

	flipped := orientationIndex/4 == 1
	if flipped {
		rotated = rotated.FlipHorizontal()
	}

	clockwiseAmount := orientationIndex % 4
	rotated = rotated.Rotate90CounterClockwise(-clockwiseAmount)

	return rotated
}

func rotateTile(tile *OrientedTile) matrix.Matrix[bool] {
	return rotateMatrix(tile.OrientationIndex, tile.Tile.Data)
}

func rotateEachTile(connectedTiles *matrix.Matrix[*OrientedTile]) matrix.Matrix[matrix.Matrix[bool]] {
	rotatedPicture := matrix.NewMatrix[matrix.Matrix[bool]](connectedTiles.Width, connectedTiles.Height)

	for x, col := range connectedTiles.Columns {
		for y, tile := range col {
			rotatedPicture.Columns[x][y] = rotateTile(tile)
		}
	}

	return rotatedPicture
}

func removeBorders(rotatedTiles matrix.Matrix[matrix.Matrix[bool]]) matrix.Matrix[bool] {
	tileWidth := rotatedTiles.Columns[0][0].Width
	tileHeight := rotatedTiles.Columns[0][0].Height

	pictureWidth := (tileWidth - 2) * rotatedTiles.Width
	pictureHeight := (tileHeight - 2) * rotatedTiles.Height

	picture := matrix.NewMatrix[bool](pictureWidth, pictureHeight)

	for x, tilesCol := range rotatedTiles.Columns {
		for y, tile := range tilesCol {
			for u, col := range tile.Columns {
				if u == 0 || u == len(tile.Columns)-1 {
					continue
				}

				for v, pixel := range col {
					if v == 0 || v == len(col)-1 {
						continue
					}

					px := x*(tileWidth-2) + u - 1
					py := y*(tileHeight-2) + v - 1

					picture.Columns[px][py] = pixel
				}
			}
		}
	}

	return picture
}

func TilesToPicture(connectedTiles *matrix.Matrix[*OrientedTile]) matrix.Matrix[bool] {
	return removeBorders(rotateEachTile(connectedTiles))
}

var monster = parsers.ParseToMatrix(strings.NewReader(utils.Msg(`
                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `)), parsers.MapperBoolean('#', ' '))

func patternMatches(x, y int, pattern, orig matrix.Matrix[bool]) bool {
	for xx := 0; xx < pattern.Width; xx++ {
		for yy := 0; yy < pattern.Height; yy++ {
			patternPixel := pattern.Columns[xx][yy]
			picturePixel := orig.Columns[x+xx][y+yy]

			if patternPixel && !picturePixel {
				return false
			}
		}
	}

	return true
}

func removePattern(x, y int, pattern, removed matrix.Matrix[bool]) {
	for xx := 0; xx < pattern.Width; xx++ {
		for yy := 0; yy < pattern.Height; yy++ {
			patternPixel := pattern.Columns[xx][yy]

			if patternPixel {
				removed.Columns[x+xx][y+yy] = false
			}
		}
	}
}

func findAndRemoveMonster(pattern, orig, removed matrix.Matrix[bool]) {
	for x, col := range orig.Columns {
		if x > orig.Width-pattern.Width {
			break
		}
		for y := range col {
			if y > orig.Height-pattern.Height {
				break
			}

			// match pattern
			if patternMatches(x, y, pattern, orig) {
				removePattern(x, y, pattern, removed)
			}
		}
	}
}

func rotateAndFlipPattern(pattern matrix.Matrix[bool]) []matrix.Matrix[bool] {
	variations := make([]matrix.Matrix[bool], 8)

	for i := range variations {
		variations[i] = rotateMatrix(i, pattern)
	}

	return variations
}

func DoWithInputPart02(world World) int {
	tile := world.Tiles[3079]
	connectedTiles := ConnectTilesUsing(tile, world.Tiles)
	picture := TilesToPicture(connectedTiles)

	// najit prisery, nerotovat obrazek, ale ty prisery
	removed := picture.Clone()
	for _, m := range rotateAndFlipPattern(monster) {
		findAndRemoveMonster(m, picture, removed)
	}

	count := 0
	for _, col := range removed.Columns {
		for _, pixel := range col {
			if pixel {
				count++
			}
		}
	}

	return count
}

func extractEdges(data matrix.Matrix[bool]) (Edges, Edges) {
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

// rotate rotates tile edges clockwise by simply shifting the indexes
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
				Id:               id,
				OrientationIndex: j,
				Edges:            edge,
				Tile:             tile,
			}
		}

		tiles[id] = tile

		return tile
	}

	parsers.ParseToGroups(r, parseTile)
	return World{Tiles: tiles}
}
