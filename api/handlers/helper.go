package handlers

import (
	"context"
	"strconv"
	"time"
	"ucode/ucode_go_chat_service/config"
	"ucode/ucode_go_chat_service/internal/socketio"

	"github.com/gin-gonic/gin"
)

func ParseLimitQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("limit", "10"))
}

func ParseOffsetQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("offset", "0"))
}

type sockErr struct {
	Function string `json:"function"`
	Message  string `json:"message"`
	Error    string `json:"error"`
	Request  any    `json:"request"`
}

func (s *socket) emitErr(sk *socketio.Socket, req sockErr) {
	s.log.Error(req)
	sk.Emit("error", sockErr{Message: req.Message, Function: req.Function})
}

func (s *socket) ctx() (context.Context, context.CancelFunc) {
	timeout := time.Duration(config.DBTimeout) * time.Second
	return context.WithTimeout(context.Background(), timeout)
}
