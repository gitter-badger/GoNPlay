package GoNPlay

import (
	"errors"
	glfw "github.com/go-gl/glfw3"
)

/*
ContextConfig is a configuration struct meant to be used with:
	github.com/GreenShift/GoNPlay/window
*/
type ContextConfig struct {
	GLMajor int `default:"3" usage:"Set Major GL Version"`
	GLMinor int `default:"3" usage:"Set Minor GL Version"`
}

type WindowConfig struct {
	// Essential Setup
	// Window and GL Setup
	WindowWidth  int    `default:"800" usage:"Sets width of the window"`
	WindowHeight int    `default:"600" usage:"Sets height of the window"`
	WindowName   string `default:"goengine - window" usage:"sets the title of the window"`
}

func SetupContext(config ContextConfig) error {

	if !glfw.Init() {
		return errors.New("GLFW failed to Init")
	}

	logger.Info("Creating GLFW Context")
	glfw.WindowHint(glfw.ContextVersionMajor, config.GLMajor)
	glfw.WindowHint(glfw.ContextVersionMinor, config.GLMinor)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	return nil
}

func SetupWindow(config WindowConfig) (error, *glfw.Window) {

	logger.Info("Generating GLFW Window")
	window, err := glfw.CreateWindow(config.WindowWidth, config.WindowHeight, config.WindowName, nil, nil)
	if err != nil {
		return errors.New("Failed to create a window"), nil
	}

	return nil, window
}
