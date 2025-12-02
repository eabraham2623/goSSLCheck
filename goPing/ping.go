package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// 1. Check for command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run simple_ping.go <URL>")
		os.Exit(1)
	}
	url := os.Args[1]

	// 2. Perform a simple GET request
	response, err := http.Get(url)

	// 3. Handle Errors
	if err != nil {
		// This handles network errors, DNS failures, etc.
		fmt.Printf("🔴 PING FAILED for %s: %v\n", url, err)
		return
	}

	// 4. Close the response body (crucial!)
	defer response.Body.Close()

	// 5. Report Status
	status := response.StatusCode

	if status >= 200 && status < 300 {
		fmt.Printf("🟢 PING SUCCESS! %s responded with status %d (%s)\n",
			url, status, http.StatusText(status))
	} else {
		// This handles application errors like 404 Not Found or 500 Server Error
		fmt.Printf("🟡 PING RECEIVED, but FAILED with status %d (%s)\n",
			status, http.StatusText(status))
	}
}
