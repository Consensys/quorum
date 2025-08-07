//go:build go1.24

package debug

import "net/http"

// Memsize is a no-op HTTP handler for Go 1.24+ where github.com/fjl/memsize
// no longer builds.
type memsizeStub struct{}

func (memsizeStub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func (memsizeStub) Add(string, interface{}) {}

var Memsize memsizeStub
