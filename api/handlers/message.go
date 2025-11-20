package handlers

import (
	"errors"
	"net/http"
	"ucode/ucode_go_chat_service/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) MessageGetList(c *gin.Context) {
	offset, err := ParseOffsetQueryParam(c)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
		return
	}
	if limit > 100 {
		limit = 100
	}

	roomId := c.Query("room_id")
	if roomId == "" {
		handleResponse(c, http.StatusBadRequest, errors.New("Room is required"))
		return
	}

	req := &models.GetListMessageReq{
		Offset: uint64(offset),
		Limit:  uint64(limit),
		RoomId: roomId,
	}

	messages, err := h.storage.Postgres().MessageGetList(
		c.Request.Context(), req,
	)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusOK, messages)
}
