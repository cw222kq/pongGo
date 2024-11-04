package main

import (
	"fmt"
	"time"
)

type game struct {
	ball        *ball
	witdh       float32
	height      float32
	startTime   time.Time
	elapsedTime float64
}

func newGame(screenWitdh float32, screenHeight float32) *game {
	return &game{
		ball:      newBall(float32(screenWitdh/2), float32(screenHeight/2), 3, 3),
		witdh:     screenWitdh,
		height:    screenHeight,
		startTime: time.Now(),
	}

}

func (g *game) start(fadeInGame int) {
	g.ball.spawn(fadeInGame)
}

func (g *game) update() {
	g.elapsedTime = time.Since(g.startTime).Seconds()

	fmt.Println(g.elapsedTime)

	g.ball.move(g.witdh, g.height)

}
