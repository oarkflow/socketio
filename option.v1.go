package socketio

import (
	"strings"
	
	eio "github.com/oarkflow/socketio/engineio"
)

// WithPath changes the path when using the SocketIO engine in
// conjunction with EngineIO. Use the engineio.WithPath to change
// the path when only using EngineIO.
func WithPath(path string) Option {
	return func(o OptionWith) {
		if v, ok := o.(*ServerV1); ok {
			if path == "" {
				return
			}
			path = "/" + strings.Trim(path, "/") + "/"
			v.path = &path
			v.eio.With(eio.WithPath(path))
		}
	}
}
