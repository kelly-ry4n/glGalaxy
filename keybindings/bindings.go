package keybindings

import (
	"galaxy/camera"
	//"galaxy/universe"
	"github.com/jteeuwen/glfw"
)

const (
	w         = 119
	a         = 97
	s         = 115
	d         = 100
	h         = 104
	movespeed = 50
)

/*
Binds each key to an action
*/
func takeKeyAction(key int, action int) {

	if action == glfw.KeyPress {

		switch {
		// WASD change the camera postion by movespeed
		case w == key:
			camera.Y += movespeed
			camera.PositionCamera(camera.X, camera.Y, 640, 480)
		case a == key:
			camera.X -= movespeed
			camera.PositionCamera(camera.X, camera.Y, 640, 480)
		case s == key:
			camera.Y -= movespeed
			camera.PositionCamera(camera.X, camera.Y, 640, 480)
		case d == key:
			camera.X += movespeed
			camera.PositionCamera(camera.X, camera.Y, 640, 480)
			// reimpliment once finished remodeling the universe.systems
			/*		case h == key:
					camera.CameraToObj(universe.Systems[0].Planets[0])
			*/

		}

	}
}

func BindKeyboard() {
	glfw.SetCharCallback(takeKeyAction)
	glfw.Enable(glfw.KeyRepeat)
}
