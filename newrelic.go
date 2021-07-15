package newrelic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const requestTimeout = time.Second * 15

type event struct {
	Message  string       `json:"message"`
	LogLevel logrus.Level `json:"logLevel"`
}

type Client struct {
	Region     string
	LicenseKey string
}

func NewClient(region string, licenseKey string) *Client {
	return &Client{Region: region, LicenseKey: licenseKey}
}

func (c *Client) Log(entry *logrus.Entry) error {
	logEvent := &event{
		Message:  entry.Message,
		LogLevel: entry.Level,
	}

	for k, v := range entry.Data {
		entry.Data[k] = v
	}

	json, err := json.Marshal(logEvent)
	if err != nil {
		return err
	}

	return c.request(json)
}

func (c *Client) request(json []byte) error {
	// Determine URL based on region provided
	var url = "https://log-api.newrelic.com/log/v1"

	if c.Region == "EU" {
		url = "https://log-api.eu.newrelic.com/log/v1"
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
