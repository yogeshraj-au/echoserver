package main

import (
	"fmt"
	"io/ioutil" // Import ioutil package
	"log"
	"net/http"
	"os"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Log the hostname
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Failed to get hostname: %v", err)
		hostname = "Unknown"
	}
	log.Printf("Hostname: %s", hostname)

	// Log pod information (currently unavailable)
	log.Println("Pod Information:")
	log.Println("\t-no pod information available-")

	// Log server values
	log.Println("Server values:")
	log.Printf("\tserver_version=%s\n", r.Header.Get("Server"))

	// Log request information
	log.Println("Request Information:")
	log.Printf("\tclient_address=%s\n", r.RemoteAddr)
	log.Printf("\tmethod=%s\n", r.Method)
	log.Printf("\treal_path=%s\n", r.URL.Path)
	log.Printf("\trequest_scheme=%s\n", r.URL.Scheme)
	log.Printf("\trequest_uri=%s\n", r.RequestURI)

	// Log request headers
	log.Println("Request Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			log.Printf("\t%s=%s\n", name, value)
		}
	}

	// Log request body
	log.Println("Request Body:")
	// Read request body
	body := r.Body
	// Check if request body is empty
	if r.ContentLength == 0 {
		log.Println("\t-no body in request-")
	} else {
		// Log the request body if present
		bodyContent, err := ioutil.ReadAll(body)
		if err != nil {
			log.Printf("\tError reading request body: %v", err)
		} else {
			log.Printf("\t%s\n", bodyContent)
		}
	}
	
	// Echo back the request path
	fmt.Fprintf(w, "Echoserver: %s\n", r.URL.Path)

	// Log the response
	log.Printf("Sent response: %d", http.StatusOK)
}

// This function will handle root path ("/") requests and display the metrics.
// This function will handle root path ("/") requests and display the metrics.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Log the hostname
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Failed to get hostname: %v", err)
		hostname = "Unknown"
	}

	// Prepare the response content
	response := fmt.Sprintf(`Hostname: %s
Pod Information:
    -no pod information available-
Server values:
    server_version=%s
Request Information:
    client_address=%s
    method=%s
    real_path=%s
    request_scheme=%s
    request_uri=%s
Request Headers:
`, hostname, r.Header.Get("Server"), r.RemoteAddr, r.Method, r.URL.Path, r.URL.Scheme, r.RequestURI)

	// Write request headers to the response
	for name, values := range r.Header {
		for _, value := range values {
			response += fmt.Sprintf("    %s=%s\n", name, value)
		}
	}

	// Write response to the client browser
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	// Define HTTP handlers
	http.HandleFunc("/", rootHandler) // Register rootHandler for the root path ("/")
	http.HandleFunc("/echo", echoHandler)

	// Start HTTP server
	port := ":8080"
	log.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server error: %v", err)
		os.Exit(1)
	}
}
