package texTools

import (
	//"fmt"
	"galaxy/data"
	gl "github.com/chsc/gogl/gl21"
)

/*  TODO This needs to be refactored so that everything uses gametextures interface
with rectangular textures instead of having a different type for each. */

/* Draws a texture to the screen given a pointer to is GameTexture struct */
func DrawRecTex(t *data.GameObject) {

	// Find the center point of the texture to rotate around
	xav := (t.X1 + t.X2) / 2
	yav := (t.Y1 + t.Y2) / 2

	//Translate there, rotate, translate back
	gl.MatrixMode(gl.MODELVIEW)
	gl.Translatef(xav, yav, 0)
	gl.Rotatef(gl.Float(t.Theta), 0, 0, 1)
	gl.Translatef(-xav, -yav, 0)

	//Bind our texture to be drawn by id
	gl.Color3f(1, 1, 1)
	gl.Enable(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, t.TexId)

	// Draw a rectangle with the texture stretched to the corners
	gl.Begin(gl.QUADS)

	// Stretch the texture to its 4 corners.
	gl.TexCoord2d(0, 0)
	gl.Vertex2f(t.X1, t.Y1)

	gl.TexCoord2d(0, 1)
	gl.Vertex2f(t.X1, t.Y2)

	gl.TexCoord2d(1, 1)
	gl.Vertex2f(t.X2, t.Y2)

	gl.TexCoord2d(1, 0)
	gl.Vertex2f(t.X2, t.Y1)

	gl.End()

	// Unbind the texture in case something else wants to draw
	gl.Disable(gl.TEXTURE_2D)

	// Reset the matrix
	gl.LoadIdentity()

}

func DrawPlanetTex(t *data.Planet) {

	// Find the center point of the texture to rotate around
	xav := (t.X1 + t.X2) / 2
	yav := (t.Y1 + t.Y2) / 2

	//Translate there, rotate, translate back
	gl.MatrixMode(gl.MODELVIEW)
	gl.Translatef(xav, yav, 0)
	gl.Rotatef(gl.Float(t.Theta), 0, 0, 1)
	gl.Translatef(-xav, -yav, 0)

	//Bind our texture to be drawn by id
	gl.Color3f(1, 1, 1)
	gl.Enable(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, t.TexId)

	// Draw a rectangle with the texture stretched to the corners
	gl.Begin(gl.QUADS)

	// Stretch the texture to its 4 corners.
	gl.TexCoord2d(0, 0)
	gl.Vertex2f(t.X1, t.Y1)

	gl.TexCoord2d(0, 1)
	gl.Vertex2f(t.X1, t.Y2)

	gl.TexCoord2d(1, 1)
	gl.Vertex2f(t.X2, t.Y2)

	gl.TexCoord2d(1, 0)
	gl.Vertex2f(t.X2, t.Y1)

	gl.End()

	// Unbind the texture in case something else wants to draw
	gl.Disable(gl.TEXTURE_2D)

	// Reset the matrix
	gl.LoadIdentity()

}

func DrawStarTex(t *data.Star) {

	// Find the center point of the texture to rotate around
	xav := (t.X1 + t.X2) / 2
	yav := (t.Y1 + t.Y2) / 2

	//Translate there, rotate, translate back
	gl.MatrixMode(gl.MODELVIEW)
	gl.Translatef(xav, yav, 0)
	gl.Rotatef(gl.Float(t.Theta), 0, 0, 1)
	gl.Translatef(-xav, -yav, 0)

	//Bind our texture to be drawn by id
	gl.Color3f(1, 1, 1)
	gl.Enable(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, t.TexId)

	// Draw a rectangle with the texture stretched to the corners
	gl.Begin(gl.QUADS)

	// Stretch the texture to its 4 corners.
	gl.TexCoord2d(0, 0)
	gl.Vertex2f(t.X1, t.Y1)

	gl.TexCoord2d(0, 1)
	gl.Vertex2f(t.X1, t.Y2)

	gl.TexCoord2d(1, 1)
	gl.Vertex2f(t.X2, t.Y2)

	gl.TexCoord2d(1, 0)
	gl.Vertex2f(t.X2, t.Y1)

	gl.End()

	// Unbind the texture in case something else wants to draw
	gl.Disable(gl.TEXTURE_2D)

	// Reset the matrix
	gl.LoadIdentity()

}
