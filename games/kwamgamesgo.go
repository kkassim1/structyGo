package main

import (
	"fmt"
	"log"

	"time"

	"github.com/veandco/go-sdl2/sdl"

	"math/rand"
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
// 			switch event := event.(type) {
// 			case *sdl.QuitEvent:
// 				running = false
// 			case *sdl.KeyboardEvent:
// 				if event.Keysym.Sym == sdl.K_ESCAPE {
// 					running = false
// 				}
// 				if event.State == sdl.PRESSED {
// 					switch event.Keysym.Sym {
// 					case sdl.K_LEFT:
// 						stickX -= int32(runningSpeed)
// 					case sdl.K_RIGHT:
// 						stickX += int32(runningSpeed)
// 					case sdl.K_UP:
// 						stickY -= int32(runningSpeed)
// 					case sdl.K_DOWN:
// 						stickY += int32(runningSpeed)
// 					}
// 				}
// 			}
// 		}

// 		// Update stick figure position
// 		if stickX > 800 {
// 			stickX = -stickWidth
// 		} else if stickX+stickWidth < 0 {
// 			stickX = 800
// 		}
// 		if stickY > 600 {
// 			stickY = -stickHeight
// 		} else if stickY+stickHeight < 0 {
// 			stickY = 600
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
//---------------------------------------------------------------
//------jj-----------------

const (
	windowWidth  = 800
	windowHeight = 600
	playerSize   = 20
	enemySize    = 20
	bulletSize   = 10
)

type gameObject struct {
	x, y, dx, dy int32
}

func run() error {
	rand.Seed(time.Now().UnixNano())

	if err := sdl.Init(0x00000020); err != nil {
		return fmt.Errorf("failed to initialize SDL: %v", err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("SDL Window", -1, -1, 800, 600, 0x00000004)
	if err != nil {
		return fmt.Errorf("failed to create window: %v", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, 0x00000002)
	if err != nil {
		return fmt.Errorf("failed to create renderer: %v", err)
	}
	defer renderer.Destroy()

	player := gameObject{x: windowWidth / 2, y: windowHeight / 2}
	enemy := gameObject{x: rand.Int31n(windowWidth), y: rand.Int31n(windowHeight)}
	playerBullet := gameObject{dx: -1}
	enemyBullet := gameObject{dx: -1}

	playerHealth := 3
	enemyShootCooldown := 100

	speed := int32(5)
	enemySpeed := int32(3)

	for running := true; running; sdl.Delay(16) {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if event.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				}
				if event.Type == sdl.KEYDOWN {
					switch event.Keysym.Sym {
					case sdl.K_LEFT:
						player.dx = -speed
					case sdl.K_RIGHT:
						player.dx = speed
					case sdl.K_UP:
						player.dy = -speed
					case sdl.K_DOWN:
						player.dy = speed
					case sdl.K_SPACE:
						if playerBullet.dx == -1 {
							playerBullet = gameObject{x: player.x, y: player.y, dx: 0, dy: -speed} // shoot bullet upwards from player's position
						}
					}
				}
				if event.Type == sdl.KEYUP {
					switch event.Keysym.Sym {
					case sdl.K_LEFT, sdl.K_RIGHT:
						player.dx = 0
					case sdl.K_UP, sdl.K_DOWN:
						player.dy = 0
					}
				}
			}
		}

		// enemy AI: move towards player but try to dodge bullets
		if enemy.x < player.x && playerBullet.dx < 0 {
			enemy.dx = enemySpeed
		} else if enemy.x > player.x && playerBullet.dx > 0 {
			enemy.dx = -enemySpeed
		}
		if enemy.y < player.y && playerBullet.dy < 0 {
			enemy.dy = enemySpeed
		} else if enemy.y > player.y && playerBullet.dy > 0 {
			enemy.dy = -enemySpeed
		}

		// enemy AI: shoot at player
		enemyShootCooldown--
		if enemyShootCooldown <= 0 {
			if enemyBullet.dx == -1 {
				enemyBullet = gameObject{x: enemy.x, y: enemy.y, dx: 0, dy: enemySpeed} // shoot bullet downwards from enemy's position
				enemyShootCooldown = 100
			}
		}

		// move player, enemy and bullets
		player.x += player.dx
		player.y += player.dy
		enemy.x += enemy.dx
		enemy.y += enemy.dy
		playerBullet.x += playerBullet.dx
		playerBullet.y += playerBullet.dy
		enemyBullet.x += enemyBullet.dx
		enemyBullet.y += enemyBullet.dy

		// check collision between enemy and player bullet
		if enemy.x < playerBullet.x+bulletSize && enemy.x+enemySize > playerBullet.x &&
			enemy.y < playerBullet.y+bulletSize && enemy.y+enemySize > playerBullet.y {
			// collision detected, respawn enemy and deactivate bullet
			enemy.x = rand.Int31n(windowWidth)
			enemy.y = rand.Int31n(windowHeight)
			playerBullet.dx = -1
		}

		// check collision between player and enemy bullet
		if player.x < enemyBullet.x+bulletSize && player.x+playerSize > enemyBullet.x &&
			player.y < enemyBullet.y+bulletSize && player.y+playerSize > enemyBullet.y {
			// collision detected, decrease player health and deactivate bullet
			playerHealth--
			enemyBullet.dx = -1
			if playerHealth <= 0 {
				running = false // game over
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.FillRect(&sdl.Rect{X: player.x, Y: player.y, W: playerSize, H: playerSize})
		renderer.FillRect(&sdl.Rect{X: enemy.x, Y: enemy.y, W: enemySize, H: enemySize})

		if playerBullet.dx != -1 { // draw bullet only if it's active
			renderer.FillRect(&sdl.Rect{X: playerBullet.x, Y: playerBullet.y, W: bulletSize, H: bulletSize})
		}
		if enemyBullet.dx != -1 { // draw bullet only if it's active
			renderer.FillRect(&sdl.Rect{X: enemyBullet.x, Y: enemyBullet.y, W: bulletSize, H: bulletSize})
		}

		renderer.Present()
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
