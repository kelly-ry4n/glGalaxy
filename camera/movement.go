package camera

import (
	// "galaxy/data"
	texs "galaxy/texTools"
	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
	"strconv"
)

var (
	X      gl.Double
	Y      gl.Double
	Width  gl.Sizei
	Height gl.Sizei
)

func PositionCamera(x, y gl.Double, width, height gl.Sizei) {
	texs.Disable2d()
	texs.Enable2d(x, y, 1000, 900)

	gl.Viewport(0, 0, 1000, 900)
	gl.LoadIdentity()

	title := strconv.FormatInt(int64(x), 10) + " , " + strconv.FormatInt(int64(y), 10)

	glfw.SetWindowTitle(title)

	X = x
	y = Y
	Width = width
	Height = height

	//UpdateActiveSystems()

}

/* Fix this once universe refactor is finished
func CameraToObj(o data.GameObject) {
	PositionCamera(gl.Double(o.X1-300),
		gl.Double(o.Y1-300), 1000, 900)

	//UpdateActiveSystems()
}*/
