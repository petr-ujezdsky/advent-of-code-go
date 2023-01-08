package day_22

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"math"
	"strings"
)

/*
###########1#
#01234567890#
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########
*/

type Room struct {
	collections.Stack[rune]
	Capacity int
}

func (r Room) ContainsOnly(ch rune) bool {
	for _, rch := range r.PeekAll() {
		if rch != ch {
			return false
		}
	}

	return r.Len() > 0
}

func (r Room) CanAccept(ch rune) bool {
	return r.Len() == 0 || r.Len() < r.Capacity && r.ContainsOnly(ch)
}

func (r Room) Clone() Room {
	return Room{
		Stack:    r.Stack.Clone(),
		Capacity: r.Capacity,
	}
}

type Building struct {
	// 00-10 = hallway
	// 11-12 = room1
	// 13-14 = room2
	// 15-16 = room3
	// 17-18 = room4
	Hallway        []rune
	Rooms          []Room
	ConsumedEnergy int
	Previous       *Building
}

func NewBuilding(room1, room2, room3, room4 string) Building {
	return NewBuildingFull("...........", room1, room2, room3, room4, 0)
}

func NewBuildingFull(hallway, room1, room2, room3, room4 string, consumedEnergy int) Building {
	b := Building{
		Hallway:        dot2zero([]rune(hallway)),
		Rooms:          make([]Room, 4),
		ConsumedEnergy: consumedEnergy,
	}

	b.Rooms[0] = Room{Stack: collections.NewStackFilled(dot2nothing([]rune(room1))), Capacity: len(room1)}
	b.Rooms[1] = Room{Stack: collections.NewStackFilled(dot2nothing([]rune(room2))), Capacity: len(room1)}
	b.Rooms[2] = Room{Stack: collections.NewStackFilled(dot2nothing([]rune(room3))), Capacity: len(room1)}
	b.Rooms[3] = Room{Stack: collections.NewStackFilled(dot2nothing([]rune(room4))), Capacity: len(room1)}

	return b
}

func dot2zero(chars []rune) []rune {
	for i := 0; i < len(chars); i++ {
		if chars[i] == '.' {
			chars[i] = 0
		}
	}
	return chars
}

func dot2nothing(chars []rune) []rune {
	for i := 0; i < len(chars); i++ {
		if chars[i] == '.' {
			return chars[0:i]
		}
	}
	return chars
}

func (b Building) IsSorted() bool {
	for iRoom := 0; iRoom < 4; iRoom++ {
		room := b.Rooms[iRoom]
		if !room.ContainsOnly(rune('A'+iRoom)) || room.Len() != room.Capacity {
			return false
		}
	}

	return true
}

func (b Building) Clone() Building {
	b2 := Building{
		Hallway:        slices.Clone(b.Hallway),
		Rooms:          make([]Room, 4),
		ConsumedEnergy: b.ConsumedEnergy,
		Previous:       &b,
	}

	b2.Rooms[0] = b.Rooms[0].Clone()
	b2.Rooms[1] = b.Rooms[1].Clone()
	b2.Rooms[2] = b.Rooms[2].Clone()
	b2.Rooms[3] = b.Rooms[3].Clone()

	return b2
}
func chZeroToDot(ch rune) rune {
	if ch == 0 {
		return '.'
	}

	return ch
}

func getChOrDot(chars []rune, i int) string {
	if i < 0 || i >= len(chars) {
		return "."
	}

	return string(chZeroToDot(chars[i]))
}

func (b Building) String() string {
	sb := strings.Builder{}

	sb.WriteString("#############\n")
	sb.WriteString("#")

	// hallway
	for _, ch := range b.Hallway {
		sb.WriteRune(chZeroToDot(ch))
	}
	sb.WriteString("#\n")

	// rooms
	for i := b.Rooms[0].Capacity - 1; i >= 0; i-- {
		if i == b.Rooms[0].Capacity-1 {
			sb.WriteString("###")
		} else {
			sb.WriteString("  #")
		}

		for _, room := range b.Rooms {
			sb.WriteString(getChOrDot(room.PeekAll(), i) + "#")
		}

		if i == b.Rooms[0].Capacity-1 {
			sb.WriteString("##\n")
		} else {
			sb.WriteString("\n")
		}
	}

	sb.WriteString("  #########")

	return sb.String()
}

func stepEnergy(ch rune) int {
	switch ch {
	case 'A':
		return 1
	case 'B':
		return 10
	case 'C':
		return 100
	case 'D':
		return 1000
	}

	panic("Unknown ch")
}

func PrintMoves(b *Building) {
	if b.Previous != nil {
		PrintMoves(b.Previous)
	}

	fmt.Println(b.String())
	fmt.Printf("======================================%v\n", b.ConsumedEnergy)
}

func Move(building Building, lowestEnergy *int) ([]Building, *Building) {
	// check best energy
	if building.ConsumedEnergy > *lowestEnergy {
		return nil, nil
	}

	// check final state
	if building.IsSorted() {
		*lowestEnergy = building.ConsumedEnergy
		return nil, &building
	}

	var buildings []Building

	hallway := building.Hallway

	// try every ch in all rooms and move it to hallway
	for iRoom := 0; iRoom < 4; iRoom++ {
		room := building.Rooms[iRoom]

		if room.Empty() {
			continue
		}

		// do not move properly positioned ch into hallway
		if room.ContainsOnly(rune('A' + iRoom)) {
			continue
		}

		stepsUp := room.Capacity + 1 - room.Len()

		// index of room in hallway
		iRoomHallway := iRoom*2 + 2

		// try all available hallway locations
		// left
		for iHallway := iRoomHallway - 1; iHallway >= 0; iHallway-- {
			// invalid hallway positions (room entrances)
			if iHallway == 2 || iHallway == 4 || iHallway == 6 || iHallway == 8 {
				continue
			}

			// check occupancy
			if hallway[iHallway] != 0 {
				// can not move here and further -> stop
				break
			}

			// calculate energy
			stepsTotal := stepsUp + iRoomHallway - iHallway
			energy := building.ConsumedEnergy + stepsTotal*stepEnergy(room.Peek())

			// check best energy
			if energy > *lowestEnergy {
				continue
			}

			// move to that position creating new state
			b := building.Clone()
			// remove ch from room
			ch := b.Rooms[iRoom].Pop()
			// add it to hallway
			b.Hallway[iHallway] = ch
			// energy
			b.ConsumedEnergy = energy

			// store it
			buildings = append(buildings, b)
		}

		// right
		for iHallway := iRoomHallway + 1; iHallway < len(hallway); iHallway++ {
			// invalid hallway positions (room entrances)
			if iHallway == 2 || iHallway == 4 || iHallway == 6 || iHallway == 8 {
				continue
			}

			// check occupancy
			if hallway[iHallway] != 0 {
				// can not move here and further -> stop
				break
			}

			// calculate energy
			stepsTotal := stepsUp + iHallway - iRoomHallway
			energy := building.ConsumedEnergy + stepsTotal*stepEnergy(room.Peek())

			// check best energy
			if energy > *lowestEnergy {
				continue
			}

			// move to that position creating new state
			b := building.Clone()
			// remove ch from room
			ch := b.Rooms[iRoom].Pop()
			// add it to hallway
			b.Hallway[iHallway] = ch
			// energy
			b.ConsumedEnergy = energy

			// store it
			buildings = append(buildings, b)
		}
	}

	// try every ch in hallway and move it to room
HALLWAY:
	for iHallway, ch := range hallway {
		// invalid hallway positions (room entrances)
		if iHallway == 2 || iHallway == 4 || iHallway == 6 || iHallway == 8 {
			continue
		}

		// check occupancy
		if ch == 0 {
			continue
		}

		// target room
		iRoom := int(ch - 'A')
		room := building.Rooms[iRoom]

		// check room eligibility
		if !room.CanAccept(ch) {
			continue
		}
		// room is eligible

		// index of room in hallway
		iRoomHallway := iRoom*2 + 2

		// check path to room
		step := utils.Signum(iRoomHallway - iHallway)

		// go left or right
		for iiHallway := iHallway + step; utils.Abs(iiHallway-iRoomHallway) > 0; iiHallway += step {
			// invalid hallway positions (room entrances)
			if iiHallway == 2 || iiHallway == 4 || iiHallway == 6 || iiHallway == 8 {
				continue
			}

			ch := hallway[iiHallway]

			// check occupancy
			if ch != 0 {
				// occupied, can not move
				continue HALLWAY
			}
		}

		// calculate energy
		stepsSide := utils.Abs(iHallway - iRoomHallway)
		stepsDown := room.Capacity - room.Len()
		stepsTotal := stepsSide + stepsDown
		energy := building.ConsumedEnergy + stepsTotal*stepEnergy(ch)

		// check best energy
		if energy > *lowestEnergy {
			continue
		}

		// move into room creating new state
		b := building.Clone()
		// remove it from hallway
		b.Hallway[iHallway] = 0
		// add ch to room
		b.Rooms[iRoom].Push(ch)
		// energy
		b.ConsumedEnergy = energy

		// store it
		buildings = append(buildings, b)
	}

	return buildings, nil
}

var metricGlobal = utils.NewMetric("Global")
var metricWinner = utils.NewMetric("Winner")
var metrics = utils.Metrics{metricGlobal, metricWinner}

func Sort(building Building) (int, *Building) {
	metricWinner.Enabled = metricGlobal.Enabled

	lowestEnergy := math.MaxInt
	buildings := []Building{building}
	var totalWinner *Building

	for len(buildings) > 0 {
		metricGlobal.TickCurrent(500_000, len(buildings))

		b := buildings[0]
		buildings = slices.RemoveUnordered(buildings, 0)

		nextBuildings, currentWinner := Move(b, &lowestEnergy)
		buildings = append(buildings, nextBuildings...)

		if currentWinner != nil {
			totalWinner = currentWinner
			metricWinner.Tick(100)
		}
	}
	metricGlobal.Finished()
	metricWinner.Finished()

	return lowestEnergy, totalWinner
}
