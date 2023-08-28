package ingest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/logicmonitor/lm-logs-sdk-go/apitoken"
)

// Log will contain information about logs being sent to logingest
type Log struct {
	Message    string                 `json:"msg"`
	Timestamp  time.Time              `json:"timestamp"`
	ResourceID map[string]string      `json:"_lm.resourceId"`
	Metadata   map[string]interface{} `json:"metadata"`
}

// Ingest will contain details about endpoint authorisation and code version details
type Ingest struct {
	CompanyName    string
	AccessID       string
	AccessKey      string
	LogSource      string
	VersionID      string
	BearerToken    string
	useBearerToken bool
}

// create new ingest
func NewLogIngester(company string, accessId string, accessKey string, bearerToken string, logsource string, version string) (*Ingest, error) {
	useBeareFrorAuth := false
	if len(accessId) == 0 || len(accessKey) == 0 {
		useBeareFrorAuth = true
	}
	if useBeareFrorAuth == true && len(bearerToken) == 0 {
		// No valid mode of auth soecified
		return nil, errors.New("Can not set up authentication with LogicMonitor. Either specify accessId and accessKey both or bearerToken")
	}

	return &Ingest{
		CompanyName:    company,
		AccessID:       accessId,
		AccessKey:      accessKey,
		LogSource:      logsource,
		VersionID:      version,
		useBearerToken: useBeareFrorAuth,
		BearerToken:    bearerToken,
	}, nil

}

// SendLogs will be used to send logs to logingest
func (in *Ingest) SendLogs(logs []Log) (*Response, error) {
	url := fmt.Sprintf("https://%s.logicmonitor.com/rest/log/ingest", in.CompanyName)

	var logsMapArray []map[string]interface{}

	for i := 0; i < len(logs); i++ {
		logMap := make(map[string]interface{})
		eachLog := logs[i]
		logMap["msg"] = eachLog.Message
		logMap["_lm.resourceId"] = eachLog.ResourceID
		logMap["timestamp"] = eachLog.Timestamp.String()
		for key, value := range eachLog.Metadata {
			logMap[key] = value
		}
		logsMapArray = append(logsMapArray, logMap)
	}

	body, err := json.Marshal(logsMapArray)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	auth := in.GenerateAuthString(body)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)
	req.Header.Set("User-Agent", in.LogSource+"/"+in.VersionID)

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

func (in *Ingest) GenerateAuthString(body []byte) string {

	if in.useBearerToken == false {
		return apitoken.GenerateLMv1Token(in.AccessID, in.AccessKey, body).String()

	} else {
		return fmt.Sprintf("Bearer %s", in.BearerToken)
	}
}
