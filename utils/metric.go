package utils

import "fmt"

type Metric struct {
	Name            string
	Ticks, Sum, Max int
	Enabled         bool
}

func NewMetric(name string) *Metric {
	return &Metric{
		Name: name,
	}
}

func (m *Metric) TickSum(period, v int) {
	m.Ticks++
	m.Sum += v

	if m.Enabled && m.Ticks%period == 0 {
		fmt.Printf("%v - tick #%d, sum = %d\n", m.Name, m.Ticks, m.Sum)
	}
}

func (m *Metric) TickCurrent(period, v int) {
	m.Ticks++

	if m.Enabled && m.Ticks%period == 0 {
		fmt.Printf("%v - tick #%d, current = %d\n", m.Name, m.Ticks, v)
	}
}

func (m *Metric) TickMax(period, v int) {
	m.Ticks++
	m.Max = Max(m.Max, v)

	if m.Enabled && m.Ticks%period == 0 {
		fmt.Printf("%v - tick #%d, max = %d\n", m.Name, m.Ticks, v)
	}
}

func (m *Metric) Tick(period int) {
	m.Ticks++

	if m.Enabled && m.Ticks%period == 0 {
		fmt.Printf("%v - tick #%d\n", m.Name, m.Ticks)
	}
}

func (m *Metric) Finished() {
	if m.Enabled {
		fmt.Printf("%v - ticks = %d, sum = %d (finished)\n", m.Name, m.Ticks, m.Sum)
	}
}

type Metrics []*Metric

func (m Metrics) Enable() {
	for _, metric := range m {
		metric.Enabled = true
	}
}
