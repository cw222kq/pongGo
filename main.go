package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWitdh := int32(1600)
	screenHeight := int32(900)

	rl.InitWindow(screenWitdh, screenHeight, "PONG")
	//rl.SetWindowState(rl.FlagWindowResizable)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	//startTime := time.Now()

	var b *ball

	fadeOutStartText := 255
	fadeInGame := 0

	for !rl.WindowShouldClose() { // monitorWidth 1920, monitorHeight 1080

		//elapsedTime := time.Since(startTime).Seconds()

		rl.BeginDrawing()

		if fadeOutStartText > 0 {
			fadeOutStartText -= 3
		}

		if fadeOutStartText == 0 {
			if fadeInGame < 255 {
				fadeInGame += 15
			}
			if b == nil {
				b = newBall(float32(screenWitdh/2), float32(screenHeight/2), 3, 3)
			} else {
				b.spawn(fadeInGame)
				b.move(float32(screenWitdh), float32(screenHeight))
			}
		}

		rl.ClearBackground(rl.Black)
		text := "Welcome to Pong!"
		textSize := int32(30)
		textWidth := int32(rl.MeasureText(text, textSize))

		x := (screenWitdh - textWidth) / 2
		y := (screenHeight - textSize) / 2
		rl.DrawText(text, x, y, textSize, rl.NewColor(255, 255, 255, byte(fadeOutStartText)))

		rl.EndDrawing()
	}
}
