package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWitdh := int32(1600)
	screenHeight := int32(900)

	rl.InitWindow(screenWitdh, screenHeight, "PONG")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	g := newGame()

	fadeOutStartText := 255
	fadeInGame := 0

	for !rl.WindowShouldClose() { // monitorWidth 1920, monitorHeight 1080

		rl.BeginDrawing()

		if fadeOutStartText > 0 {
			fadeOutStartText -= 3
		}

		if fadeOutStartText == 0 {
			if fadeInGame < 255 {
				fadeInGame += 15
			}
			/*if b == nil {
				b = newBall(float32(screenWitdh/2), float32(screenHeight/2), 3, 3)
			} else {
				b.spawn(fadeInGame)
				b.move(float32(screenWitdh), float32(screenHeight))
			}*/
			if fadeInGame == 255 {
				g.update()
			}
			g.start(fadeInGame)
		}

		rl.ClearBackground(rl.Black)
		text := "Welcome to Pong!"
		textSize := int32(30)
		textWidth := int32(rl.MeasureText(text, textSize))

		x := (int32(rl.GetScreenWidth()) - textWidth) / 2
		y := (int32(rl.GetScreenHeight()) - textSize) / 2
		rl.DrawText(text, x, y, textSize, rl.NewColor(255, 255, 255, byte(fadeOutStartText)))

		rl.EndDrawing()
	}
}
