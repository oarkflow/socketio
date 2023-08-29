package engineio

import (
	"net/http"
	
	eiot "github.com/oarkflow/socketio/engineio/transport"
)

func amp(str string) *string { return &str }

func eioVersionFrom(r *http.Request) EIOVersionStr { return EIOVersionStr(r.URL.Query().Get("EIO")) }
func sessionIDFrom(r *http.Request) SessionID      { return SessionID(r.URL.Query().Get("sid")) }
func transportNameFrom(r *http.Request) eiot.Name  { return eiot.Name(r.URL.Query().Get("transport")) }
