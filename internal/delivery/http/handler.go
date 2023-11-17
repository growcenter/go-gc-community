package http

import (
	"fmt"
	"go-gc-community/docs"
	"go-gc-community/internal/config"
	health "go-gc-community/internal/delivery/http/health"
	v1 "go-gc-community/internal/delivery/http/v1"
	"go-gc-community/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	usecase *usecases.Usecases
}

func NewHandler(usecase *usecases.Usecases) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

// @title GO-GC-COMMUNITY API DOCUMENTATION
// @version 1.0
// @description This is a go-gc-community api docs.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @schemes https

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors,
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Http.Host, cfg.Http.Port)
	if cfg.Environment != config.EnvLocal {
		docs.SwaggerInfo.Host = cfg.Http.Host
	}
	
	// Init Swagger
	if cfg.Environment != config.Prod {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Init router
	router.GET("/", func(ctx *gin.Context) {
		message := fmt.Sprintf("Welcome to %s", cfg.App.Name)
		ctx.String(http.StatusOK, message)
	})
	
	version1 := v1.NewV1Handler(*h.usecase)
	health := health.NewHealthHandler(*h.usecase)
	api := router.Group("/api")
	{
		version1.Init(api)
		health.Init(api)
	}

	//fs := http.FileServer(http.Dir("public"))
	//router.StaticFS("/lele", http.Dir("public/index.html"))
	//router.LoadHTMLGlob("../../public/*")
	router.LoadHTMLFiles("public/index.html")
	router.GET("lele", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", map[string]string{"title": "homepage"})
	})
	return router
}