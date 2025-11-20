package handlers

import (
	"net/http"
	"ucode/ucode_go_chat_service/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) RoomMemberCreate(c *gin.Context) {
	req := &models.CreateRoomMember{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
		return
	}

	roomMember, err := h.storage.Postgres().RoomMemberCreate(
		c.Request.Context(),
		req,
	)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, roomMember)
}
