package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	// INIT VARS
	colorWhite := color{255, 255, 255}
	colorRed := color{255, 0, 0}
	colorTeal := color{0, 255, 242}

	// FOR THE BACKGROUND
	frequency := float32(0.015)
	gain := float32(0.2)
	lacunarity := float32(10)
	octaves := 3

	var state gameState = paused

	var frameStart time.Time
	var elapsedTime float32

	window, renderer, texture := setupSDL(windowTitle)
	defer sdl.Quit()
	defer window.Destroy()
	defer renderer.Destroy()
	defer texture.Destroy()

	// MAKE PIXELS - AN ARRAY OF BYTES - THIS IS LIKE A malloc IN C
	pixels := make([]byte, noOfPixels)

	playerPaddle := paddle{position{paddleStartOffset, paddleStartY}, paddleXSize, paddleYSize, paddleStartSpeed, paddleStartScore, colorWhite}
	ball := ball{position{ballStartX, ballStartY}, ballSize, ballStartXSpeed, ballStartYSpeed, colorWhite}

	aiPaddle := paddle{position{winWidth - paddleStartOffset, paddleStartY}, paddleXSize, paddleYSize, paddleStartSpeed, paddleStartScore, colorWhite}

	// KEYBOARD STATE ARRAY
	keyState := sdl.GetKeyboardState()

	// INFINITE LOOP POLLING FOR EVENTS - NEEDS TO CONTAIN DRAWING CODE
	for {
		frameStart = time.Now()

		// POLL FOR EVENTS - LOOK FOR QUIT EVENT
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		// UPDATE GAME STATE
		if state == play {
			if playerPaddle.score == maxScore || aiPaddle.score == maxScore {
				state = paused
			} else {
				playerPaddle.update(keyState, elapsedTime)
				aiPaddle.aiUpdate(&ball)
				ball.update(&playerPaddle, &aiPaddle, elapsedTime)
			}

		} else if state == paused {
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				playerPaddle.score = 0
				aiPaddle.score = 0
				state = play
			}
		}

		// DRAW NOISE BACKGROUND
		makeNoise(pixels, frequency, gain, lacunarity, octaves, colorRed, colorTeal)

		// DRAW BALL AND playerPaddle and aiPaddle
		playerPaddle.draw(pixels)
		aiPaddle.draw(pixels)
		ball.draw(pixels)

		drawPixels(pixels, window, renderer, texture)

		elapsedTime = float32(time.Since(frameStart).Seconds()) * 1000.0 // IN MILISECONDS

		// FRAMERATE FIXING
		if elapsedTime < frameTime {
			sdl.Delay(uint32(frameTime) - uint32(elapsedTime))
			elapsedTime = float32(time.Since(frameStart).Seconds()) * 1000.0
		}
	}
}
