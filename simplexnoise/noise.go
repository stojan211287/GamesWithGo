package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	window, renderer, texture := setupSDL()

	defer sdl.Quit()
	defer window.Destroy()
	defer renderer.Destroy()
	defer texture.Destroy()

	pixels := make([]byte, windowHeight*windowWidth*channelDepth)
	keyState := sdl.GetKeyboardState()

	multiplier := float32(1.0)

	frequency := float32(0.005)
	gain := float32(0.2)
	lacunarity := float32(3)
	octaves := 3

	// SETUP EVENT LOOP
	for {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		shiftPressed := keyState[sdl.SCANCODE_LSHIFT] != 0 || keyState[sdl.SCANCODE_RSHIFT] != 0
		frequencyKeyPressed := keyState[sdl.SCANCODE_F] != 0
		gainKeyPressed := keyState[sdl.SCANCODE_G] != 0
		lacunarityKeyPressed := keyState[sdl.SCANCODE_L] != 0
		octaveKeyPressed := keyState[sdl.SCANCODE_O] != 0

		if shiftPressed {
			multiplier *= -1.0
		}
		if frequencyKeyPressed {
			frequency = frequency + 0.001*multiplier
		}
		if octaveKeyPressed {
			octaves = octaves + int(1*multiplier)
		}
		if gainKeyPressed {
			gain = gain + 0.1*multiplier
		}
		if lacunarityKeyPressed {
			lacunarity = lacunarity + 0.5*multiplier
		}

		makeNoise(pixels, frequency, gain, lacunarity, octaves)
		showPixels(pixels, renderer, texture)

		fmt.Println("Multiplier : ", multiplier)
		fmt.Println("Gain :", gain)
		fmt.Println("Lacunarity : ", lacunarity)
		fmt.Println("Frequency : ", frequency)
		fmt.Println("Octaves : ", octaves)
	}

}
