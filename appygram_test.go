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

func TestSendingAppygram(t *testing.T) {
	client := getClient()
	m := AppygramMessage{
		Name: "Test Gopher", Email: "gopher@gopher.net",
		Phone: "8889991234", Message: "This is a test message",
	}
	status, err := client.sendAppygramMessage(m)
	if err != nil {
		t.Errorf("Error from appygram client %s", err.Error())
	}
	if status.OK != true {
		t.Errorf("Appygram not sent successfully")
	}
}

func TestSendingAppygramTrace(t *testing.T) {
	client := getClient()
	btrace := MakeErrorBackTrace(make([]byte, 1))
	tr := AppygramTrace{Class: "appygram", Backtrace: btrace}
	tr.Name = "Test Gopher"
	tr.Email = "gopher@gopher.net"
	tr.Phone = "8889991234"
	tr.Message = "This is a test message"
	status, err := client.sendAppygramTrace(tr)
	if err != nil {
		t.Errorf("Error from appygram client %s", err.Error())
	}
	if status.OK != true {
		t.Errorf("Appygram trace not sent successfully")
	}
}
