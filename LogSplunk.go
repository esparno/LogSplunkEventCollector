package LogSplunkEventCollector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func LogToSplunk(url string, port int, event *Event, token string, timeout time.Duration) (*http.Response, error) {
	logUrl := fmt.Sprintf("%s:%d/services/collector/event", url, port)
	byteSlice, err := json.Marshal(&event)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", logUrl, bytes.NewBuffer(byteSlice))
	if err != nil {
		panic(err)
	}
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Splunk %s", token))
	header.Set("Content-Type", "application/json")
	req.Header = header
	client := http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	return resp, err
}

type Event struct {
	Host         string      `json:"host"`
	SourceType   string      `json:"sourcetype"`
	Index        string      `json:"index"`
	EventMessage interface{} `json:"event"`
}
