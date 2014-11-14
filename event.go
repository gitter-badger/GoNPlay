package GoNPlay

import (
	"fmt"
	glfw "github.com/go-gl/glfw3"
)

type Event interface {
	execute(entity Entity) error
}
