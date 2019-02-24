package main

type ball struct {
	// THIS IS SIMPLE INHERITANCE OF STRUCTS IN GO - NOW YOU CAN ACCESS ball.x INSTEAD OF ball.position.x
	position
	radius int
	xv     float32
	yv     float32
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

func (ball *ball) update(leftPaddle *paddle, rightPaddle *paddle, elapsedTime float32) {

	ball.x += int(ball.xv * elapsedTime)
	ball.y += int(ball.yv * elapsedTime)

	switch ball.hitsScreen() {
	case left:
		rightPaddle.upScore()
		ball.newKickoff()
	case right:
		leftPaddle.upScore()
		ball.newKickoff()
	case top:
		ball.bounceBallY()
	case bottom:
		ball.bounceBallY()
	}

	if leftPaddle.hitByBall(ball) {
		ball.bounceBallX()
	}

	if rightPaddle.hitByBall(ball) {
		ball.bounceBallX()
	}
}

func (ball *ball) bounceBallX() {
	ball.xv = -1.0 * ball.xv
	// HACK TO ESCAPE SEE-SAW COLLISIONS
	ball.x += int(ball.xv) * ball.radius / 2
}

func (ball *ball) bounceBallY() {
	ball.yv = -1.0 * ball.yv
	// HACK TO ESCAPE SEE-SAW COLLISIONS
	ball.y += int(ball.yv) * ball.radius / 2
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
