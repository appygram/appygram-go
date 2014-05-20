package appygram

import (
	"net/http"
	"os"
	"testing"
)

func getApiKey() string {
	key := os.Getenv("API_KEY")
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
	if topics[0] != "Feedback" {
		t.Errorf("Topics[0] should be Feedback, not %s", topics[0])
	}
}
