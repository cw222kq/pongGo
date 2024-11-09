package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type game struct {
	ball      *ball
	paddleL   *paddle
	paddleR   *paddle
	sound     *sound
	scoreSize int32
}

func NewGame() *game {
	g := &game{
		ball:      NewBall(),
		paddleL:   NewPaddle(float32(80.0), rl.NewColor(0, 102, 204, 255)),
		paddleR:   NewPaddle(float32(rl.GetScreenWidth()-80.0-20.0), rl.NewColor(255, 204, 0, 255)),
		sound:     NewSound(),
		scoreSize: 60,
	}

	g.sound.play(g.sound.start)

	return g
}

func (g *game) updateScore() {
	if g.ball.position.X < 0 {
		(*g.paddleR).score += 1
	}

	if g.ball.position.X > float32(rl.GetScreenWidth()) {
		(*g.paddleL).score += 1
	}

}

func (g *game) drawScores(fadeInGame int, textSize int32) {
	leftScoreText := strconv.Itoa(g.paddleL.score)
	rightScoreText := strconv.Itoa(g.paddleR.score)
	dashText := " - "

	totalTextWidth := rl.MeasureText(leftScoreText, textSize) + rl.MeasureText(dashText, textSize) + rl.MeasureText(rightScoreText, textSize)
	screenWidth := float32(rl.GetScreenWidth())
	startX := (screenWidth - float32(totalTextWidth)) / 2
	yPos := float32(rl.GetScreenHeight()) * 0.1
	if !g.isGameOngoing() {
		yPos *= 2
		endText := "Press 'N' to start a new game or 'Q' to quit."
		endTextSize := int32(60)
		endTextWidth := rl.MeasureText(endText, endTextSize)
		endTextstartX := (screenWidth - float32(endTextWidth)) / 2
		rl.DrawText(endText, int32(endTextstartX), int32(yPos+600), endTextSize, rl.NewColor(255, 255, 255, byte(fadeInGame)))

	}

	rl.DrawText(leftScoreText, int32(startX), int32(yPos), textSize, rl.NewColor(0, 102, 204, byte(fadeInGame)))

	dashX := startX + float32(rl.MeasureText(leftScoreText, textSize))
	rl.DrawText(dashText, int32(dashX), int32(yPos), textSize, rl.NewColor(255, 255, 255, byte(fadeInGame)))

	rightScoreX := dashX + float32(rl.MeasureText(dashText, textSize))
	rl.DrawText(rightScoreText, int32(rightScoreX), int32(yPos), textSize, rl.NewColor(255, 204, 0, byte(fadeInGame)))

}

func (g *game) isGameOngoing() bool {
	if g.paddleL.score < 5 && g.paddleR.score < 5 {
		return true
	}
	return false
}

func (g *game) isBallOutOfScreen() bool {
	if g.ball.position.X < 0 || g.ball.position.X > float32(rl.GetScreenWidth()) {
		g.sound.play(g.sound.ballOutOfScreen)
		g.updateScore()
		if g.isGameOngoing() {
			g.ball = nil
		} else {
			g.sound.play(g.sound.gameOver)
		}
		return true
	}

	return false
}

func (g *game) start(fadeInGame int) {
	if g.isGameOngoing() {
		g.ball.spawn(fadeInGame)
		g.paddleL.spawn()
		g.paddleR.spawn()
		g.drawScores(fadeInGame, 60)
	}
}

func (g *game) update() {

	if g.isGameOngoing() {
		if g.isBallOutOfScreen() {
			g.ball = NewBall()
		}

		if g.ball.isChangingDir {
			g.sound.play(g.sound.hit)
			g.ball.isChangingDir = false
		}

		g.ball.move(rl.Rectangle{X: g.paddleL.position.X, Y: g.paddleL.position.Y, Width: g.paddleL.size.X, Height: g.paddleL.size.Y}, rl.Rectangle{X: g.paddleR.position.X, Y: g.paddleR.position.Y, Width: g.paddleR.size.X, Height: g.paddleR.size.Y})
		g.paddleL.move(rl.KeyW, rl.KeyS)
		g.paddleR.move(rl.KeyUp, rl.KeyDown)
		g.drawScores(255, g.scoreSize)
		return
	}

	if g.scoreSize < 600 {
		g.drawScores(255, g.scoreSize)
		g.scoreSize += 6
	}

	g.drawScores(255, g.scoreSize)

}
