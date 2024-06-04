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

	// Algorithm for Handling http requests.
	// 1) First we will create a client struct from the request information.

	// Create the client
	client := Client{
		IPAddress:  GetIPAddress(r),
		UserAgent:  r.UserAgent(),
		Referer:    r.Referer(),
		RequestURI: r.RequestURI,
	}

	// 2) Next we will handle methods

	// Handling methods, we will first handle post method.
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)

		//Ensures body entegrity
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusInternalServerError)
			return
		}

		//This variable will hold the data from body
		var data map[string]string

		// Unmarshals the JSON into a slice variable. We use pointers here.
		err = json.Unmarshal(body, &data)

		//Verify json integrity and cause JsonBasedRequestEevnt
		if err == nil {

			//Create the request recieved event
			requestevent := JsonBasedRequestEvent{
				Handled: false,
				Client:  client,
				Data:    data,
				Writer:  w,
			}

			// Send json request event through channel
			jsonbasedrequesteventchannel <- &requestevent

			//Increment waitgroup in advance
			eventswait.Add(1)

			//Wait for handler to decrement
			eventswait.Wait()

			log.Print("Sync Wait finished for JsonBasedRequestEevnt..")

			if requestevent.PreventDefault {
				return
			}

			// Log the entity information
			log.Printf("Client: %v triggered JsonBasedRequestEvent", client.IPAddress)

			return
		}

		log.Printf("Couldn't cast: %v to map[string]string", body)
	}

	// We will now handle get method for static and abstract file paths
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
			if osfiles.FileExists("index.html") {
				requestedPath = "/index.html"
				log.Printf("Client: %v was served existing index.html", client.IPAddress)
			} else {
				//Send invalid path even through channel
				invalidpatheventchannel <- &invalidpathevent

				//Increment waitgroup in advance
				eventswait.Add(1)

				//Wait for handler to decrement
				eventswait.Wait()
			}

			log.Print("Sync Wait finished..")
		}

		// Construct the file path
		filePath := filepath.Join(".", requestedPath)

		//Otherwise, we're either handling a legitimate url path or an abstract one

		//Define variables fileContent for response
		var fileContent []byte

		// Define the content type to respond client
		var contentType string

		//Check if file exists and handle serving of existing files
		if osfiles.FileExists(filePath) {
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

			//Define error variable for reading file
			var err error

			//Populate file contents
			fileContent, err = os.ReadFile(filePath)

			if err != nil {
				log.Printf("Something went wrong reading from file: %v ", filePath)
			}
		}

		// Check if the file exists
		if !osfiles.FileExists(filePath) {

			//Log filepart not found so causing InvalidPathEvent
			log.Printf("Client: %v caused InvalidPathEvent for: %v", client.IPAddress, invalidpathevent.Url)

			//Send InvalidPathEvent through channel
			invalidpatheventchannel <- &invalidpathevent

			//Increment waitgroup in advance
			eventswait.Add(1)

			//Wait for handler to decrement
			eventswait.Wait()

			//Log sync wait finish
			log.Print("Sync Wait finished..")

			//Check if PreventDefault was ignored and requesting url is 404
			if !invalidpathevent.PreventDefault && !osfiles.FileExists(filePath) {
				http.NotFound(w, r)
				return
			}

			//Check if PreventDefault was set and no optional data was given
			if invalidpathevent.PreventDefault && invalidpathevent.Optional == nil {
				http.NotFound(w, r)
				return
			}

			//Check if prevent default was set and content type was given
			if invalidpathevent.PreventDefault && invalidpathevent.ContentType != "" {
				contentType = invalidpathevent.ContentType

				if invalidpathevent.Optional != nil {
					fileContent = invalidpathevent.Optional

					//Log served client with abstract path
					log.Printf("Client: %v was served abstract file path for: %v", client.IPAddress, invalidpathevent.Url)
				}
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
