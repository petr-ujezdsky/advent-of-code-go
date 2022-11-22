package day_12

import (
	"bufio"
	"io"
	"strings"
)

type World struct {
	template string
	rules    map[string]string
}

func ParseInput(r io.Reader) (World, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	template := scanner.Text()

	rules := make(map[string]string)
	scanner.Scan()

	for scanner.Scan() {
		ruleParts := strings.Split(scanner.Text(), " -> ")
		rules[ruleParts[0]] = ruleParts[1]
	}

	return World{template, rules}, scanner.Err()
}
