package http

import (
	"errors"
	"go-gc-community/internal/response"
	"go-gc-community/internal/usecases"
	"go-gc-community/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	usecase *usecases.Usecases
}

func NewHealthHandler(usecase usecases.Usecases) *HealthHandler {
	return &HealthHandler{
		usecase: &usecase,
	}
}

func (hh *HealthHandler) Init(api *gin.RouterGroup) {
	health := api.Group("/health")
	{
		health.GET("", hh.Check)
	}
}

// @Summary Health Check
// @Tags default-health
// @Description This is the endpoint to check the system database
// @ModuleID Check
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response "Response indicates that the request succeeded and the resources has been fetched and transmitted in the message body"
// @Failure 500 {object} response.Response "failed Connect to the Database"
// @Router /api/health [get] 
func (h *HealthHandler) Check(ctx *gin.Context) {
	err := h.usecase.Health.Check()
	if err != nil {
		logger.Error("failed connect to the database")
		response.Error(ctx.Writer, http.StatusInternalServerError, "00", "00", errors.New("failed Connect to the Database"))
		return
	}

	response.Default(ctx.Writer, http.StatusOK, "00", response.SUCCESS_DEFAULT)
}