package rusrelic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

const requestTimeout = time.Second * 15

var (
	EndpointGeneric = "https://log-api.newrelic.com/log/v1"
	EndpointEU      = "https://log-api.eu.newrelic.com/log/v1"
)

type Client struct {
	Region     string
	LicenseKey string
}

type event struct {
	Message string         `json:"Message"`
	Level   logrus.Level   `json:"Level"`
	Caller  *runtime.Frame `json:"Caller"`
	Data    logrus.Fields  `json:"Data"`
}

func NewClient(region string, licenseKey string) (*Client, error) {
	if licenseKey == "" {
		return nil, fmt.Errorf("please specify a New Relic License Key")
	}
	return &Client{Region: region, LicenseKey: licenseKey}, nil
}

func (c *Client) Log(entry *logrus.Entry) (*event, error) {
	evt := &event{
		Message: entry.Message,
		Level:   entry.Level,
		Caller:  entry.Caller,
		Data:    entry.Data,
	}
	json, err := json.Marshal(evt)
	if err != nil {
		return nil, err
	}

	return evt, c.request(json)
}

func (c *Client) request(json []byte) error {
	// Determine URL based on region provided
	var url = EndpointGeneric
	if c.Region == "EU" {
		url = EndpointEU
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return fmt.Errorf("could not make a request to the New Relic Log API: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF-8")
	req.Header.Set("X-License-Key", c.LicenseKey)

	client := &http.Client{Timeout: requestTimeout}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return err
}
