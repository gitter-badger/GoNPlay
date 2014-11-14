package GoNPlay

import (
	glfw "github.com/go-gl/glfw3"
)

// Container for GLFW keystroke information
type KeyStroke struct {
	key    glfw.Key
	action glfw.Action
	mod    glfw.ModifierKey
}

type KeyboardEventHandlers struct {
	events map[KeyStroke]Event
}
