package main

// add a
import (
	"fmt"
	"galaxy/ai"
	"galaxy/camera"
	"galaxy/data"
	"galaxy/keybindings"
	texs "galaxy/texTools"
	"galaxy/universe"
	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
	"os"
)

const (
	Title  = "2d Shapes"
	Width  = 1000
	Height = 900
)

var (
	texture    gl.Uint
	rotx, roty gl.Float
	ambient    []gl.Float = []gl.Float{0.5, 0.5, 0.5, 1}
	diffuse    []gl.Float = []gl.Float{1, 1, 1, 1}
	lightpos   []gl.Float = []gl.Float{-5, 5, 10, 0}

	Objects []*data.GameObject
)

func main() {
	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.Terminate()

	glfw.OpenWindowHint(glfw.WindowNoResize, 0)

	if err := glfw.OpenWindow(Width, Height, 0, 0, 0, 0, 16, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.CloseWindow()

	glfw.SetSwapInterval(60)
	glfw.SetWindowTitle(Title)

	if err := gl.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "gl: %s\n", err)
	}

	keybindings.BindKeyboard()

	gl.ClearColor(0, 0, 0, 1)

	GameLoop()

}

func GameLoop() {

	t := 0.0
	dt := 0.1

	CreateUniverse()

	camera.PositionCamera(0, 0, 1000, 900)

	for glfw.WindowParam(glfw.Opened) == 1 {

		t += dt

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.Clear(gl.DEPTH_BUFFER_BIT | gl.LIGHTING_BIT)

		for i := range universe.Systems {
			for j := range universe.Systems[i].Planets {
				BufferPlanet(universe.Systems[i].Planets[j])
				ai.PlanetOrbit(universe.Systems[i].Planets[j], t)
			}
			for j := range universe.Systems[i].Stars {
				BufferStar(universe.Systems[i].Stars[j])
				ai.StarStationary(universe.Systems[i].Stars[j], t)
			}
		}

		glfw.SwapBuffers()
	}
}

func CreateUniverse() {
	universe.LoadTextures()
	universe.SetSeed(10)
	universe.PlaceSystems(10000, 10)
}

func BufferPlanet(o *data.Planet) {
	texs.DrawPlanetTex(o)
}

func BufferStar(o *data.Star) {
	texs.DrawStarTex(o)
}

func destroyScene() {
	gl.DeleteTextures(1, &texture)
}
