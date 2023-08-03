package core

import rl "github.com/gen2brain/raylib-go/raylib"
import "image/color"

var buttonColors map[int32]color.RGBA = map[int32]color.RGBA{
	rl.MouseLeftButton:    rl.Maroon,
	rl.MouseMiddleButton:  rl.Lime,
	rl.MouseRightButton:   rl.DarkBlue,
	rl.MouseSideButton:    rl.Purple,
	rl.MouseExtraButton:   rl.Yellow,
	rl.MouseForwardButton: rl.Orange,
	rl.MouseBackButton:    rl.Beige,
}

const screenWidth = 800
const screenHeight = 450
const framesPerSecond = 60

func MouseInput() {
	rl.InitWindow(screenWidth, screenHeight, "raylib-go [core] example - mouse input")
	defer rl.CloseWindow()
	ballPosition := rl.Vector2{X: 0, Y: 0}
	ballColor := rl.DarkBlue
	rl.SetTargetFPS(framesPerSecond)
	// GAME LOOP
	for !rl.WindowShouldClose() {
		// UPDATE
		ballPosition = rl.GetMousePosition()
		// Moving the button-color mapping into a map may cause color prioritization behavior.
		// else-if is an explicit priority order. a go map is iterated through in a random order.
		for button := range buttonColors {
			if rl.IsMouseButtonPressed(button) {
				ballColor = buttonColors[button]
			}
		}
		// DRAW
		{
			rl.BeginDrawing()
			defer rl.EndDrawing()
			rl.ClearBackground(rl.RayWhite)
			rl.DrawCircleV(ballPosition, 40, ballColor)
			rl.DrawText("move ball with mouse and click mouse button to change color", 10, 10, 20, rl.DarkGray)
		}
	}
}
