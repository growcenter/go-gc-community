package v1

import (
	"fmt"
	"go-gc-community/internal/models"
	"go-gc-community/internal/response"
	"go-gc-community/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *V1Handler) eventRoutes(api *gin.RouterGroup) {
	event := api.Group("/event", h.Authorize)
	{
		event.GET("/list", h.List)
	}
}

// @Summary Event List
// @Tags event-list
// @Description This is the endpoint retrieve event list
// @ModuleID User
// @Accept  json
// @Produce  json
// @Success 201 {object} models.UserLoginResponse "Response indicates that the request succeeded and user account is created"
// @Success 200 {object} models.UserLoginResponse "Response indicates that the request succeeded and user is logged in"
// @Failure 422 {object} response.Response "There is something wrong with how user input the data"
// @Router api/v1.0/user/inquiry [get]
func (eh *V1Handler) List(ctx *gin.Context) {
	event, err := eh.usecase.Event.GetAll()
	if err != nil {
		logger.Error(err)
		response.Error(ctx.Writer, http.StatusBadRequest, "01", "01", err)
	}

	count := len(event)

	list := make([]models.EventResponseDetail, len(event))
	for i, p := range event {
		list[i] = models.EventResponseDetail{
			EventID: p.ID,
			EventName: p.Name,
			EventDescription: p.Description,
			EventCode: p.Code,
			Status: p.Status,
		}
	}

	response.Success(ctx.Writer, http.StatusOK, models.GetEventResponse{
		ResponseCode: fmt.Sprintf("%d%s%s", http.StatusOK, "00", "00"),
		ResponseMessage: response.SUCCESS_DEFAULT,
		EventCount: count,
		Events: list,
	})
}
