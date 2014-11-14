// Golang Enviornment file sets up the global enviornment that
// GoNPlay used throughout the project.

package GoNPlay

import (
	"github.com/spacemonkeygo/spacelog"
)

var (
	// Logger Tools I will use across GoNPlay. This is based
	// on Space Monkey's "spacelog" Go library. This can be
	// found at github.com/spacemonkeygo/spacelog
	logger = spacelog.GetLogger()
)
