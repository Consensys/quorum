// +build !go1.7

package xhandler

import (
        "net/http"

        "golang.org/x/net/context"
)


// New creates a conventional http.Handler injecting the provided root
// context to sub handlers. This handler is used as a bridge between conventional
// http.Handler and context aware handlers.
func New(ctx context.Context, h HandlerC) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                h.ServeHTTPC(ctx, w, r)
        })
}
