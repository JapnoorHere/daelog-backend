package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/japnoor/daelog/internal/services"
	"github.com/japnoor/daelog/pkg/constants"
	"github.com/japnoor/daelog/pkg/logger"
	"github.com/japnoor/daelog/pkg/response"
)

type SessionHandler struct {
	service *services.SessionService
}

func NewSessionHandler(service *services.SessionService) *SessionHandler {
	return &SessionHandler{service: service}
}

func (h *SessionHandler) GetByDate(c *gin.Context) {
	from, err := strconv.ParseInt(c.Query("from"), 10, 64)
	if err != nil {
		response.BadRequest(c, constants.ErrInvalidParam, "from is required and must be a unix ms timestamp")
		return
	}

	to, err := strconv.ParseInt(c.Query("to"), 10, 64)
	if err != nil {
		response.BadRequest(c, constants.ErrInvalidParam, "to is required and must be a unix ms timestamp")
		return
	}

	sessions, err := h.service.GetByDate(from, to)
	if err != nil {
		logger.Error("failed to fetch sessions", err)
		response.InternalError(c, constants.ErrFetchFailed, "failed to fetch sessions")
		return
	}

	logger.Info("sessions fetched").Int("count", len(sessions)).Send()
	response.OK(c, sessions)
}
