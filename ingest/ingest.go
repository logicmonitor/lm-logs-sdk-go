package ingest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"../apitoken"
)

type Log struct {
	Message    string            `json:"msg"`
	Timestamp  time.Time         `json:"timestamp"`
	ResourceID map[string]string `json:"_lm.resourceId"`
}

type Ingest struct {
	CompanyName string
	AccessID    string
	AccessKey   string
}

func (in *Ingest) SendLogs(logs []Log) (*Response, error) {
	url := fmt.Sprintf("https://%s.logicmonitor.com/rest/log/ingest", in.CompanyName)

	body, err := json.Marshal(logs)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	lMv1Token := apitoken.GenerateLMv1Token(in.AccessID, in.AccessKey, body)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", lMv1Token.String())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	ingestResponse, err := convertHTTPToIngestResponse(resp)
	if err != nil {
		return ingestResponse, err
	}

	return ingestResponse, nil
}
