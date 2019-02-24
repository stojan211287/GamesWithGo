package main

const winHeight int = 800
const winWidth int = 800
const channelDepth int = 4

const noOfPixels int = winHeight * winWidth * channelDepth
const screenPitch int = winWidth * channelDepth

const paddleStartOffset int = 50
const paddleStartY int = winHeight / 2

const paddleScreenBounce int = 10

const ballStartX int = winHeight / 2
const ballStartY int = winWidth / 2

const ballStartXSpeed float32 = 0.5
const ballStartYSpeed float32 = 1.0

const paddleXSize int = 20
const paddleYSize int = 100
const paddleStartSpeed float32 = 1.5 * ballStartYSpeed
const paddleStartScore int = 0

const ballSize int = 20

const scoreYOffset int = 50
const maxScore int = 3

const frameRate int = 200 // IN FRAMES PER SECOND
const frameTime float32 = (1 / 200) * 1000.0

const windowTitle string = "Psychodelic PONG"

// GAME STATES
type gameState int

const (
	play = iota + 1
	paused
)
