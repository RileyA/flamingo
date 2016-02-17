package gfx

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

// This is the root for all graphics-related state. The scene graph is rooted
// here, the GL context and Window are created/managed here, etc.
type Context struct {
	window *glfw.Window
}

func CreateContext(width int, height int) (*Context, error) {
	if err := glfw.Init(); err != nil {
		fmt.Println("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(width, height, "flamingo", nil, nil)

	if err != nil {
		glfw.Terminate()
		return nil, err
	}

	window.MakeContextCurrent()

	// Initialize Glow
	if err = gl.Init(); err != nil {
		glfw.Terminate()
		return nil, err
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	return &Context{window}, nil
}

func (context *Context) SwapBuffers() {
	context.window.SwapBuffers()
}

func (context *Context) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (context *Context) SetBackgroundColor(color *Colorf) {
	gl.ClearColor(color[0], color[1], color[2], color[3])
}

func (context *Context) Release() {
	glfw.Terminate()
}
