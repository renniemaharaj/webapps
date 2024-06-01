package ws

// Client represents a client entity connecting to the web server.
type Client struct {
	IPAddress  string
	UserAgent  string
	Referer    string
	RequestURI string
}
