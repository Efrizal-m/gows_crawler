package http

import (
	"github.com/gin-gonic/gin"
)

func (lc Liputan6Controller) RegisterRoutes(rg *gin.Engine) {
	newsroute := rg.Group("/crawl")
	// router.GET("/detik", controllers.CrawlDetik)
	newsroute.GET("/liputan6", lc.Crawl)

}
