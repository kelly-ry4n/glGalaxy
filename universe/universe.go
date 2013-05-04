package universe

import (
	//"fmt"
	"galaxy/data"
	gl "github.com/chsc/gogl/gl21"
	"math/rand"
)

var (
	Objects             []data.GameObject
	Systems             []System
	ActiveSystems       []*System
	ActivePlanetObjects []*data.GameObject
	ActivePlanets       []data.Planet

	NumSys  int = 0
	SysSize int = 5000
)

type System struct {
	X          int
	Y          int
	NumStars   int
	NumPlanets int
	Planets    []*data.Planet
	Stars      []*data.Star
}

func SetSeed(s int64) {
	rand.Seed(s)
}

// Creates a new system using (temorarily) hard-coded constants 
func NewSystem(X, Y int) System {
	numStars := 1
	numPlanets := rand.Int() % 12

	// Make the slice of Planet

	planets := make([]*data.Planet, numPlanets+1)

	// TODO get rid of hard coded constants
	for i := 0; i <= numPlanets; i++ {
		orbitSizeX := rand.Float64()*4000 + 2000
		orbitSizeY := rand.Float64()*4000 + 2000
		orbitSpeed := rand.Float64() / ((orbitSizeX + orbitSizeY) / 10)

		size := gl.Float(rand.Float64() * 100)

		planets[i] = &data.Planet{data.GameObject{PlanetTextures[0], 0, 0, size, size, 0},
			orbitSizeX, orbitSizeY, orbitSpeed, float64(X), float64(Y)}
	}

	// Make the slice of Stars for each system
	stars := make([]*data.Star, numStars+1)

	for i := 0; i <= numStars; i++ {
		stars[i] = &data.Star{data.GameObject{PlanetTextures[0], gl.Float(X), gl.Float(Y), gl.Float(X + 300), gl.Float(Y + 300), 0}}
	}

	return System{X, Y, numStars, numPlanets, planets, stars}
}

// Adds numSys random systems of maximum size size on elliptical orbits
// to the global state.
func PlaceSystems(size, numSys int) {
	systems := []System{}

	for i := 0; i <= numSys; i++ {
		systems = append(systems, NewSystem(rand.Int()%size, rand.Int()%size))
	}

	Systems = systems

}
