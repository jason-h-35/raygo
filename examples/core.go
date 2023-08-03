package core

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidth = 800
const screenHeight = 450
const framesPerSecond = 60

func MouseInput() {
	var buttonColors map[int32]color.RGBA = map[int32]color.RGBA{
		rl.MouseLeftButton:    rl.Maroon,
		rl.MouseMiddleButton:  rl.Lime,
		rl.MouseRightButton:   rl.DarkBlue,
		rl.MouseSideButton:    rl.Purple,
		rl.MouseExtraButton:   rl.Yellow,
		rl.MouseForwardButton: rl.Orange,
		rl.MouseBackButton:    rl.Beige,
	}
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

func Camera3DMode() {
	rl.InitWindow(screenWidth, screenHeight, "raylib-go [core] example - 3D camera mode")
	defer rl.CloseWindow()
	position := rl.NewVector3(0, 10, 10)
	target := rl.NewVector3(0, 0, 0)
	up := rl.NewVector3(0, 1, 0)
	var fovy float32 = 45.0
	proj := rl.CameraPerspective
	camera := rl.NewCamera3D(position, target, up, fovy, proj)
	cubePosition := rl.NewVector3(0, 0, 0)
	rl.SetTargetFPS(framesPerSecond)

	for !rl.WindowShouldClose() {
		func() {
			rl.BeginDrawing()
			defer rl.EndDrawing()
			rl.ClearBackground(rl.RayWhite)
			func() {
				rl.BeginMode3D(camera)
				defer rl.EndMode3D()
				rl.DrawCube(cubePosition, 2, 2, 2, rl.Red)
				rl.DrawCubeWires(cubePosition, 2, 2, 2, rl.Maroon)
				rl.DrawGrid(10, 10)
			}()
			rl.DrawText("Welcome to the third dimension!", 10, 40, 20, rl.DarkGray)
			rl.DrawFPS(10, 10)
		}()
	}
}
