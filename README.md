# rusrelic [![Go Reference](https://pkg.go.dev/badge/github.com/abrunner94/rusrelic.svg)](https://pkg.go.dev/github.com/abrunner94/rusrelic)

[Logrus](https://github.com/sirupsen/logrus) hook for the [New Relic Log API](https://docs.newrelic.com/docs/logs/log-management/log-api/introduction-log-api/).

## Installation

```bash
go get github.com/mrcrilly/logrus-newrelic-hook
```

## Usage

```golang
// Specify the EU region with "EU" if necessary and your New Relic License Key
rusrelicClient, _ := rusrelic.NewClient("US", "NEW_RELIC_LICENSE_KEY")
rusrelicHook := rusrelic.NewHook(rusrelicClient, rusrelic.DefaultLevels)

// Add the hook to logrus
log.AddHook(rusrelicHook)
```
