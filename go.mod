module github.com/ozoncp/ocp-poll-api

go 1.15

replace github.com/ozoncp/ocp-poll-api/internal/models => ../modules

require (
	github.com/golang/mock v1.6.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.15.0
)

replace github.com/ozoncp/ocp-poll-api/internal/flusher => ./flusher

replace github.com/ozoncp/ocp-poll-api/internal/repo => ../repo
