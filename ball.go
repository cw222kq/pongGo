package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ball struct {
	position rl.Vector2
	speed    rl.Vector2
	radius   float32
}

func NewBall() *ball {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	speedX := float32(3)
	speedY := float32(3)

	if r.Intn(2) == 0 {
		speedX = -speedX
	}

	if r.Intn(2) == 0 {
		speedY = -speedY
	}

	return &ball{
		position: rl.Vector2{X: float32(rl.GetScreenWidth() / 2), Y: float32(rl.GetScreenHeight() / 2)},
		speed:    rl.Vector2{X: speedX, Y: speedY},
		radius:   10.0,
	}
}

func (b *ball) spawn(fadeValue int) {
	rl.DrawCircleV((*b).position, float32((*b).radius), (rl.NewColor(255, 255, 255, byte(fadeValue))))
}

func (b *ball) isCollidingWithPaddle(paddlePos rl.Rectangle) bool {
	return rl.CheckCollisionCircleRec(b.position, b.radius, paddlePos)
}

func (b *ball) isCollidingWithWall() bool {
	minY := b.radius
	maxY := float32(rl.GetScreenHeight()) - b.radius

	if minY > b.position.Y || maxY < b.position.Y {
		return true
	}

	return false
}

func (b *ball) move(paddleLPos rl.Rectangle, paddleRPos rl.Rectangle) { // - left & - up
	b.position.X += b.speed.X
	b.position.Y += b.speed.Y

	if b.isCollidingWithPaddle(paddleLPos) || b.isCollidingWithPaddle(paddleRPos) {
		fmt.Println(b.speed.X)
		b.speed.X = -b.speed.X
		if b.speed.X < 0 {
			b.speed.X = b.speed.X - 1.5
		} else {
			b.speed.X = b.speed.X + 1.5
		}
	}

	if b.isCollidingWithWall() {
		b.speed.Y = -b.speed.Y
	}
}
