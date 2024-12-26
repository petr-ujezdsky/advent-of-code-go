package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"strconv"
	"strings"
)

type World struct {
	DiskMap []int
}

type Block struct {
	FileId         int
	Length         int
	Previous, Next *Block
}

func (block *Block) Split(leftLength int) (*Block, *Block) {
	b1 := block
	b2 := &Block{
		FileId:   block.FileId,
		Length:   b1.Length - leftLength,
		Previous: b1,
		Next:     block.Next,
	}

	b1.Next = b2
	b1.Length = leftLength

	return b1, b2
}

type Blocks struct {
	Head, Tail *Block
	FilesIndex []*Block
}

func (blocks Blocks) String() string {
	var sb strings.Builder

	for block := blocks.Head; block != nil; block = block.Next {
		for i := 0; i < block.Length; i++ {
			if block.FileId == -1 {
				sb.WriteRune('.')
			} else {
				sb.WriteString(strconv.Itoa(block.FileId))
			}
		}

	}

	return sb.String()
}

func (blocks Blocks) Checksum() int {
	sum := 0
	position := 0

	for block := blocks.Head; block != nil; block = block.Next {
		if block.FileId != -1 {
			sum += block.FileId * utils.SumNtoM(position, position+block.Length-1)
		}

		position += block.Length
	}

	return sum
}

func parseBlocks(diskMap []int) Blocks {
	fileIdSeq := 0
	var filesIndex []*Block

	var head *Block
	var previous *Block
	for i, size := range diskMap {
		isFile := i%2 == 0

		block := &Block{
			FileId:   -1,
			Length:   size,
			Previous: previous,
			Next:     nil,
		}

		if isFile {
			block.FileId = fileIdSeq
			fileIdSeq++
			filesIndex = append(filesIndex, block)
		}

		if head == nil {
			head = block
		}

		if previous != nil {
			previous.Next = block
		}

		previous = block
	}

	return Blocks{
		Head:       head,
		Tail:       previous,
		FilesIndex: filesIndex,
	}
}

func printBlocks(blocks Blocks) {
	//fmt.Println(blocks.String())
}

func DoWithInputPart01(world World) int {
	blocks := parseBlocks(world.DiskMap)
	printBlocks(blocks)

	lastFile := blocks.Tail
	i := 0
	for block := blocks.Head; block != nil; block = block.Next {
		i++
		if i > 100000 {
			panic("Too much")
		}
		if block == lastFile {
			break
		}
		if block.FileId != -1 {
			// skip files
			continue
		}

		// find last file block
		for lastFile.FileId == -1 {
			lastFile = lastFile.Previous
		}

		usableSize := utils.Min(block.Length, lastFile.Length)

		if block.Length > usableSize {
			block.Split(usableSize)
		}

		block.FileId = lastFile.FileId
		lastFile.Length -= usableSize

		if lastFile.Length == 0 {
			// remove it
			lastFile.Previous.Next = lastFile.Next
			if lastFile.Next != nil {
				lastFile.Next.Previous = lastFile.Previous
			}

			lastFile = lastFile.Previous
		}

		printBlocks(blocks)
	}

	return blocks.Checksum()
}

func DoWithInputPart02(world World) int {
	blocks := parseBlocks(world.DiskMap)

	fileIndex := slices.Reverse(blocks.FilesIndex)

	for _, file := range fileIndex {

		//fmt.Printf("#%v: ", file.FileId)
		printBlocks(blocks)

		for block := blocks.Head; block != nil; block = block.Next {
			if block == file {
				break
			}

			if block.FileId != -1 {
				// skip files
				continue
			}

			if block.Length < file.Length {
				// too small
				continue
			}

			usableSize := utils.Min(block.Length, file.Length)

			if block.Length > file.Length {
				block.Split(usableSize)
			}

			block.FileId = file.FileId
			// make it free space
			file.FileId = -1
			break
		}
	}

	printBlocks(blocks)

	return blocks.Checksum()
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var numbers []int
	for scanner.Scan() {
		for _, numberChar := range scanner.Text() {
			numbers = append(numbers, int(numberChar-'0'))
		}
	}

	return World{DiskMap: numbers}
}
