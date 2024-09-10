package main

import (
	"metaballs-demo/src/state"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint | rl.FlagWindowResizable) // Optionally add `rl.FlagFullscreenMode`
	rl.InitWindow(512, 512, "Metaballs Demo")
	// rl.SetWindowMonitor(0) // For development only
	rl.SetTargetFPS(60)

	metaballs := state.NewMetaballs(rl.GetScreenWidth(), rl.GetScreenHeight())

	// Game loop
	for !rl.WindowShouldClose() {
		if rl.IsWindowResized() {
			metaballs.OnResize(rl.GetScreenWidth(), rl.GetScreenHeight())
		}

		metaballs.Update()

		rl.BeginDrawing()
		metaballs.Draw()
		rl.EndDrawing()
	}

	metaballs.Unload()
	rl.CloseWindow()
}
