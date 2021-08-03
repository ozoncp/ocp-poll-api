module github.com/ozoncp/ocp-poll-api

go 1.15

replace github.com/ozoncp/ocp-poll-api/internal/models => ../modules

require github.com/golang/mock v1.6.0

replace github.com/ozoncp/ocp-poll-api/internal/flusher => ./flusher

replace github.com/ozoncp/ocp-poll-api/internal/repo => ../repo
