package ws

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"newrennie/osfiles"
	"os"
	"path/filepath"
	"strings"
)

// Handler handles incoming HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	// Create the client
	client := Client{
		IPAddress:  GetIPAddress(r),
		UserAgent:  r.UserAgent(),
		Referer:    r.Referer(),
		RequestURI: r.RequestURI,
	}

	// Handle POST requests
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)

		//Ensures body entegrity
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusInternalServerError)
			return
		}

		var data map[string]interface{}

		err = json.Unmarshal(body, &data) // Unmarshals the JSON into a slice variable
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		//Create the request recieved event
		requestevent := RequestRecievedEvent{
			Handled: false,
			Client:  client,
			Data:    data,
			Writer:  w,
		}

		// Log the entity information
		log.Printf("\nTriggering client: %+v sending request: %v", client, requestevent)

		// TriggerRequestEvent()
		TriggerRequestEvent(EventChannel, requestevent, false)

		return
	}

	// Handle only GET requests for static files
	if r.Method == http.MethodGet {
		// Extract the requested file path from the URL
		requestedPath := r.URL.Path

		// If requestPath is '/' check for index.html or php.html and return
		if requestedPath == "/" {
			requestedPath = "/index.html"
		}

		// Construct the file path
		filePath := filepath.Join(".", requestedPath)

		// Check if the file exists
		if !osfiles.FileExists(filePath) {
			http.NotFound(w, r)
			return
		}

		// Determine the content type based on the file extension
		var contentType string
		switch strings.ToLower(filepath.Ext(filePath)) {
		case ".html":
			contentType = "text/html"
		case ".css":
			contentType = "text/css"
		case ".js":
			contentType = "application/javascript"
		case ".png":
			contentType = "image/png"
		case ".jpg", ".jpeg":
			contentType = "image/jpeg"
		case ".gif":
			contentType = "image/gif"
		default:
			contentType = "text/plain"
		}

		// Set the content type
		w.Header().Set("Content-Type", contentType)

		// Read the content of the file
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Could not read %s: %s\n", filePath, err.Error())
		}

		// Write the file content to the response
		fmt.Fprint(w, string(fileContent))

		// Optionally, include the query parameters in the response
		queryParams := r.URL.Query()
		if len(queryParams) > 0 {
			fmt.Fprint(w, "\n\nQuery Parameters:\n")
			for key, values := range queryParams {
				for _, value := range values {
					fmt.Fprintf(w, "%s: %s\n", key, value)
				}
			}
		}
		return
	}

	// Respond with 405 Method Not Allowed for other methods
	w.Header().Set("Allow", http.MethodGet)
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
