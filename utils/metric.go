package utils

import (
	"fmt"
	"time"
)

type Metric struct {
	Name              string
	Ticks, Sum, Max   int
	PreviousTimestamp int64
	Enabled           bool
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

func (m *Metric) TickTotal(period, total int) {
	m.Ticks++

	if m.Enabled && m.Ticks%period == 0 {
		percent := float64(100*m.Ticks) / float64(total)
		fmt.Printf("%v - tick #%d, %d / %d (%.4f%%)\n", m.Name, m.Ticks, m.Ticks, total, percent)
	}
}

func (m *Metric) TickMax(period, v int) {
	m.Ticks++
	m.Max = Max(m.Max, v)

	if m.Enabled && m.Ticks%period == 0 {
		fmt.Printf("%v - tick #%d, max = %d\n", m.Name, m.Ticks, m.Max)
	}
}

func (m *Metric) TickTime(period int) {
	// initialize previous timestamp in first iteration
	if m.Enabled && m.Ticks == 0 {
		m.PreviousTimestamp = time.Now().UnixMilli()
	}

	m.Ticks++

	if m.Enabled && m.Ticks%period == 0 {
		currentTimestamp := time.Now().UnixMilli()
		elapsedMs := int(currentTimestamp - m.PreviousTimestamp)

		fmt.Printf("%v - tick #%d, elapsed time = %dms (%d ticks/s)\n", m.Name, m.Ticks, elapsedMs, 1000*period/elapsedMs)

		m.PreviousTimestamp = currentTimestamp
	}
}

func (m *Metric) Tick(period int) {
	m.Ticks++

	if m.Enabled && m.Ticks%period == 0 {
		fmt.Printf("%v - tick #%d\n", m.Name, m.Ticks)
	}
}

func (m *Metric) Enable() *Metric {
	m.Enabled = true
	return m
}

func (m *Metric) Disable() *Metric {
	m.Enabled = false
	return m
}

func (m *Metric) Finished() {
	if m.Enabled {
		fmt.Printf("%v - ticks = %d, sum = %d (finished)\n", m.Name, m.Ticks, m.Sum)
	}
}

type Metrics []*Metric

func (m Metrics) Enable() {
	for _, metric := range m {
		metric.Enable()
	}
}
