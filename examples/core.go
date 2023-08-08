package examples

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidth = 1920
const screenHeight = 1080
const framesPerSecond = 120

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

func BunnyMark() {
	const MaxBunnies = 500000
	const MaxBatchElements = 8192
	type Bunny struct {
		position rl.Vector2
		speed    rl.Vector2
		color    rl.Color
		texture  *rl.Texture2D
	}
	rl.InitWindow(screenWidth, screenHeight, "raylib example - bunnymark")
	defer rl.CloseWindow()
	paths := []string{"assets/roach_american.png", "assets/roach_missoula.png", "assets/roach_oriental.png"}
	var texBunnies []rl.Texture2D
	for _, path := range paths {
		texBunnies = append(texBunnies, rl.LoadTexture(path))
	}

	bunnies := make([]Bunny, MaxBunnies)
	bunniesCount := 0
	rl.SetTargetFPS(framesPerSecond)
	for !rl.WindowShouldClose() {
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			spawnedPerFrame := 100
			for i := 0; i < spawnedPerFrame; i++ {
				if bunniesCount < MaxBunnies {
					bunnies[bunniesCount].position = rl.GetMousePosition()
					bunnies[bunniesCount].speed.X = float32(rl.GetRandomValue(-250, 250) / 60)
					bunnies[bunniesCount].speed.Y = float32(rl.GetRandomValue(-250, 250) / 60)
					bunnies[bunniesCount].color = rl.NewColor(
						uint8(rl.GetRandomValue(50, 240)),
						uint8(rl.GetRandomValue(80, 240)),
						uint8(rl.GetRandomValue(100, 240)), 255)
					bunnies[bunniesCount].texture = &texBunnies[i%3]
					bunniesCount += 1
				}
			}
		}
		// Update Bunnies
		for i := 0; i < bunniesCount; i++ {
			bunnies[i].position = rl.Vector2Add(bunnies[i].position, bunnies[i].speed) // move bunny position
			texBunny := texBunnies[0]                                                  // assume tex's have the same dimensions for now.
			// when position exceeds boundary, change direction of speed for a bounce effect
			if ((bunnies[i].position.X + float32(texBunny.Width/2)) > float32(rl.GetScreenWidth())) ||
				((bunnies[i].position.X + float32(texBunny.Width/2)) < 0) {
				bunnies[i].speed.X *= -1
			}
			if ((bunnies[i].position.X + float32(texBunny.Width/2)) > float32(rl.GetScreenWidth())) ||
				((bunnies[i].position.X + float32(texBunny.Width/2)) < 0) {
				bunnies[i].speed.Y *= -1
			}
		}
		// Draw
		func() {
			rl.BeginDrawing()
			defer rl.EndDrawing()
			rl.ClearBackground(rl.RayWhite)
			for i := 0; i < bunniesCount; i++ {
				rl.DrawTexture(*bunnies[i].texture,
					int32(bunnies[i].position.X), int32(bunnies[i].position.Y), bunnies[i].color)
			}

			rl.DrawRectangle(0, 0, screenWidth, 40, rl.Black)
			rl.DrawText(fmt.Sprintf("bunnies: %v", bunniesCount), 120, 10, 20, rl.Green)
			rl.DrawText(fmt.Sprintf("batched draw calls: %v", 1+bunniesCount/MaxBatchElements), 320, 10, 20, rl.Maroon)
			rl.DrawFPS(10, 10)
		}()
	}
	for _, tex := range texBunnies {
		rl.UnloadTexture(tex)
	}
}

func FPSCamera() {
	const MaxColumns = 20
	rl.InitWindow(screenWidth, screenHeight, "raylib example - FPS Camera")
	Camera := rl.NewCamera3D
}
