package main

import (
  rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
  rl.InitWindow(400, 400, "lorenz-attractor")
  rl.SetTargetFPS(60)
  for !rl.WindowShouldClose() {
    rl.BeginDrawing()
    rl.ClearBackground(rl.RayWhite)
    rl.DrawText("Hello, world!", 12, 12, 20, rl.Maroon)
    rl.EndDrawing()
  }
  rl.CloseWindow()
}
    