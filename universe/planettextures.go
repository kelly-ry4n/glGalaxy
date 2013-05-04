package universe

import (
	"fmt"
	"galaxy/TEXTURES"
	texs "galaxy/texTools"
	gl "github.com/chsc/gogl/gl21"
)

var PlanetTextures []gl.Uint

/* Temorary function which loads up the planet texture. I really nead an artist.
If you're an artist and are also reading this code, please get in touch */
func LoadTextures() {
	ID, err := texs.CreateTextureFromBytes(TEXTURES.Planet_png)

	if err != nil {
		fmt.Println("TEXTURE LOAD ERROR", err)
	}

	PlanetTextures = append(PlanetTextures, ID)
}
