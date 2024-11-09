package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWitdh := int32(1600)
	screenHeight := int32(900)

	rl.InitWindow(screenWitdh, screenHeight, "PONG")
	rl.InitAudioDevice()
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	g := NewGame()

	fadeOutStartText := 255
	fadeInGame := 0

	for !rl.WindowShouldClose() { // monitorWidth 1920, monitorHeight 1080

		rl.BeginDrawing()

		if !g.isGameOngoing() {
			if rl.IsKeyDown(rl.KeyN) {
				g = NewGame()
			}

			if rl.IsKeyDown(rl.KeyQ) {
				return
			}
		}

		if fadeOutStartText > 0 {
			fadeOutStartText -= 3
		}

		if fadeOutStartText == 0 {
			if fadeInGame < 255 {
				fadeInGame += 15
			}

			if fadeInGame == 255 {
				g.update()
			}
			g.start(fadeInGame)
		}

		rl.ClearBackground(rl.Black)
		text := "Welcome to Pong!"
		textSize := int32(40)
		textWidth := int32(rl.MeasureText(text, textSize))

		startX := (int32(rl.GetScreenWidth()) - textWidth) / 2
		yPos := (int32(rl.GetScreenHeight()) - textSize) / 2
		rl.DrawText(text, startX, yPos, textSize, rl.NewColor(255, 255, 255, byte(fadeOutStartText)))

		rl.EndDrawing()
	}
}
