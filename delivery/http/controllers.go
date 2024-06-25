package http

import (
	domainU "gows_crawler/domain/usecase"

	"gows_crawler/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Liputan6Controller struct {
	Liputan6Usecase domainU.Liputan6UsecaseI
}

func New(liputan6service domainU.Liputan6UsecaseI) Liputan6Controller {
	return Liputan6Controller{
		Liputan6Usecase: liputan6service,
	}
}

func (lc *Liputan6Controller) Crawl(ctx *gin.Context) {
	var user models.Liputan6News
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := lc.Liputan6Usecase.CrawlNews(ctx, &user)

	var (
		result gin.H
	)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"message": "processing",
		}
	}
	ctx.JSON(http.StatusOK, result)
}

func (lc *Liputan6Controller) GetAll(ctx *gin.Context) {
	liputan6News, err := lc.Liputan6Usecase.GetAll(ctx)
	var (
		result gin.H
	)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"data": liputan6News,
		}
	}
	ctx.JSON(http.StatusOK, result)
}
