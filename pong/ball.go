package main

type ball struct {
	// THIS IS SIMPLE INHERITANCE OF STRUCTS IN GO - NOW YOU CAN ACCESS ball.x INSTEAD OF ball.position.x
	position
	radius int
	xv     int
	yv     int
	color  color
}

func (ball *ball) draw(pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(ball.x+x, ball.y+y, ball.color, pixels)
			}
		}
	}
}

func (ball *ball) update(leftPaddle *paddle, rightPaddle *paddle) {

	ball.x += ball.xv
	ball.y += ball.yv

	switch ball.hitsScreen() {
	case left:
		ball.newKickoff()
	case right:
		ball.newKickoff()
	case top:
		ball.reverseY()
	case bottom:
		ball.reverseY()
	}

	switch leftPaddle.hitByBall(ball) {
	case ballHitsPaddle:
		ball.reverseX()
	}

	switch rightPaddle.hitByBall(ball) {
	case ballHitsPaddle:
		ball.reverseX()
	}
}

func (ball *ball) reverseX() {
	ball.xv = -ball.xv
}

func (ball *ball) reverseY() {
	ball.yv = -ball.yv
}

func (ball *ball) newKickoff() {
	ball.x = ballStartX
	ball.y = ballStartY
}

func (ball *ball) getRightTip() int {
	return ball.x + ball.radius
}

func (ball *ball) getLeftTip() int {
	return ball.x - ball.radius
}

func (ball *ball) hitsScreen() screenEdge {
	if ball.y-ball.radius < 0 {
		return top
	}
	if ball.y+ball.radius > winHeight {
		return bottom
	}
	if ball.x-ball.radius < 0 {
		return left
	}
	if ball.x+ball.radius > winWidth {
		return right
	}
	return none
}
