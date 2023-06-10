package httpd

import "strings"

// HealthCheck is a special callback used by functions that would provide a healthcheck
// of some kind, such as whether the IRC client is up or whether the database is opened.
//
// Callbacks will be called by the healthcheck engine whenever the status of the service
// is requested. They should return a boolean that returns true unless the service is not
// OK, and a string that may further clarify the current status of the service.
type HealthCheck func() (bool, string)

// healthcheck is the internal data structure used to represent a healthcheck list.
type healthcheck struct {
	checks map[string]HealthCheck
}

// newHealthcheck creates a healthcheck engine.
func newHealthcheck() healthcheck {
	var (
		checks = make(map[string]HealthCheck)
		engine = healthcheck{checks: checks}
	)
	return engine
}

// AddCheck registers a new healthcheck callback for the service with the given name.
func (hc *healthcheck) AddCheck(name string, callback HealthCheck) {
	hc.checks[name] = callback
}

// GetResults executes all the healthchecks and returns the outcome of the checks,
// both as a string with a pretty representation that can be used in the HTTP
// endpoint, and as a number with the amount of failing checks.
func (hc *healthcheck) GetResults() (string, int64) {
	var (
		builder       = strings.Builder{}
		errors  int64 = 0
	)
	for name, check := range hc.checks {
		builder.WriteString(name)
		success, outcome := check()
		if success {
			builder.WriteString(" OK ")
		} else {
			builder.WriteString(" ERROR ")
			errors++
		}
		if len(outcome) > 0 {
			builder.WriteString(": " + outcome)
		}
		builder.WriteString("\n")
	}
	return builder.String(), errors
}
