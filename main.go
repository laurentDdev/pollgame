package main

import (
	"billarggame/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	rl.InitWindow(1280, 720, "Billiard Game")
	defer rl.CloseWindow()

	camera := rl.Camera3D{
		Position: rl.Vector3{
			X: 10,
			Y: 50,
			Z: 10,
		},
		Target:     rl.Vector3{X: 0, Y: 0, Z: 0},
		Up:         rl.Vector3{X: 0, Y: 1, Z: 0},
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}

	g := game.NewGame()

	g.Init()

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera, rl.CameraFree)

		//var deltaTime = rl.GetFrameTime()

		camera.Target = rl.Vector3{0, 0, 0}

		//Camera.Target = posCar

		//posCar.Z += 1 * deltaTime
		//rotCar.Y += 1 * deltaTime
		//car.Transform = rl.MatrixRotateXYZ(rotCar)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode3D(camera)
		//rl.DrawModel(car, posCar, 1.0, rl.Blue)
		rl.DrawGrid(20, 10)

		g.Update()

		rl.EndMode3D()
		rl.EndDrawing()
	}

	g.Unload()

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
