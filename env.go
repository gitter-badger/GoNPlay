// Golang Enviornment file sets up the global enviornment that
// GoNPlay used throughout the project.

package GoNPlay

import (
	glfw "github.com/go-gl/glfw3"
	"github.com/spacemonkeygo/spacelog"
)

// Logger Tools I will use across GoNPlay. This is based
// on Space Monkey's "spacelog" Go library. This can be
// found at github.com/spacemonkeygo/spacelog
var (
	logger = spacelog.GetLogger()
)

// Generic Conversion from golang `type` to glfw.type
// We've created these mappings so we can make
// decisions based on context
var (
	glBool = map[bool]int{true: glfw.True, false: glfw.False}
	glAPI  = map[string]int{
		"OPENGL":    glfw.OpenglApi,
		"OPENGL ES": glfw.OpenglEsApi,
	}
	glProfile = map[string]int{
		"ANY":     glfw.OpenglAnyProfile,
		"COMPACT": glfw.OpenglCompatProfile,
		"CORE":    glfw.OpenglCoreProfile,
	}
)
