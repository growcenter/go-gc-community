package v1

import (
	"go-gc-community/internal/usecases"
	"go-gc-community/pkg/authorization"

	"github.com/gin-gonic/gin"
)

type V1Handler struct {
	usecase *usecases.Usecases
	authorization *authorization.Auth

}

func NewV1Handler(usecase usecases.Usecases, authorization authorization.Auth) *V1Handler {
	return &V1Handler{
		usecase: &usecase,
		authorization: &authorization,
	}
}

func (v1h *V1Handler) Init(api *gin.RouterGroup) {
	api = api.Group("/v1.0")
	{
		v1h.userRoutes(api)
		v1h.eventRoutes(api)
		v1h.internalRoutes(api)
	}
}