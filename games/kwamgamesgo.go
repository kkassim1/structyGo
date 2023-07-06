package main

import (
	"fmt"
	"log"

	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// func run() error {
// 	if err := sdl.Init(0x00000020); // Replace 0x00000020 with the appropriate flag value
// 	err != nil {
// 		return fmt.Errorf("failed to initialize SDL: %v", err)
// 	}
// 	defer sdl.Quit()

// 	// Create a window
// 	window, err := sdl.CreateWindow("SDL Window", -1, -1, 800, 600, 0x00000004)
// 	if err != nil {
// 		return fmt.Errorf("failed to create window: %v", err)
// 	}
// 	defer window.Destroy()

// 	// Create a renderer
// 	renderer, err := sdl.CreateRenderer(window, -1, 0x00000002)
// 	if err != nil {
// 		return fmt.Errorf("failed to create renderer: %v", err)
// 	}
// 	defer renderer.Destroy()

// 	// Main loop
// 	running := true
// 	for running {
// 		// Event handling
// 		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
// 			switch event.(type) {
// 			case *sdl.QuitEvent:
// 				running = false
// 				break
// 			}
// 		}

// 		// Rendering
// 		renderer.Clear()
// 		// Add your rendering code here
// 		renderer.Present()
// 	}

// 	return nil
// }

// func main() {
// 	if err := run(); err != nil {
// 		log.Fatal(err)
// 	}
// }

//--------------------------------------------------------------
//-------------------------------------------------------------

// func run() error {
// 	if err := sdl.Init(0x00000020); err != nil {
// 		return fmt.Errorf("failed to initialize SDL: %v", err)
// 	}
// 	defer sdl.Quit()

// 	// Create a window
// 	window, err := sdl.CreateWindow("SDL Window", -1, -1, 800, 600, 0x00000004)
// 	if err != nil {
// 		return fmt.Errorf("failed to create window: %v", err)
// 	}
// 	defer window.Destroy()

// 	// Create a renderer
// 	renderer, err := sdl.CreateRenderer(window, -1, 0x00000002)
// 	if err != nil {
// 		return fmt.Errorf("failed to create renderer: %v", err)
// 	}
// 	defer renderer.Destroy()

// 	// Stick figure variables
// 	stickX := int32(400)           // Initial x-coordinate of the stick figure
// 	stickY := int32(300)           // Initial y-coordinate of the stick figure
// 	stickWidth := int32(20)        // Width of the stick figure
// 	stickHeight := int32(60)       // Height of the stick figure
// 	runningSpeed := 5              // Speed at which the stick figure moves
// 	delay := time.Millisecond * 16 // Delay between each iteration (adjust as needed for desired speed)

// 	// Main loop
// 	running := true
// 	for running {
// 		// Event handling
// 		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
// 			switch event.(type) {
// 			case *sdl.QuitEvent:
// 				running = false
// 				break
// 			}
// 		}

// 		// Update stick figure position
// 		stickX += int32(runningSpeed)
// 		if stickX > 800 {
// 			stickX = -stickWidth
// 		}

// 		// Rendering
// 		renderer.SetDrawColor(0, 0, 0, 255) // Set the draw color to black
// 		renderer.Clear()

// 		// Draw stick figure
// 		renderer.SetDrawColor(255, 255, 255, 255) // Set draw color to white
// 		renderer.DrawRect(&sdl.Rect{X: stickX, Y: stickY, W: stickWidth, H: stickHeight})

// 		renderer.Present()

// 		time.Sleep(delay) // Introduce a delay between each iteration
// 	}

// 	return nil
// }

// func main() {
// 	if err := run(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// ---------erfef----------------------

func run() error {
	if err := sdl.Init(0x00000020); err != nil {
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

	// Stick figure variables
	stickX := int32(400)           // Initial x-coordinate of the stick figure
	stickY := int32(300)           // Initial y-coordinate of the stick figure
	stickWidth := int32(20)        // Width of the stick figure
	stickHeight := int32(60)       // Height of the stick figure
	runningSpeed := 5              // Speed at which the stick figure moves
	delay := time.Millisecond * 16 // Delay between each iteration (adjust as needed for desired speed)

	// Main loop
	running := true
	for running {
		// Event handling
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if event.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				}
				if event.State == sdl.PRESSED {
					switch event.Keysym.Sym {
					case sdl.K_LEFT:
						stickX -= int32(runningSpeed)
					case sdl.K_RIGHT:
						stickX += int32(runningSpeed)
					case sdl.K_UP:
						stickY -= int32(runningSpeed)
					case sdl.K_DOWN:
						stickY += int32(runningSpeed)
					}
				}
			}
		}

		// Update stick figure position
		if stickX > 800 {
			stickX = -stickWidth
		} else if stickX+stickWidth < 0 {
			stickX = 800
		}
		if stickY > 600 {
			stickY = -stickHeight
		} else if stickY+stickHeight < 0 {
			stickY = 600
		}

		// Rendering
		renderer.SetDrawColor(0, 0, 0, 255) // Set the draw color to black
		renderer.Clear()

		// Draw stick figure
		renderer.SetDrawColor(255, 255, 255, 255) // Set draw color to white
		renderer.DrawRect(&sdl.Rect{X: stickX, Y: stickY, W: stickWidth, H: stickHeight})

		renderer.Present()

		time.Sleep(delay) // Introduce a delay between each iteration
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
