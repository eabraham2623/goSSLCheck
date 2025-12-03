package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func validate_args() string {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ping.go <HTTP_URL>")
		log.Fatalf("URL not given, please give HTTP URL")
	}
	url := os.Args[1]
	return url
}

func try_ping(url string) (*http.Response, error) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("🔴 PING FAILED for %s: %v\n", url, err)
		return nil, err
	}

	return response, nil
}

func check_ping_response(response *http.Response, err error, url string) {
	defer response.Body.Close()
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

func ping_error_check(err error) {
	if err != nil {
		log.Fatalf("")
	}
}

func main() {
	url := validate_args()
	response, err := try_ping(url)
	ping_error_check(err)
	check_ping_response(response, err, url)
}
