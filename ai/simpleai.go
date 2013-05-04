package ai

import (
	//"fmt"
	"galaxy/data"
	gl "github.com/chsc/gogl/gl21"
	"math"
)

type AIer interface {
	Move()
	AbsMove()
	Rotate()
	AbsRotate()
	GetAIParams()
}

const (
	Tolerance = 10
	Speed     = 3
)

type Stage struct {
	MinX float32
	MinY float32

	MaxX float32
	MaxY float32

	Actors []int
}

var (
	State = Stage{0, 0, 1000, 900, make([]int, 10)}
)

// TODO refactor to use gameobject instead of planet

/* Updates o's position using its movement function.

The movemet function looks something like 
	(x,y) = (f(t),g(t)) where for elliptic orbits

	f = Acos(t)
	g = Bsin(t)
*/
func PlanetOrbit(o *data.Planet, t float64) {
	A := o.OrbitParam1
	B := o.OrbitParam2
	C := o.OrbitParam3

	x0 := o.OrbitParam4
	y0 := o.OrbitParam5

	o.AbsMove(gl.Float(A*math.Sin(t*C)+x0), gl.Float(B*math.Cos(t*C)+y0))
}

func StarStationary(o *data.Star, t float64) {
	// Do animation stuff if needed, currently don't have animated sprites
}
