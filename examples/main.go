package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	sio "github.com/oarkflow/socketio"
	eio "github.com/oarkflow/socketio/engineio"
	eiot "github.com/oarkflow/socketio/engineio/transport"
	seri "github.com/oarkflow/socketio/serialize"
)

func main() {
	port := ":8001"
	
	server := sio.NewServer(
		eio.WithPingInterval(300*1*time.Millisecond),
		eio.WithPingTimeout(200*1*time.Millisecond),
		eio.WithMaxPayload(1000000),
		eio.WithTransportOption(eiot.WithGovernor(1500*time.Microsecond, 500*time.Microsecond)),
	)
	
	server.OnConnect(func(s *sio.SocketV4) error {
		log.Println("connected:", s.ID())
		s.Of("/").On("notice", CustomWrap(func(a string) error {
			return s.Emit("reply", seri.String("have "+a))
		}))
		s.Of("/").On("bye", CustomWrap(func(a string) error {
			return s.Emit("bye", seri.String(a))
			return nil
		}))
		s.Of("/").On("msg", CustomWrap(func(a string) error {
			return s.Emit("bye", seri.String(a))
			return nil
		}))
		return nil
	})
	server.OnDisconnect(func(reason string) {
		log.Println("closed", reason)
	})
	
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("asset")))
	log.Printf("serving port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Define a custom wrapper 
type CustomWrap func(string) error

// Define your callback
func (cc CustomWrap) Callback(data ...any) error {
	a, aOK := data[0].(string)
	
	if !aOK {
		return fmt.Errorf("bad parameters")
	}
	
	return cc(a)
}
