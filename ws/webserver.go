package ws

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"
)

// HostServer starts the web server and registers the handler function
func (webaddress WebAddress) HostServer() bool {
	// Register the handler function for the root URL
	http.HandleFunc("/", Handler)

	// Start the web server on port 8080
	log.Printf("Starting server at %v:%v", webaddress.IP, webaddress.PORT)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
	return true
}

// GetIPAddress extracts the IP address from the request
func GetIPAddress(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func (webaddress WebAddress) HostAddress() {
	webaddress.HostServer()
	err := webaddress.OpenInBrowser()
	if err != nil {
		log.Fatalf("Failed to open browser: %v", err)
	}
}

// openBrowser opens the specified URL in the default browser of the user.
func (webaddress WebAddress) OpenInBrowser() error {
	var url string = fmt.Sprintf("%v:%v", webaddress.IP, webaddress.PORT)
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // assume linux or other unix-like
		cmd = "xdg-open"
		args = []string{url}
	}

	return exec.Command(cmd, args...).Start()
}
