package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
)

func main() {
	// register the handler for all requests
	http.HandleFunc("/", handleRequest)

	// get port from environment variable, default to 80
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	// start the server
	fmt.Printf("Server starting on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

// handleRequest displays all request information including headers
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// set content type to plain text for easy reading
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	var sb strings.Builder

	// basic request info
	sb.WriteString("=== REQUEST INFO ===\n\n")
	sb.WriteString(fmt.Sprintf("Method:         %s\n", r.Method))
	sb.WriteString(fmt.Sprintf("URL:            %s\n", r.URL.String()))
	sb.WriteString(fmt.Sprintf("Path:           %s\n", r.URL.Path))
	sb.WriteString(fmt.Sprintf("Raw Query:      %s\n", r.URL.RawQuery))
	sb.WriteString(fmt.Sprintf("Protocol:       %s\n", r.Proto))
	sb.WriteString(fmt.Sprintf("Host:           %s\n", r.Host))
	sb.WriteString(fmt.Sprintf("Remote Address: %s\n", r.RemoteAddr))
	sb.WriteString(fmt.Sprintf("Request URI:    %s\n", r.RequestURI))

	// headers section - sorted alphabetically for consistent output
	sb.WriteString("\n=== HEADERS ===\n\n")
	headerNames := make([]string, 0, len(r.Header))
	for name := range r.Header {
		headerNames = append(headerNames, name)
	}
	sort.Strings(headerNames)

	for _, name := range headerNames {
		// headers can have multiple values
		values := r.Header[name]
		sb.WriteString(fmt.Sprintf("%s: %s\n", name, strings.Join(values, ", ")))
	}

	// query parameters if any
	if len(r.URL.Query()) > 0 {
		sb.WriteString("\n=== QUERY PARAMETERS ===\n\n")
		queryKeys := make([]string, 0, len(r.URL.Query()))
		for key := range r.URL.Query() {
			queryKeys = append(queryKeys, key)
		}
		sort.Strings(queryKeys)

		for _, key := range queryKeys {
			values := r.URL.Query()[key]
			sb.WriteString(fmt.Sprintf("%s: %s\n", key, strings.Join(values, ", ")))
		}
	}

	// write the response
	fmt.Fprint(w, sb.String())
}
