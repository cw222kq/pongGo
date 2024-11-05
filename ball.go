package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ball struct {
	position rl.Vector2
	speed    rl.Vector2
	radius   float32
}

func NewBall(posX float32, posY float32, speedX float32, speedY float32) *ball {
	return &ball{
		position: rl.Vector2{X: posX, Y: posY},
		speed:    rl.Vector2{X: speedX, Y: speedY},
		radius:   10.0,
	}
}

func (b *ball) spawn(fadeValue int) {
	rl.DrawCircleV((*b).position, float32((*b).radius), (rl.NewColor(255, 255, 255, byte(fadeValue))))
}

func (b *ball) move(screenWitdh float32, screenHeight float32) { // - left & - up
	b.position.X += b.speed.X
	b.position.Y += b.speed.Y

	minX := b.radius
	minY := b.radius
	maxX := float32(screenWitdh) - b.radius
	maxY := float32(screenHeight) - b.radius

	if minX > b.position.X || maxX < b.position.X {
		b.speed.X = -b.speed.X
	}

	if minY > b.position.Y || maxY < b.position.Y {
		b.speed.Y = -b.speed.Y
	}
}
