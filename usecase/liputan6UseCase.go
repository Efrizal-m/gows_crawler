package usecase

import (
	"context"
	domainR "gows_crawler/domain/repository"
	domainU "gows_crawler/domain/usecase"

	"gows_crawler/domain/models"
	engine "gows_crawler/services"

	"log"
)

type liputan6ServiceImpl struct {
	liputan6Repo domainR.Liputan6RepositoryI
	ctx          context.Context
}

func (u liputan6ServiceImpl) CrawlNews(ctx context.Context, req *models.Liputan6News) error {
	//check if context is nil
	if ctx == nil {
		ctx = context.Background()
	}

	liputan6News := engine.CrawlLiputan6()

	//insert data
	for _, v := range liputan6News {
		err := u.liputan6Repo.InsertData(ctx, v)
		if err != nil {
			return err
		}
	}
	log.Println("Successfully Inserted Data Liputan6")

	return nil
}

func (u liputan6ServiceImpl) GetAll(ctx context.Context) ([]models.Liputan6News, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := u.liputan6Repo.GetAllData(ctx)
	if err != nil {
		log.Println("failed to show data news with default log")
		return list, err
	}

	return list, err
}

func NewLiputan6Usecase(liputan6Repo domainR.Liputan6RepositoryI, ctx context.Context) domainU.Liputan6UsecaseI {
	return &liputan6ServiceImpl{
		liputan6Repo: liputan6Repo,
		ctx:          ctx,
	}
}
