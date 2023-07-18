package main

import (
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	winWidth  = 1280
	winHeight = 720
)

func main() {

	err := sdl.Init(0xFFFF)
	if err != nil {
		panic(err)
	}

	defer sdl.Quit()

	window, err := sdl.CreateWindow("hello triangle", 200, 200, winWidth, winHeight, 0x00000002) // Open a window.
	if err != nil {
		panic(err)
	}

	window.GLCreateContext()
	defer window.Destroy()
	gl.Init()
	ver := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("version ,", ver)
	vertices := []float32{
		.5, .5, .0,
		.5, -.5, 0.0,
		.0, .5, .0,
	}

	var VBO uint32

	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

}
