package main

import rl "github.com/gen2brain/raylib-go/raylib"

type sound struct {
	start           rl.Sound
	ballOutOfScreen rl.Sound
	gameOver        rl.Sound
	hit             rl.Sound
}

func NewSound() *sound {
	return &sound{
		start:           rl.LoadSound("./assets/start.mp3"),
		ballOutOfScreen: rl.LoadSound("./assets/windSound.wav"),
		gameOver:        rl.LoadSound("./assets/gameOver.mp3"),
		hit:             rl.LoadSound("./assets/click.wav"),
	}
}

func (sound) play(s rl.Sound) {
	rl.PlaySound(s)
}
