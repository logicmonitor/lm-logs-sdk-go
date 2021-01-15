# lm-logs-sdk-go(beta)
Go SDK for sending logs to LogicMonitor

**NOTE:** This SDK was created for use by LogicMonitor-built log integrations and is not intended to be used or supported otherwise.

```go
// Initialize the library
lmIngest := ingest.Ingest{
	CompanyName: "<<account-name>>",
	AccessID:  "<<accesss-id>>",
	AccessKey: "<<access-key>>",
}

// Create logs
logs := []ingest.Log{{
    Message:    "Hello from LogicMonitor!",
    ResourceID: map[string]string{"<<lm-property>>": "<<lm-property-value>>"},
}}


// Send logs to Logic Monitor
ingestResponse, err := lmIngest.SendLogs(logs)
log.Println(ingestResponse) // {"success":true,"message":"Accepted","errors":null,"RequestId":"c952611b-edbf-b670-f94a-9023b38bfdba"}
```
