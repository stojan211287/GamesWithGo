package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	colorWhite := color{255, 255, 255}

	// INIT THE EVENT SYSTEM
	err := sdl.Init(sdl.INIT_EVERYTHING)
	checkError(err)
	defer sdl.Quit()

	// MAKE WINDOW
	window, err := sdl.CreateWindow("Testin SDL2, FINALLY!", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(winHeight), int32(winWidth), sdl.WINDOW_SHOWN)

	checkError(err)
	defer window.Destroy()

	// MAKE RENDERER
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	checkError(err)
	defer renderer.Destroy()

	// MAKE TEXTURE - FROM RENDERER
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	checkError(err)
	defer texture.Destroy()

	// MAKE PIXELS - AN ARRAY OF BYTES - THIS IS LIKE A malloc IN C
	pixels := make([]byte, noOfPixels)

	playerPaddle := paddle{position{paddleStartOffset, paddleStartY}, 20, 100, 0, 0, colorWhite}
	ball := ball{position{ballStartX, ballStartY}, ballSize, 1, 2, colorWhite}

	aiPaddle := paddle{position{winWidth - paddleStartOffset, paddleStartY}, 20, 100, 0, 0, colorWhite}

	// KEYBOARD STATE ARRAY
	keyState := sdl.GetKeyboardState()

	// INFINITE LOOP POLLING FOR EVENTS - NEEDS TO CONTAIN DRAWING CODE
	for {

		// POLL FOR EVENTS - LOOK FOR QUIT EVENT
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		// UPDATE GAME STATE
		playerPaddle.update(keyState)
		// UPDATE AMAZING AI
		aiPaddle.aiUpdate(&ball)
		// UPDATE BALL
		ball.update(&playerPaddle, &aiPaddle)

		// CLEAR SCREEN
		clearScreen(pixels)

		// DRAW BALL AND playerPaddle and aiPaddle
		playerPaddle.draw(pixels)
		aiPaddle.draw(pixels)
		ball.draw(pixels)

		// UPDATE TEXTURE WITH NEW PIXELS ARRAY
		texture.Update(nil, pixels, screenPitch)
		renderer.Copy(texture, nil, nil)
		renderer.Present()
	}
}
