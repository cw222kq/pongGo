package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type game struct {
	ball        *ball
	paddleL     *paddle
	paddleR     *paddle
	startTime   time.Time
	elapsedTime float64
}

func newGame() *game {
	return &game{
		ball:      NewBall(),
		paddleL:   NewPaddle(float32(80.0)),
		paddleR:   NewPaddle(float32(rl.GetScreenWidth() - 80.0)),
		startTime: time.Now(),
	}
}

func (g *game) start(fadeInGame int) {
	g.ball.spawn(fadeInGame)
	g.paddleL.spawn()
	g.paddleR.spawn()
}

func (g *game) isBallOutOfScreen() bool {
	if g.ball.position.X < 0 || g.ball.position.X > float32(rl.GetScreenWidth()) {
		g.ball = nil
		return true
	}

	return false
}

func (g *game) update() {
	g.elapsedTime = time.Since(g.startTime).Seconds()

	//fmt.Println(g.elapsedTime)
	if g.isBallOutOfScreen() {
		g.ball = NewBall()
	}

	g.ball.move(rl.Rectangle{X: g.paddleL.position.X, Y: g.paddleL.position.Y, Width: g.paddleL.size.X, Height: g.paddleL.size.Y}, rl.Rectangle{X: g.paddleR.position.X, Y: g.paddleR.position.Y, Width: g.paddleR.size.X, Height: g.paddleR.size.Y})
	g.paddleL.move(rl.KeyW, rl.KeyS)
	g.paddleR.move(rl.KeyUp, rl.KeyDown)

}
