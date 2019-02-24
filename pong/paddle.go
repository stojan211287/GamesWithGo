package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type paddle struct {
	position
	width  int
	height int
	speed  float32
	score  int
	color  color
}

func (paddle *paddle) update(keyStateArray []uint8, elapsedTime float32) {

	upPressed := keyStateArray[sdl.SCANCODE_UP] != 0
	downPressed := keyStateArray[sdl.SCANCODE_DOWN] != 0

	if upPressed || downPressed {
		screenHit := paddle.hitsScreen()
		switch screenHit {
		case top:
			paddle.y = paddle.height/2 + paddleScreenBounce
		case bottom:
			paddle.y = winHeight - paddle.height/2 - paddleScreenBounce
		default:
			// CONTROLS ARE REVERSED BECAUSE OF THE IMAGE COORDINATE SYSTEM
			if upPressed {
				paddle.y -= int(paddle.speed * elapsedTime)
			}
			if downPressed {
				paddle.y += int(paddle.speed * elapsedTime)
			}
		}
	}
}

func (paddle *paddle) hitsScreen() screenEdge {

	if paddle.y+paddle.height/2 >= winHeight {
		return bottom
	}
	if paddle.y <= paddle.height/2 {
		return top
	}
	return none
}

func (paddle *paddle) aiUpdate(ball *ball) {
	paddle.y = ball.y
}

func (paddle *paddle) draw(pixels []byte) {

	startDrawX := paddle.x - paddle.width/2
	startDrawY := paddle.y - paddle.height/2

	for y := 0; y < paddle.height; y++ {
		for x := 0; x < paddle.width; x++ {
			setPixel(startDrawX+x, startDrawY+y, paddle.color, pixels)
		}
	}

	numX := lerp(float32(paddle.x), float32(winHeight/2), 0.2)
	numY := scoreYOffset

	scorePosition := position{int(numX), numY}

	drawNumbers(scorePosition, paddle.color, 10, paddle.score, pixels)
}

func (paddle *paddle) hitByBall(ball *ball) bool {

	paddlesRightEdge := paddle.x + paddle.width/2
	paddlesLeftEdge := paddle.x - paddle.width/2

	ballsLeftTip := ball.getLeftTip()
	ballsRightTip := ball.getRightTip()

	ballBetweenTopAndBottom := ball.y >= paddle.y-paddle.height/2 && ball.y <= paddle.y+paddle.height/2

	if ballBetweenTopAndBottom {
		if ballsLeftTip <= paddlesRightEdge && ballsRightTip >= paddlesLeftEdge ||
			ballsRightTip <= paddlesLeftEdge && ballsLeftTip >= paddlesRightEdge {
			return true
		}
	}
	return false
}

func (paddle *paddle) upScore() {
	paddle.score++
}
