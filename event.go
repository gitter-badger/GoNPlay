package GoNPlay

import (
	glfw "github.com/go-gl/glfw3"
)

func SetupEventHandler() {
	logger.Info("Configuring EventHandler")
}

type KeyboardEventHandler struct {
	events map[glfw.Key]Event
}

type EventHandler struct {
}

type Event interface {
	execute(entity *Entity) error
}
