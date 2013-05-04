package data

import (
	gl "github.com/chsc/gogl/gl21"
	//texs "galaxy/texTools"
)

type GameObject struct {
	TexId gl.Uint
	X1    gl.Float
	Y1    gl.Float
	X2    gl.Float
	Y2    gl.Float
	Theta gl.Float
}

type Planet struct {
	GameObject
	OrbitParam1 float64
	OrbitParam2 float64
	OrbitParam3 float64
	OrbitParam4 float64
	OrbitParam5 float64
}

type Star struct {
	GameObject
}

const (
	Tolerance = 10
)

// Returns true if (X,Y) is contained within o's blit rectangle
// else false
func (o *GameObject) IsOnLocation(X, Y float32) bool {
	oxav := float32((o.X2-o.X1)/2 + o.X1)
	oyav := float32((o.Y2-o.Y1)/2 + o.Y1)

	var out bool

	if abs(oxav-X) < Tolerance && abs(oyav-Y) < Tolerance {
		out = true
	} else {
		out = false
	}
	return out
}

func (t *GameObject) Move(x, y gl.Float) {
	t.X1 += x
	t.X2 += x

	t.Y1 += y
	t.Y2 += y
}

func (o *GameObject) AbsMove(x, y gl.Float) {
	w := o.X2 - o.X1
	h := o.Y2 - o.Y1

	o.X1 = x
	o.Y1 = y

	o.X2 = o.X1 + w
	o.Y2 = o.Y1 + h

}

func (t *GameObject) Stretch(x, y gl.Float) {
	t.X2 += x
	t.Y2 += y
}

func (t *GameObject) Rotate(theta gl.Float) {
	t.Theta += theta
}

func (t *GameObject) AbsRotate(theta gl.Float) {
	t.Theta = theta
}
