package ingest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	resourcePath = "/log/ingest"
	method       = "POST"
)

type Log struct {
	Message    string            `json:"msg"`
	Timestamp  time.Time         `json:"timestamp"`
	ResourceId map[string]string `json:"_lm.resourceId"`
}

type Ingest struct {
	HostUrl   string
	AccessID  string
	AccessKey string
}

func (in *Ingest) SendLogs(logs []Log) (*Response, error) {
	url := in.HostUrl + "/rest/log/ingest"

	body, err := json.Marshal(logs)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	lMv1Token := generateLMv1Token(in.AccessID, in.AccessKey, body)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", lMv1Token.String())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	ingestResponse, err := convertHttpToIngestResponse(resp)
	if err != nil {
		return ingestResponse, err
	}

	return ingestResponse, nil
}
