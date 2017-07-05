# LogSplunkEventCollector
Library containing simple function for logging to Splunk Event Collector

## Example Code
```
event := LogSplunkEventCollector.Event{
		Host: "myhost",
		SourceType: "mysourcetype",
		Index: "myindex",
		EventMessage: "test event message",
	}
	resp, err := LogSplunkEventCollector.LogToSplunk("https://urlforlogging.com", 443, &event, "TOKEN", time.Minute)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(r))
```
