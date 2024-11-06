package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type game struct {
	ball        *ball
	paddleL     *paddle
	paddleR     *paddle
	witdh       float32
	height      float32
	startTime   time.Time
	elapsedTime float64
}

func newGame(screenWitdh float32, screenHeight float32) *game {
	return &game{
		ball:      NewBall(float32(screenWitdh/2), float32(screenHeight/2), -3, 3),
		paddleL:   NewPaddle(float32(screenHeight), 80.0),
		paddleR:   NewPaddle(float32(screenHeight), float32(screenWitdh-80.0)),
		witdh:     screenWitdh,
		height:    screenHeight,
		startTime: time.Now(),
	}

}

func (g *game) start(fadeInGame int) {
	g.ball.spawn(fadeInGame)
	g.paddleL.spawn()
	g.paddleR.spawn()
}

func (g *game) update() {
	g.elapsedTime = time.Since(g.startTime).Seconds()

	//fmt.Println(g.elapsedTime)

	g.ball.move(rl.Rectangle{X: g.paddleL.position.X, Y: g.paddleL.position.Y, Width: g.paddleL.size.X, Height: g.paddleL.size.Y}, rl.Rectangle{X: g.paddleR.position.X, Y: g.paddleR.position.Y, Width: g.paddleR.size.X, Height: g.paddleR.size.Y})
	g.paddleL.move(rl.KeyW, rl.KeyS)
	g.paddleR.move(rl.KeyUp, rl.KeyDown)

}
