# lm-logs-sdk-go(beta)
Go SDK for sending logs to Logic Monitor

```go
// Initialize the library
lmIngest := ingest.Ingest{
	HostUrl:   "https://<<account-name>>.logicmonitor.com",
	AccessID:  "<<accesss-id>>",
	AccessKey: "<<access-key>>",
}

// Create logs
logs := []ingest.Log{{
	Message:    "Hello from Logic Monitor!",
	ResourceId: map[string]string{"<<lm-property>>": "<<lm-property-value>>"},
}}

// Send logs to Logic Monitor
ingestResponse, err := lmIngest.SendLogs(logs)
log.Println(ingestResponse) // {"success":true,"message":"Accepted","errors":null,"RequestId":"c952611b-edbf-b670-f94a-9023b38bfdba"}
```
