package main

import rl "github.com/gen2brain/raylib-go/raylib"

type paddle struct {
	position rl.Vector2
	size     rl.Vector2
	color    rl.Color
	score    int
}

func NewPaddle(posX float32, c rl.Color) *paddle {
	return &paddle{
		position: rl.Vector2{X: posX, Y: float32(rl.GetScreenHeight())/2 - 200/2},
		size:     rl.Vector2{X: 20, Y: 200},
		color:    c,
		score:    0,
	}
}

func (p *paddle) spawn() {
	rl.DrawRectangleV(p.position, p.size, p.color)
}

func (p *paddle) move(up int32, down int32) {
	if rl.IsKeyDown(up) && p.position.Y > 0 {
		(*p).position.Y -= 4
	}

	if rl.IsKeyDown(down) && p.position.Y < float32(rl.GetScreenHeight())-float32(p.size.Y) {
		(*p).position.Y += 4
	}
}
