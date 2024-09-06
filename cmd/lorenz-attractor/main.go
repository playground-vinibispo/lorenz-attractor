package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var x, y, z float32
	x = 0.01
	sigma := float32(10)
	beta := float32(8) / 3
	ro := float32(28)
	dt := float32(0.01)

	rl.InitWindow(800, 600, "raylib lorenz attractor")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	var points []rl.Vector3
	camera := rl.Camera3D{
		Position:   rl.NewVector3(0, 0, 500),
		Target:     rl.NewVector3(0, 0, 0),
		Up:         rl.NewVector3(0, 1, 0),
		Fovy:       45,
		Projection: rl.CameraPerspective,
	}

	for !rl.WindowShouldClose() {
		// rl.UpdateCamera(&camera, rl.CameraOrbital)
		dx := float32(sigma*(y-x)) * dt
		dy := float32(x*(ro-z)-y) * dt
		dz := float32(x*y-beta*z) * dt

		x += dx
		y += dy
		z += dz
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawText(fmt.Sprintf("x: %f", x), 10, 10, 20, rl.White)
		rl.DrawText(fmt.Sprintf("y: %f", y), 10, 30, 20, rl.White)
		rl.DrawText(fmt.Sprintf("z: %f", z), 10, 50, 20, rl.White)
		rl.DrawText("WASD to move the camera", 10, 70, 20, rl.White)
		rl.BeginMode3D(camera)
		scaleFactor := float32(7)
		// Let's create a keymap to zoom in and out
		if rl.IsKeyDown(rl.KeyW) {
			camera.Position = rl.NewVector3(camera.Position.X, camera.Position.Y, camera.Position.Z+1)
		}
		if rl.IsKeyDown(rl.KeyS) {
			camera.Position = rl.NewVector3(camera.Position.X, camera.Position.Y, camera.Position.Z-1)
		}
		if rl.IsKeyDown(rl.KeyA) {
			camera.Position = rl.NewVector3(camera.Position.X-1, camera.Position.Y, camera.Position.Z)
		}
		if rl.IsKeyDown(rl.KeyD) {
			camera.Position = rl.NewVector3(camera.Position.X+1, camera.Position.Y, camera.Position.Z)
		}
		points = append(points, rl.NewVector3(x, y, z))
		var hue float32
		for i := 0; i < len(points)-1; i++ {
			startPos := rl.Vector3Scale(points[i], scaleFactor)
			endPos := rl.Vector3Scale(points[i+1], scaleFactor)
			rl.DrawLine3D(startPos, endPos, rl.ColorFromHSV(hue, 1, 1))
			hue += 0.1
			if hue > 360 {
				hue = 0
			}
		}
		rl.EndMode3D()

		rl.EndDrawing()
	}
}

