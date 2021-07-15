package newrelic

import (
	"runtime"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewClientNoLicenseKey(t *testing.T) {
	_, err := NewClient("US", "")
	assert.EqualError(t, err, "please specify a New Relic License Key")
}

func TestLog(t *testing.T) {
	client, _ := NewClient("US", "abc")
	entry := &logrus.Entry{
		Message: "test message",
		Level:   logrus.ErrorLevel,
		Data: logrus.Fields{
			"field1": "abc",
			"field2": "def",
		},
		Caller: &runtime.Frame{
			Function: "testFn",
			Line:     1243,
		},
	}

	actual, _ := client.Log(entry)
	expected := &event{
		Message: "test message",
		Level:   logrus.ErrorLevel,
		Data: logrus.Fields{
			"field1": "abc",
			"field2": "def",
		},
		Caller: &runtime.Frame{
			Function: "testFn",
			Line:     1243,
		},
	}

	assert.Equal(t, expected, actual)
}
