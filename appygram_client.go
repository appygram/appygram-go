package appygram

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const base = "https://arecibo.appygram.com"

type AppygramClient struct {
	ApiKey     string
	HttpClient http.Client
}

func (ac *AppygramClient) buildRequest(method string, urlStr string) (r *http.Request, err error) {
	r, err = http.NewRequest(method, urlStr, nil)
	r.Header.Add("Accept", "application/json")
	return
}

func (ac *AppygramClient) getTopics() (at AppygramTopics, err error) {
	url := base + "/topics/" + ac.ApiKey
	request, err := ac.buildRequest("GET", url)
	response, err := ac.HttpClient.Do(request)
	if err != nil {
		at.OK = false
		return
	}
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

func (ac *AppygramClient) sendAppygramMessage(m AppygramMessage) (AppygramResult, error) {
	return AppygramResult{false}, nil
}

func (ac *AppygramClient) sendAppygramTrace(s AppygramTrace) (AppygramResult, error) {
	return AppygramResult{false}, nil
}
