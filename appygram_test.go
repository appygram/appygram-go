package appygram

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func getApiKey() string {
	key := os.Getenv("API_KEY")
	fmt.Println(key)
	if key == "" {
		panic("Please set API_KEY in the shell environment.")
	}
	return key
}

func getClient() AppygramClient {
	return Client(getApiKey(), http.Client{})
}

func TestGetTopics(t *testing.T) {
	client := getClient()
	topics := client.getTopics()
	if len(topics) == 0 {
		t.Errorf("Topics should have length > 0")
	}
}
