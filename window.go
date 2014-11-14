package GoNPlay

import (
	"errors"
	glfw "github.com/go-gl/glfw3"
	"strings"
)

// Configuration for the GLFW Window.
type WindowConfig struct {
	WindowWidth  int    `default:"800" usage:"Sets width of the window"`
	WindowHeight int    `default:"600" usage:"Sets height of the window"`
	WindowName   string `default:"goengine - window" usage:"sets the title of the window"`
}

//
// GLFW3's Context generally refers to the OpenGL context used in the rendering process.
// GLContext is one of the tools
// Allows for configuration of:
// * OpenGL version
// * OpenGL or OpenGLES Profile
//
// #TODO: Explain OpenGL Context better
//
type GLConfig struct {
	GLMajor            int    `default:"3" usage:"Set Major GL Version"`
	GLMinor            int    `default:"3" usage:"Set Minor GL Version"`
	GLAPI              string `default:"OpenGL" usage:"Set Which API. 'OpenGL','OpenGL ES'"`
	GLProfile          string `default:"Core" usage:"Set Profile. 'Any', 'Compact', 'Core'`
	GLFowardCompatible bool   `default:"true" usage:"Should the context be forward compatible.`
}

type GLContext struct {
	config GLConfig
}

//Configures the OpenGL context.
func (glc *GLContext) Setup(config GLConfig) error {
	glc.config = config
	if !glfw.Init() {
		return errors.New("GLFW failed to Init")
	}

	var profile int

	if val, ok := glProfile[strings.ToUpper(glc.config.GLProfile)]; !ok {
		profile = glfw.OpenglCoreProfile
		logger.Error("glProfile flag was not correctly set: Default to Core")
	} else {
		profile = val
	}

	var api int

	if val, ok := glProfile[strings.ToUpper(glc.config.GLProfile)]; !ok {
		api = glfw.OpenglCoreProfile
		logger.Error("glProfile flag was not correctly set: Default to OpenGL")
	} else {
		api = val
	}

	logger.Info("Creating GLFW Context")
	glfw.WindowHint(glfw.ContextVersionMajor, glc.config.GLMajor)
	glfw.WindowHint(glfw.ContextVersionMinor, glc.config.GLMinor)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glBool[glc.config.GLFowardCompatible])
	glfw.WindowHint(glfw.OpenglProfile, profile)
	glfw.WindowHint(glfw.ClientApi, api)

	return nil
}

//Destroys the OpenGL Context
func (glc *GLContext) Terminate() {
	glfw.Terminate()
}

// Creates the Initial glfw Window context.
// Returns ( error, glfw3.window )
//
//	glfw3 Window Documentation: http://godoc.org/github.com/go-gl/glfw3#Window
func SetupWindow(config WindowConfig) (error, *glfw.Window) {

	logger.Info("Generating GLFW Window")
	window, err := glfw.CreateWindow(config.WindowWidth, config.WindowHeight, config.WindowName, nil, nil)
	if err != nil {
		return errors.New("Failed to create a window"), nil
	}

	return nil, window
}
