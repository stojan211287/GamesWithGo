package main

type color struct {
	r, g, b byte
}

type position struct {
	x, y int
}

type screenEdge int

const (
	top screenEdge = iota + 1
	bottom
	left
	right
	none
)

func setPixel(x, y int, col color, pixels []byte) {

	indexOfChange := (y*winWidth + x) * channelDepth

	if (indexOfChange < len(pixels)-channelDepth) && (indexOfChange >= 0) {
		pixels[indexOfChange] = col.r
		pixels[indexOfChange+1] = col.g
		pixels[indexOfChange+2] = col.b
	}
}

func clearScreen(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}
