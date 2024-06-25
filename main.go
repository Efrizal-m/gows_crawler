package main

import (
	"context"
	"fmt"
	"log"

	"gows_crawler/config"
	"gows_crawler/delivery/http"
	domainU "gows_crawler/domain/usecase"
	"gows_crawler/repository"
	"gows_crawler/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server      *gin.Engine
	lu          domainU.Liputan6UsecaseI
	lc          http.Liputan6Controller
	ctx         context.Context
	mongoClient *mongo.Client
)

func init() {
	ctx = context.TODO()

	// Mongo
	mongoCon, err := config.Connect()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("mongo connection established")
	lr := repository.NewLiputan6Repository(mongoCon)
	lu = usecase.NewLiputan6Usecase(lr, ctx)
	lc = http.New(lu)
	server = gin.Default()
}

func main() {
	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}(mongoClient, ctx)

	lc.RegisterRoutes(server)

	log.Fatal(server.Run(":9090"))

}
