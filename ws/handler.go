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

		//This variable will hold the data from body
		var data map[string]interface{}

		// Unmarshals the JSON into a slice variable. We use pointers here.
		err = json.Unmarshal(body, &data)

		//Verify json integrity and respond with bad request.
		if err == nil {

			//Create the request recieved event
			requestevent := JsonBasedRequestEvent{
				Handled: false,
				Client:  client,
				Data:    data,
				Writer:  w,
			}

			// Send json request event through channel
			jsonbasedrequesteventchannel <- requestevent

			if requestevent.PreventDefault {
				return
			}

			// Log the entity information
			log.Printf("\nClient: %+v, triggered InvalidPathEvent sending request: %v", client.IPAddress, requestevent.Data)
		}

		if err != nil {
			log.Printf("\nGot bad json from client:%v", client.IPAddress)
			return
		}

		return
	}

	// Handle only GET requests for static files
	if r.Method == http.MethodGet {
		// Extract the requested file path from the URL
		requestedPath := r.URL.Path

		//Create the invalidpathevent
		invalidpathevent := InvalidPathEvent{
			Url:            requestedPath,
			Optional:       nil,
			Handled:        false,
			Client:         client,
			Writer:         w,
			PreventDefault: false,
		}

		// If requestPath is '/' check for index.html or php.html.
		if requestedPath == "/" {
			if osfiles.FileExists("/index.html") {
				requestedPath = "/index.html"
			} else {
				invalidpatheventchannel <- &invalidpathevent
			}

		}

		// Construct the file path
		filePath := filepath.Join(".", requestedPath)

		// Check if the file exists
		if !osfiles.FileExists(filePath) {

			//Send invalid path even through channel
			invalidpatheventchannel <- &invalidpathevent

			//Increment waitgroup in advance
			eventswait.Add(1)

			//Wait for handler to decrement
			eventswait.Wait()
		}

		log.Print("Sync Wait finished..")

		//Check if PreventDefault was set to true and return else continue
		if invalidpathevent.PreventDefault {
			if invalidpathevent.Optional == nil {
				// Log the entity information
				log.Printf("Client: %v caused InvalidPathEvent for: %v", client.IPAddress, invalidpathevent.Url)
			}
		} else {
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

		if invalidpathevent.ContentType != "" {
			contentType = invalidpathevent.ContentType
		}
		// Read the content of the file
		fileContent, err := os.ReadFile(filePath)

		if err != nil {
			if invalidpathevent.Optional != nil {
				fileContent = invalidpathevent.Optional
			}
		}

		// Set the content type
		w.Header().Set("Content-Type", contentType)

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
