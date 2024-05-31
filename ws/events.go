package ws

import (
	"net/http"
)

var EventChannel chan<- RequestRecievedEvent

// Events are currently broken. Events are sent, the Handler(), ends before external handling
// Function to trigger events with recursion
func TriggerRequestEvent(eventChan chan<- RequestRecievedEvent, event RequestRecievedEvent, unhandled bool) {
	eventChan <- event
}

func AttachEventChannel(eventChan chan<- RequestRecievedEvent) {
	EventChannel = eventChan
}

// represents a Request Receieved Event
type RequestRecievedEvent struct {
	Handled bool
	Data    interface{}
	Client  Client
	Writer  http.ResponseWriter
}
