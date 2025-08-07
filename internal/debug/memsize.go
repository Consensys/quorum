//go:build !go1.24

package debug

import "github.com/fjl/memsize/memsizeui"

// Memsize exposes memory statistics at /memsize.
var Memsize memsizeui.Handler
