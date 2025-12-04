package main

import (
	"testing"
)

func TestPing(t *testing.T) {
	url := "https://reddit.com"
	response, err := TryPing(url)
	if err != nil {
		t.Errorf("Unable to get a ping from %v", url)
	}

	// CRITICAL: You must close the body when you are done!
	defer response.Body.Close()

}
