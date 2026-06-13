package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/japnoor/daelog/internal/models"
	"github.com/japnoor/daelog/internal/services"
	"github.com/japnoor/daelog/pkg/constants"
	"github.com/japnoor/daelog/pkg/logger"
	"github.com/japnoor/daelog/pkg/response"
)

type EventHandler struct {
	service *services.EventService
}

func NewEventHandler(service *services.EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) Create(c *gin.Context) {
	var req models.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("invalid create event request").Err(err).Send()
		response.BadRequest(c, constants.ErrValidation, err.Error())
		return
	}

	event, err := h.service.Create(&req)
	if err != nil {
		logger.Error("failed to create event", err)
		response.InternalError(c, constants.ErrCreateFailed, "failed to create event")
		return
	}

	logger.Info("event created").Str("id", event.ID.Hex()).Str("type", string(event.EventType)).Send()
	response.Created(c, event)
}

func (h *EventHandler) GetByDate(c *gin.Context) {
	from, err := strconv.ParseInt(c.Query("from"), 10, 64)
	if err != nil {
		response.BadRequest(c, constants.ErrInvalidParam, "from is required and must be a timestamp")
		return
	}

	to, err := strconv.ParseInt(c.Query("to"), 10, 64)
	if err != nil {
		response.BadRequest(c, constants.ErrInvalidParam, "to is required and must be a timestamp")
		return
	}

	events, err := h.service.GetByDate(from, to)
	if err != nil {
		logger.Error("failed to fetch events", err)
		response.InternalError(c, constants.ErrFetchFailed, "failed to fetch events")
		return
	}

	logger.Info("events fetched").Int("count", len(events)).Send()
	response.OK(c, events)
}
