package ws

import (
	"net/http"
	"sync"
)

var eventswait sync.WaitGroup

var jsonbasedrequesteventchannel chan<- *JsonBasedRequestEvent

// Struct representing request receieved event.
type JsonBasedRequestEvent struct {
	Data           map[string]string
	Client         Client
	Writer         http.ResponseWriter
	Handled        bool
	PreventDefault bool
}

var invalidpatheventchannel chan<- *InvalidPathEvent

// Struct representing request receieved event.
type InvalidPathEvent struct {
	Url            string
	Parameters     string
	Client         Client
	Writer         http.ResponseWriter
	Handled        bool
	Optional       []byte
	ContentType    string
	PreventDefault bool
}

// This functions attaches an event channel to the web server. Use pointers.
func AttachJsonBasedRequestEventChannel(eventChan chan<- *JsonBasedRequestEvent) *sync.WaitGroup {
	jsonbasedrequesteventchannel = eventChan
	return &eventswait
}

// This functions attaches an event channel to the web server. Use pointers.
func AttachInvalidPathEventChannel(eventChan chan<- *InvalidPathEvent) *sync.WaitGroup {
	invalidpatheventchannel = eventChan
	return &eventswait
}
