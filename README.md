# Logrus New Relic Hook

[Logrus](https://github.com/sirupsen/logrus) hook for the [New Relic Log API](https://docs.newrelic.com/docs/logs/log-management/log-api/introduction-log-api/).

## Installation

```bash
go get github.com/mrcrilly/logrus-newrelic-hook
```

## Usage

```golang
// Specify the EU region with "EU" if necessary and your New Relic License Key
newrelicHookClient, _ := logrusnewrelic.NewClient("US", "NEW_RELIC_LICENSE_KEY")
newrelicHook := logrusnewrelic.NewHook(newrelicHookClient, logrusnewrelic.DefaultLevels)

// Add the hook to logrus
log.AddHook(newrelicHook)
```

## Why fork?

I didn't want to rely on an external repository that could disappear at any time. And I didn't like the `rusreclic`.
