package day_22

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
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

type Building2 struct {
	// 00-10 = hallway
	// 11-12 = room1
	// 13-14 = room2
	// 15-16 = room3
	// 17-18 = room4
	data           []rune
	ConsumedEnergy int
	Previous       *Building2
}

func NewBuilding2(room1, room2, room3, room4 string) Building2 {
	b := Building2{
		data:           make([]rune, 19),
		ConsumedEnergy: 0,
	}

	utils.Copy([]rune(room1), b.Room(0))
	utils.Copy([]rune(room2), b.Room(1))
	utils.Copy([]rune(room3), b.Room(2))
	utils.Copy([]rune(room4), b.Room(3))

	return b
}
func NewBuilding2Full(hallway, room1, room2, room3, room4 string, consumedEnergy int) Building2 {
	b := Building2{
		data:           make([]rune, 19),
		ConsumedEnergy: consumedEnergy,
	}

	utils.Copy([]rune(hallway), b.Hallway())

	utils.Copy([]rune(room1), b.Room(0))
	utils.Copy([]rune(room2), b.Room(1))
	utils.Copy([]rune(room3), b.Room(2))
	utils.Copy([]rune(room4), b.Room(3))

	for i := 0; i < len(b.data); i++ {
		if b.data[i] == '.' {
			b.data[i] = 0
		}
	}

	return b
}

func (b Building2) Hallway() []rune {
	return b.data[0:11]
}

func (b Building2) Room(i int) Room {
	return b.data[11+i*2 : 11+i*2+2]
}

func (b Building2) IsSorted() bool {
	for iRoom := 0; iRoom < 4; iRoom++ {
		room := b.Room(iRoom)
		if int(room[0]) != ('A'+iRoom) || int(room[1]) != ('A'+iRoom) {
			return false
		}
	}

	return true
}

func (b Building2) Clone() Building2 {
	return Building2{
		data:           utils.ShallowCopy(b.data),
		ConsumedEnergy: b.ConsumedEnergy,
		Previous:       &b,
	}
}
func chToString(ch rune) rune {
	if ch == 0 {
		return '.'
	}

	return ch
}

func (b Building2) String() string {
	sb := strings.Builder{}

	sb.WriteString("#############\n")
	sb.WriteString("#")

	// hallway
	for _, ch := range b.Hallway() {
		sb.WriteRune(chToString(ch))
	}
	sb.WriteString("#\n")

	// rooms
	sb.WriteString("###")
	for i := 0; i < 4; i++ {
		room := b.Room(i)
		sb.WriteString(string(chToString(room[1])) + "#")
	}
	sb.WriteString("##\n")

	sb.WriteString("  #")
	for i := 0; i < 4; i++ {
		room := b.Room(i)
		sb.WriteString(string(chToString(room[0])) + "#")
	}
	sb.WriteString("\n")
	sb.WriteString("  #########")

	return sb.String()
}

type Room []rune

func (r Room) Length() int {
	if r[1] != 0 {
		return 2
	}

	if r[0] != 0 {
		return 1
	}

	return 0
}

func (r Room) Empty() bool {
	return r.Length() == 0
}

func (r Room) Push(ch rune) {
	r[r.Length()] = ch
}

func (r Room) Pop() rune {
	ch := r[r.Length()-1]
	r[r.Length()-1] = 0
	return ch
}

func (r Room) Peek() rune {
	return r[r.Length()-1]
}

func (r Room) ContainsOnly(ch rune) bool {
	if r[1] != 0 && r[1] != ch {
		return false
	}

	if r[0] != 0 && r[0] != ch {
		return false
	}

	return r.Length() > 0
}

func (r Room) EligibleFor(ch rune) bool {
	length := r.Length()

	if length == 0 {
		return true
	}

	if length == 2 {
		return false
	}

	return r[0] == ch
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

func PrintMoves(b *Building2) {
	if b.Previous != nil {
		PrintMoves(b.Previous)
	}

	fmt.Println(b.String())
	fmt.Printf("======================================%v\n", b.ConsumedEnergy)
}

func Move(building Building2, lowestEnergy *int) ([]Building2, *Building2) {
	// check best energy
	if building.ConsumedEnergy > *lowestEnergy {
		return nil, nil
	}

	// check final state
	if building.IsSorted() {
		*lowestEnergy = building.ConsumedEnergy
		return nil, &building
	}

	var buildings []Building2

	hallway := building.Hallway()

	// try every ch in all rooms and move it to hallway
	for iRoom := 0; iRoom < 4; iRoom++ {
		room := building.Room(iRoom)

		if room.Empty() {
			continue
		}

		// do not move properly positioned ch into hallway
		if room.ContainsOnly(rune('A' + iRoom)) {
			continue
		}

		stepsUp := 3 - room.Length()

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
			ch := b.Room(iRoom).Pop()
			// add it to hallway
			b.Hallway()[iHallway] = ch
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
			ch := b.Room(iRoom).Pop()
			// add it to hallway
			b.Hallway()[iHallway] = ch
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
		room := building.Room(iRoom)

		// check room eligibility
		if !room.EligibleFor(ch) {
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
		stepsDown := 2 - room.Length()
		stepsTotal := stepsSide + stepsDown
		energy := building.ConsumedEnergy + stepsTotal*stepEnergy(ch)

		// check best energy
		if energy > *lowestEnergy {
			continue
		}

		// move into room creating new state
		b := building.Clone()
		// remove it from hallway
		b.Hallway()[iHallway] = 0
		// add ch to room
		b.Room(iRoom).Push(ch)
		// energy
		b.ConsumedEnergy = energy

		// store it
		buildings = append(buildings, b)
	}

	return buildings, nil
}

var metric = utils.NewMetric("Global")
var metricSolution = utils.NewMetric("Winner")

func Sort(building Building2) (int, *Building2) {
	metricSolution.Enabled = metric.Enabled

	lowestEnergy := math.MaxInt
	buildings := []Building2{building}
	var totalWinner *Building2

	for len(buildings) > 0 {
		metric.TickCurrent(500_000, len(buildings))

		b := buildings[0]
		buildings = utils.RemoveUnordered(buildings, 0)

		nextBuildings, currentWinner := Move(b, &lowestEnergy)
		buildings = append(buildings, nextBuildings...)

		if currentWinner != nil {
			totalWinner = currentWinner
			metricSolution.Tick(100)
		}
	}
	metric.Finished()
	metricSolution.Finished()

	return lowestEnergy, totalWinner
}
