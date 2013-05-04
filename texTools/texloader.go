package texTools

import (
	"bytes"
	"errors"
	gl "github.com/chsc/gogl/gl21"
	"image"
	"image/png"
	"io"
)

/*
Creates a gl texture, returning its id and an error.

data should be a png build using go_images.bat
*/
func CreateTextureFromBytes(data []byte) (gl.Uint, error) {
	r := bytes.NewBuffer(data)
	return createTexture(r)
}

/*
Uses the image library to turn a PNG into a gl texture
*/
func createTexture(r io.Reader) (textureId gl.Uint, err error) {
	img, err := png.Decode(r)
	if err != nil {
		return 0, err
	}

	rgbaImg, ok := img.(*image.RGBA)
	if !ok {
		return 0, errors.New("texture must be an RGBA image")
	}

	gl.GenTextures(1, &textureId)
	gl.BindTexture(gl.TEXTURE_2D, textureId)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

	// flip image: first pixel is lower left corner
	imgWidth, imgHeight := img.Bounds().Dx(), img.Bounds().Dy()
	data := make([]byte, imgWidth*imgHeight*4)
	lineLen := imgWidth * 4
	dest := len(data) - lineLen
	for src := 0; src < len(rgbaImg.Pix); src += rgbaImg.Stride {
		copy(data[dest:dest+lineLen], rgbaImg.Pix[src:src+rgbaImg.Stride])
		dest -= lineLen
	}
	gl.TexImage2D(gl.TEXTURE_2D, 0, 4, gl.Sizei(imgWidth), gl.Sizei(imgHeight), 0, gl.RGBA,
		gl.UNSIGNED_BYTE, gl.Pointer(&data[0]))

	return textureId, nil
}

func Enable2d(x1, y1, w, h gl.Double) {

	// Save a copy of the proj matrix so we can restore it
	// when we want to do more 3d rendering
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	// Set up orthographic projection
	gl.Ortho(x1, x1+w, y1, y1+h, 0.0, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	// Make sure depth testing and lighting are disabled for 2d redering
	// until we are fninished rendering in 2d
	gl.PushAttrib(gl.DEPTH_BUFFER_BIT | gl.LIGHTING_BIT)
	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.LIGHTING)
}

func Disable2d() {
	// Tear down from Enable2d() assuing the matrix stacks have not been changed between calls.
	gl.PopAttrib()
}
