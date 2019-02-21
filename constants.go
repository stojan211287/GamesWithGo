package main

const winHeight int = 800
const winWidth int = 800
const channelDepth int = 4

const noOfPixels int = winHeight * winWidth * channelDepth
const screenPitch int = winWidth * channelDepth

const paddleStartOffset int = 50
const paddleStartY int = winHeight / 2

const paddleSpeed int = 5
const paddleScreenBounce int = 10

const ballStartX int = winHeight / 2
const ballStartY int = winWidth / 2

const ballSize int = 20
