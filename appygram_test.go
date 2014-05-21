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
	return Client(getApiKey(), &http.Client{})
}

func TestGetTopics(t *testing.T) {
	client := getClient()
	topics, err := client.GetTopics()
	if err != nil {
		t.Errorf("Topics top level error %s", err.Error())
		t.FailNow()
	}
	if len(topics.Topics) == 0 {
		t.Errorf("Topics should have length > 0")
		t.FailNow()
	}
	if topics.Topics[0].Name != "Feedback" {
		t.Errorf("Topics[0].Name should be Feedback, not %s", topics.Topics[0])
	}
}

func TestSendingAppygram(t *testing.T) {
	client := getClient()
	m := AppygramMessage{
		Name: "Test Gopher", Email: "gopher@gopher.net",
		Phone: "8889991234", Message: "This is a test message",
		Topic: "Feedback",
	}
	status, err := client.SendAppygramMessage(m)
	if err != nil {
		t.Errorf("Error from appygram client %s", err.Error())
	}
	if status.OK != true {
		t.Errorf("Appygram not sent successfully")
	}
}

func TestSendingAppygramTrace(t *testing.T) {
	client := getClient()
	//btrace := MakeErrorBackTrace(make([]byte, 1))
	tr := AppygramTrace{Class: "appygram", Message: "Gophers"} //, Backtrace: btrace}
	m := AppygramMessage{
		Name: "Test Gopher", Email: "gopher@gopher.net",
		Phone: "8889991234", Message: "This is a test message",
		Topic: "Exception",
	}
	trm := AppygramTraceWithMessage{AppygramMessage: m, Trace: tr}
	status, err := client.SendAppygramTrace(trm)
	if err != nil {
		t.Errorf("Error from appygram client %s", err.Error())
	}
	if status.OK != true {
		t.Errorf("Appygram trace not sent successfully")
	}
}
