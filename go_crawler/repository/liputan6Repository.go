package repository

import (
	"context"
	"gows_crawler/domain/models"
	"log"

	domainR "gows_crawler/domain/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Liputan6Repository struct {
	mongoDB *mongo.Database
}

func (c Liputan6Repository) GetAllData(ctx context.Context) (liputan6Resp []models.Liputan6News, err error) {
	query, err := c.mongoDB.Collection("liputan6").Find(ctx, bson.D{})
	if err != nil {
		log.Println("error", err)
		return []models.Liputan6News{}, err
	}
	defer query.Close(ctx)

	listDataLiputan6News := make([]models.Liputan6News, 0)
	for query.Next(ctx) {
		var row models.Liputan6News
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataLiputan6News = append(listDataLiputan6News, row)
	}

	return listDataLiputan6News, err
}

func (c Liputan6Repository) InsertData(ctx context.Context, req *models.Liputan6News) error {
	_, err := c.mongoDB.Collection("liputan6").InsertOne(ctx, req)
	if err != nil {
		log.Println("error")
	}

	return err
}

func NewLiputan6Repository(mongo *mongo.Database) domainR.Liputan6RepositoryI {
	return &Liputan6Repository{mongoDB: mongo}
}
