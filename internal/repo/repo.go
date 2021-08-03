package repo

import "github.com/ozoncp/ocp-poll-api/internal/models"

type Repo interface {
	AddEntities(entities []models.Poll) error
	ListEntities(limit, offset uint64) ([]models.Poll, error)
	DescribeEntity(entityId uint64) (*models.Poll, error)
}
