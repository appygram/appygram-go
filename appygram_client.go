package appygram

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const base = "https://arecibo.appygram.com"

type AppygramClient struct {
	ApiKey     string
	HttpClient *http.Client
}

func (ac *AppygramClient) buildRequest(method string, urlStr string, reader io.Reader) (r *http.Request, err error) {
	r, err = http.NewRequest(method, urlStr, reader)
	r.Header.Add("Accept", "application/json")
	return
}

func (ac *AppygramClient) GetTopics() (at AppygramTopics, err error) {
	url := base + "/topics/" + ac.ApiKey
	request, err := ac.buildRequest("GET", url, nil)
	response, err := ac.HttpClient.Do(request)
	if err != nil {
		at.OK = false
		return
	}
  defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		at.OK = false
		return
	} else {
		at.OK = true
	}
	var tops []AppygramTopic
	json.Unmarshal(body, &tops)
	at.Topics = tops
	return
}

func (ac *AppygramClient) SendAppygramMessage(m AppygramMessage) (ar AppygramResult, err error) {
	url := base + "/appygrams"
	withKey := AppygramMessageWithKey{AppygramMessage: m, ApiKey: ac.ApiKey}
	bytes, err := json.Marshal(withKey)
	if err != nil {
		ar.OK = false
		return
	}
	request, err := ac.buildRequest("POST", url, strings.NewReader(string(bytes)))
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		ar.OK = false
		return
	}
	response, err := ac.HttpClient.Do(request)
	if err != nil {
		ar.OK = false
		return
	}
  defer response.Body.Close()
	if response.StatusCode == 200 {
		ar.OK = true
	}
	return
}

func (ac *AppygramClient) SendAppygramTrace(s AppygramTraceWithMessage) (ar AppygramResult, err error) {
	url := base + "/traces"
	withKey := AppygramTraceWithKey{AppygramTraceWithMessage: s, ApiKey: ac.ApiKey}
	bytes, err := json.Marshal(withKey)
	if err != nil {
		ar.OK = false
		return
	}
	request, err := ac.buildRequest("POST", url, strings.NewReader(string(bytes)))
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		ar.OK = false
		return
	}
	response, err := ac.HttpClient.Do(request)
	if err != nil {
		ar.OK = false
		return
	}
  defer response.Body.Close()
	if response.StatusCode == 200 {
		ar.OK = true
	}
	return
}
