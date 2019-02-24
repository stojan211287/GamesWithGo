package main

func drawNumbers(pos position, color color, size int, num int, pixels []byte) {

	numbers := [][]byte{
		{1, 1, 1,
			1, 0, 1,
			1, 0, 1,
			1, 0, 1,
			1, 1, 1},
		{1, 1, 0,
			0, 1, 0,
			0, 1, 0,
			0, 1, 0,
			1, 1, 1},
		{1, 1, 1,
			0, 0, 1,
			1, 1, 1,
			1, 0, 0,
			1, 1, 1},
		{1, 1, 1,
			0, 0, 1,
			1, 1, 1,
			0, 0, 1,
			1, 1, 1}}

	startX := int(pos.x) - (size*3)/2
	startY := int(pos.y) - (size*5)/2

	for index, value := range numbers[num] {
		if value == 1 {
			for y := startY; y < startY+size; y++ {
				for x := startX; x < startX+size; x++ {
					setPixel(x, y, color, pixels)
				}
			}
		}
		startX += size
		if (index+1)%3 == 0 {
			startY += size
			startX -= size * 3
		}
	}
}
