package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type color struct {
	red, green, blue byte
}

func setPixels(x, y int, colorToSet color, pixels []byte) {

	pixelIndex := (channelDepth*y + x) * channelDepth

	if pixelIndex >= 0 && pixelIndex < len(pixels)-channelDepth {
		pixels[pixelIndex] = colorToSet.red
		pixels[pixelIndex+1] = colorToSet.green
		pixels[pixelIndex+2] = colorToSet.blue
	}
}

func showPixels(pixels []byte, renderer *sdl.Renderer, texture *sdl.Texture) {

	texture.Update(nil, pixels, screenPitch)
	renderer.Copy(texture, nil, nil)
	renderer.Present()
	sdl.Delay(16)
}
