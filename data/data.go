package data

import (
	"math"
)

type Instruction struct {
	X      float32
	Y      float32
	Param  float32
	Action string
}

type InstrStack struct {
	top  *InstrElement
	size int
}

type InstrElement struct {
	value Instruction
	next  *InstrElement
}

// Return the stack's length
func (s *InstrStack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *InstrStack) Push(value Instruction) {
	s.top = &InstrElement{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *InstrStack) Pop() (value Instruction) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return Instruction{}
}

func DirectionBetween(x1, y1, x2, y2 float32) (x, y float32) {
	tx := float64(x2 - x1)
	ty := float64(y2 - y1)

	mag := math.Sqrt(math.Pow(tx, 2) + math.Pow(ty, 2))

	x = float32(tx / mag)
	y = float32(ty / mag)

	return
}

func abs(a float32) float32 {
	if a < 0 {
		a = -a
	}
	return a
}
