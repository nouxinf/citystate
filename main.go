package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - 3d mode")
	defer rl.CloseWindow()

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(500.0, 500.0, 500.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 1.0
	camera.Projection = rl.CameraPerspective

	//cubePosition := rl.NewVector3(0.0, 0.0, 0.0)
	cityMap := [][]string{{"s","s","s", "s"}, {"s", "s", "s"}, {"s", "s", "s"}}
	

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)
		rl.DrawGrid(10, 1.0)
		//rl.DrawCube(cubePosition, 2.0, 2.0, 2.0, rl.Red)
		for z, row := range cityMap {
			for x, cell := range row {
				if cell == "s" {
					rl.DrawCube(rl.NewVector3(float32(x * 2.0), 0.0, float32(z * 2.0)), 2.0, 2.0, 2.0, rl.Red)
				}
			}
		}
		
		rl.EndMode3D()
		rl.DrawText("Welcome to the third dimension!", 10, 40, 20, rl.DarkGray)
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}