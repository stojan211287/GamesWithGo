package main

import (
	"runtime"
	"sync"
)

func rescaleAndDraw(rawNoise []float32, minNoise, maxNoise float32, colorGradient []color, pixels []byte) {

	scale := 255.0 / (maxNoise - minNoise)
	offset := minNoise * scale

	for i := range rawNoise {

		rawNoise[i] = rawNoise[i]*scale - offset

		color := colorGradient[clamp(0, 255, int(rawNoise[i]))]

		pixels[i*channelDepth] = color.red
		pixels[i*channelDepth+1] = color.green
		pixels[i*channelDepth+2] = color.blue
	}
}

func turbulenceNoise(x, y, frequency, lacunarity, gain float32, octaves int) float32 {
	var sum float32

	amplitude := float32(1.0)

	for i := 0; i < octaves; i++ {
		f := snoise2(x*frequency, y*frequency) * amplitude

		if f < 0.0 {
			f = -1.0 * f
		}

		sum += f
	}
	return sum
}

func fractionalBrownianMotion(x, y, frequency, lacunarity, gain float32, octaves int) float32 {

	var sum float32

	amplitude := float32(1.0)

	for i := 0; i < octaves; i++ {

		sum += snoise2(x*frequency, y*frequency) * amplitude

		frequency *= lacunarity
		amplitude *= gain
	}

	return sum
}

func makeNoise(pixels []byte, frequency, gain, lacunarity float32, octaves int) {

	noise := make([]float32, windowHeight*windowWidth)

	numRoutines := runtime.NumCPU()
	routineBatchSize := len(noise) / numRoutines

	var minNoise float32
	var maxNoise float32

	// MAKE WAITGROUP FOR GOROUTINES
	var waitGroup sync.WaitGroup
	waitGroup.Add(numRoutines)

	// USE GOROUTINES TO SAMPLE NOISE
	for i := 0; i < numRoutines; i++ {
		go func(i int) {
			defer waitGroup.Done()

			var routineMin float32
			var routineMax float32
			var mutex = &sync.Mutex{}

			startIndex := i * routineBatchSize
			endIndex := startIndex + routineBatchSize - 1

			for j := startIndex; j < endIndex; j++ {
				x := j % windowWidth
				y := (j - x) / windowHeight
				noise[j] = turbulenceNoise(float32(x), float32(y), frequency, lacunarity, gain, octaves)

				if j == 0 {
					routineMin = noise[j]
					routineMax = noise[j]
				}
				if noise[j] < routineMin {
					routineMin = noise[j]
				} else if noise[j] > routineMax {
					routineMax = noise[j]
				}
			}

			// MODIFY GLOBAL MIN AND MAX
			mutex.Lock()
			if routineMin < minNoise {
				minNoise = routineMin
			}
			if routineMax > maxNoise {
				maxNoise = routineMax
			}
			mutex.Unlock()

		}(i)
	}

	// WAIT UNTIL ALL GOROUTINES DONE
	waitGroup.Wait()

	gradient := getColorGradient(color{255, 0, 0}, color{0, 255, 0})
	rescaleAndDraw(noise, minNoise, maxNoise, gradient, pixels)
}

func lerp(b1 byte, b2 byte, percent float32) byte {
	return byte(float32(b1) + percent*(float32(b2)-float32(b1)))
}

func colorLerp(c1, c2 color, percent float32) color {
	return color{lerp(c1.red, c2.red, percent), lerp(c1.blue, c2.blue, percent), lerp(c1.green, c2.green, percent)}
}

func getColorGradient(c1, c2 color) []color {

	colorGradient := make([]color, 256)

	for i := range colorGradient {
		percent := float32(i) / float32(255)
		colorGradient[i] = colorLerp(c1, c2, percent)
	}

	return colorGradient
}

func clamp(min, max, value int) int {
	if value < min {
		value = min
	} else if value > max {
		value = max
	}
	return value
}
