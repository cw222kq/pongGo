package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWitdh := int32(1280)
	screenHeight := int32(720)

	rl.InitWindow(screenWitdh, screenHeight, "raylib [core] example - basic window")
	//rl.SetWindowState(rl.FlagWindowResizable)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		text := "Congrats! You created your first window!"
		textSize := int32(20)
		textWidth := int32(rl.MeasureText(text, textSize))

		x := (screenWitdh - textWidth) / 2
		y := screenHeight / 2

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText(text, x, y, textSize, rl.LightGray)

		rl.EndDrawing()
	}
}
