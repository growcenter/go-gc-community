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
		event.GET("/:id/session", h.SessionList)
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
// @Failure 400 {object} response.Response "There is something wrong with how user input the data"
// @Router api/v1.0/event/list [get]
func (eh *V1Handler) List(ctx *gin.Context) {
	event, time, err := eh.usecase.Event.GetAll()
	if err != nil {
		logger.Error(err)
		response.Error(ctx.Writer, http.StatusBadRequest, "01", "01", err)
		return
	}

	count := len(event)

	list := make([]models.EventResponseDetail, len(event))
	for i, p := range event {
		list[i] = models.EventResponseDetail{
			EventID: p.ID,
			EventName: p.Name,
			EventDescription: p.Description,
			EventCode: p.Code,
			OpenRegistration: p.OpenRegistration,
			ClosedRegistration: p.ClosedRegistration,
			Status: p.Status,
		}
	}

	response.Success(ctx.Writer, http.StatusOK, models.GetEventResponse{
		ResponseCode: fmt.Sprintf("%d%s%s", http.StatusOK, "00", "00"),
		ResponseMessage: response.SUCCESS_DEFAULT,
		EventCount: count,
		CurrentTime: time,
		Events: list,
	})
}

func (eh *V1Handler) SessionList(ctx *gin.Context) {
	eventId := ctx.Param("id")

	session, event, time, err := eh.usecase.Event.GetSessionByEvent(eventId)
	if err != nil {
		logger.Error(err)
		response.Error(ctx.Writer, http.StatusBadRequest, "01", "02", err)
		return
	}

	count := len(session)
	
	list := make([]models.SessionResponseDetail, len(session))
	for i, p := range session {
		list[i] = models.SessionResponseDetail{
			Name: p.Name,
			SessionID: p.ID,
			EventID: p.EventsId,
			Description: p.Description,
			Time: p.Time,
			MaxSeating: p.MaxSeating,
			Status: p.Status,
			OpenRegistration: p.OpenRegistration,
			ClosedRegistration: p.ClosedRegistration,
		}
	}

	response.Success(ctx.Writer, http.StatusOK, models.GetSessionResponse{
		ResponseCode: fmt.Sprintf("%d%s%s", http.StatusOK, "00", "00"),
		ResponseMessage: response.SUCCESS_DEFAULT,
		EventName: event.Name,
		EventId: event.ID,
		SessionCount: count,
		CurrentTime: time,
		Events: list,
	})
}