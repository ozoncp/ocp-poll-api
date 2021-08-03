package flusher

import (
	"github.com/ozoncp/ocp-poll-api/internal/models"
	"github.com/ozoncp/ocp-poll-api/internal/repo"
	"github.com/ozoncp/ocp-poll-api/internal/utils"
)

type Flusher interface {
	Flush(entities []models.Poll) ([]models.Poll, error)
}

func NewFlusher(chunkSize int, pollRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		pollRepo:  pollRepo,
	}
}

type flusher struct {
	chunkSize int
	pollRepo  repo.Repo
}

func (f *flusher) Flush(polls []models.Poll) ([]models.Poll, error) {
	slices, err := utils.SplitPoll(polls, f.chunkSize)
	if err != nil {
		return nil, err
	}
	for i := range slices {
		if err := f.pollRepo.AddEntities(slices[i]); err != nil {
			return slices[i], err
		}
	}
	return nil, nil
}
