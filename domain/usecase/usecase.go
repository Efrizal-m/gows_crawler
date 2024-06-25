package domain

import (
	"context"
	"gows_crawler/domain/models"
)

type Liputan6UsecaseI interface {
	CrawlNews(ctx context.Context, req *models.Liputan6News) error
	GetAll(ctx context.Context) ([]models.Liputan6News, error)
}
