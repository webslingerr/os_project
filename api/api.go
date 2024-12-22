package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {
	handler := handler.NewHandler(cfg, store, logger)

	r.POST("/register", handler.RegisterUser)
	r.POST("/login", handler.LoginUser)

	r.GET("/user/:id", handler.GetByIdUser)

	r.POST("/post", handler.CreatePost)
	r.GET("/post/:id", handler.GetByIdPost)
	r.POST("/post/get-list", handler.GetListPost)
	r.PUT("/post/:id", handler.UpdatePost)
	r.PUT("/post/status/:id", handler.UpdateStatus)
	r.DELETE("/post/:id", handler.DeletePost)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
