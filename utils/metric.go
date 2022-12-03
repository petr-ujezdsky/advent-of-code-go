package utils

import "fmt"

type Metric struct {
	Iterations, Sum int
	Enabled         bool
}

func (m *Metric) TickSum(period, v int) {
	m.Iterations++
	m.Sum += v

	if m.Enabled && m.Iterations%period == 0 {
		fmt.Printf("Iteration #%d, sum = %d\n", m.Iterations, m.Sum)
	}
}

func (m *Metric) Tick(period int) {
	m.Iterations++

	if m.Enabled && m.Iterations%period == 0 {
		fmt.Printf("Iteration #%d\n", m.Iterations)
	}
}
