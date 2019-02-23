package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func setupSDL() (*sdl.Window, *sdl.Renderer, *sdl.Texture) {

	// INIT THE EVENT SYSTEM
	err := sdl.Init(sdl.INIT_EVERYTHING)
	checkError(err)

	// MAKE WINDOW
	window, err := sdl.CreateWindow("SimplexNoiseTest", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(windowWidth), int32(windowHeight), sdl.WINDOW_SHOWN)
	checkError(err)

	// MAKE RENDERER OF WINDOW
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	checkError(err)

	// MAKE TEXTURE
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(windowWidth), int32(windowHeight))
	checkError(err)

	return window, renderer, texture
}
