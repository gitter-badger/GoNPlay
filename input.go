package GoNPlay

import (
	glfw "github.com/go-gl/glfw3"
)

type KeyStroke struct {
	key    glfw.Key
	action glfw.Action
	mod    glfw.ModifierKey
}

type KeyboardEventHandlers struct {
	events map[KeyStroke]Event
}
