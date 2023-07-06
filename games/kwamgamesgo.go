package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func run() error {
	if err := sdl.Init(0x00000020); // Replace 0x00000020 with the appropriate flag value
	err != nil {
		return fmt.Errorf("failed to initialize SDL: %v", err)
	}
	defer sdl.Quit()

	// Create a window
	window, err := sdl.CreateWindow("SDL Window", -1, -1, 800, 600, 0x00000004)
	if err != nil {
		return fmt.Errorf("failed to create window: %v", err)
	}
	defer window.Destroy()

	// Create a renderer
	renderer, err := sdl.CreateRenderer(window, -1, 0x00000002)
	if err != nil {
		return fmt.Errorf("failed to create renderer: %v", err)
	}
	defer renderer.Destroy()

	// Main loop
	running := true
	for running {
		// Event handling
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		// Rendering
		renderer.Clear()
		// Add your rendering code here
		renderer.Present()
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
