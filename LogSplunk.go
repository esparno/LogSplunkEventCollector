package LogSplunkEventCollector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)
// LogToSplunk will log to Splunk Event Collector
func LogToSplunk(url string, port int, event *Event, token string, timeout time.Duration) (*http.Response, error) {
	logURL := fmt.Sprintf("%s:%d/services/collector/event", url, port)
	byteSlice, err := json.Marshal(&event)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	req, err := http.NewRequest("POST", logURL, bytes.NewBuffer(byteSlice))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	header := http.Header{}
	header.Add("Authorization", fmt.Sprintf("Splunk %s", token))
	header.Set("Content-Type", "application/json")
	req.Header = header
	client := http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	return resp, err
}

// Event represents object sent in http request to Splunk
type Event struct {
	Host         string      `json:"host"`
	Time         int         `json:"time"`
	Source       string      `json:"source"`
	SourceType   string      `json:"sourcetype"`
	Index        string      `json:"index"`
	EventMessage interface{} `json:"event"`
}
