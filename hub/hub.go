package hub

import (
	"bytes"
	"io"
	"reflect"
	"sync"

	"volume/event"
	"volume/log"
)

// Event hub clients
var (
	clientsLock sync.Mutex
	clients     = make([]Client, 0)
)

// Orchestration
var (
	closeC  = make(chan bool)
	closeWg sync.WaitGroup
)

// Recieved event channel
var messageC = make(chan message, 10)

// Initialiser
func init() {
	// Spawn handler routine
	go handler()
}

// Hub client interfaces
type (
	Reader interface {
		Read() ([]byte, error)
	}
	Writer interface {
		Write([]byte) (int, error)
	}
	Client interface {
		Reader
		Writer
	}
)

// Event Handler Interface
type Handler interface {
	Handle(event.Event, Client) error
}

// Handler Func (similar to http.HandlerFunc)
type HandlerFunc func(event.Event, Client) error

// Implements the Handler interface
func (f HandlerFunc) Handle(e event.Event, c Client) error {
	return f(e, c)
}

// Supported event topics mapped to their handlers
var handlers = map[event.Topic]Handler{
	event.VolumeIncreaseTopic: HandlerFunc(increaseVolHandler),
	event.VolumeDecreaseTopic: HandlerFunc(decreaseVolHandler),
}

// Events are stored witht their message payloads and the client
// they came from, this way we can write back just to one specific
// client if required
type message struct {
	Event  event.Event
	Client Client
}

// Consumes events from the event channel and
// calls an appropriate handler function
func handler() {
	for {
		select {
		case msg := <-messageC:
			if handler, ok := handlers[msg.Event.Topic]; ok {
				handler.Handle(msg.Event, msg.Client)
			}
		case <-closeC:
			return
		}
	}
}

// Reads from a Reader and places decoded events onto the
// hub message channel
func read(client Client) {
	defer Remove(client)
	// Read messages from the client and place them on the hub message
	// channel for handling
	for {
		raw, err := client.Read() // Blocking
		if err != nil {
			if err != io.EOF {
				log.WithError(err).Error("unexpected event hub client read error")
			}
			break
		}
		evnt := event.Event{}
		decoder := event.NewDecoder(bytes.NewReader(raw))
		if err := decoder.Decode(&evnt); err != nil {
			log.WithError(err).Error()
		} else {
			messageC <- message{
				Event:  evnt,
				Client: client,
			}
		}
	}
}

// Add a new client to the hub, this will spawn a new reader
// which consumes events from the client
func Add(c Client) {
	// Add to clients list
	clientsLock.Lock()
	clients = append(clients)
	clientsLock.Unlock()
	// Spawn read goroutine
	go read(c)
}

// Removes a client from the clients list
func Remove(client Client) {
	clientsLock.Lock()
	for i := range clients {
		if reflect.DeepEqual(clients[i], client) {
			clients[len(clients)-1], clients[i] = clients[i], clients[len(clients)-1]
			clients = clients[:len(clients)-1]
			break
		}
	}
	clientsLock.Unlock()
}

// Broadcast a message to all connected clients, if a client write error
// occurs the client is removed from the clients list and the error logged
func Broadcast(msg []byte) {
	clientsLock.Lock()
	cl := make([]Client, len(clients))
	copy(cl, clients)
	clientsLock.Unlock()
	for i := range cl {
		client := cl[i]
		if _, err := client.Write(msg); err != nil {
			log.WithError(err).Error("error broadcasting to client")
			defer Remove(client)
		}
	}
}
