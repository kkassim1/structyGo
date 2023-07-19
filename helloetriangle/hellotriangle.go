package main

import (
	"fmt"
	"strings"

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

	vertexshadersrc :=
		`#version 330 core
	layout (location = 0) in vec3 aPos;

	void main(){
		gl_Position = vec4(aPos.x,aPos.y,aPos.z,1.0);
	}`

	vertexshader := gl.CreateShader(gl.VERTEX_SHADER)
	csource, free := gl.Strs(vertexshadersrc + "\x00")
	gl.ShaderSource(vertexshader, 1, csource, nil)
	free()
	gl.CompileShader(vertexshader)

	var status int32

	gl.GetShaderiv(vertexshader, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var loglen int32
		gl.GetShaderiv(vertexshader, gl.INFO_LOG_LENGTH, &loglen)
		log := strings.Repeat("\n", int(loglen+1))

		gl.GetShaderInfoLog(vertexshader, loglen, nil, gl.Str(log))
		panic("failed to compile vertex shader : \n" + log)
	}

	fragmentshadersrc :=
		`#version 330 core
 out vec4 FragColor;
 
 void main()
 {
	FragColor = vec4(1.0f,.5f,.2f,1.0f);
 }`
	fragmentshader := gl.CreateShader(gl.FRAGMENT_SHADER)
	csource, free = gl.Strs(fragmentshadersrc + "\x00")
	gl.ShaderSource(fragmentshader, 1, csource, nil)
	free()
	gl.CompileShader(fragmentshader)

	gl.GetShaderiv(fragmentshader, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var loglen int32
		gl.GetShaderiv(fragmentshader, gl.INFO_LOG_LENGTH, &loglen)
		log := strings.Repeat("\n", int(loglen+1))

		gl.GetShaderInfoLog(fragmentshader, loglen, nil, gl.Str(log))
		panic("failed to compile fragments shader : \n" + log)
	}

	// uint 32 shaderprogra

}
