package v1

import (
	"go-gc-community/internal/usecases"

	"github.com/gin-gonic/gin"
)

type V1Handler struct {
	usecase *usecases.Usecases

}

func NewV1Handler(usecase usecases.Usecases) *V1Handler {
	return &V1Handler{
		usecase: &usecase,
	}
}

func (v1h *V1Handler) Init(api *gin.RouterGroup) {
	api = api.Group("/v1.0")
	{
		v1h.userRoutes(api)
	}
}