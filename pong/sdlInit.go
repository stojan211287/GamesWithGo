package main

import "github.com/veandco/go-sdl2/sdl"

func setupSDL() (*sdl.Window, *sdl.Renderer, *sdl.Texture) {
	// INIT THE EVENT SYSTEM
	err := sdl.Init(sdl.INIT_EVERYTHING)
	checkError(err)

	// MAKE WINDOW
	window, err := sdl.CreateWindow("Testin SDL2, FINALLY!", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(winHeight), int32(winWidth), sdl.WINDOW_SHOWN)

	checkError(err)

	// MAKE RENDERER
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	checkError(err)

	// MAKE TEXTURE - FROM RENDERER
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	checkError(err)

	return window, renderer, texture

}
