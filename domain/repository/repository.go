package domain

import (
	"context"
	"gows_crawler/domain/models"
)

type Liputan6RepositoryI interface {
	GetAllData(ctx context.Context) (user []models.Liputan6News, err error)
	InsertData(ctx context.Context, req *models.Liputan6News) error
}
