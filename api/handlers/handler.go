package handlers

import (
	"ucode/ucode_go_chat_service/config"
	"ucode/ucode_go_chat_service/pkg/logger"
	"ucode/ucode_go_chat_service/storage"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type handler struct {
	log     *logger.Logger
	cfg     config.Config
	storage storage.StorageI
}

type HandlerConfig struct {
	Logger   *logger.Logger
	Cfg      config.Config
	Postgres storage.StorageI
}

type Response struct {
	Body  any    `json:"body"`
	Error string `json:"error"`
}

func New(c *HandlerConfig) *handler {
	return &handler{
		log:     c.Logger,
		cfg:     c.Cfg,
		storage: c.Postgres,
	}
}

func handleResponse(c *gin.Context, status int, data any) {
	if status >= 400 {
		c.JSON(status, Response{
			Error: cast.ToString(data),
		})
		return
	}

	c.JSON(status, Response{
		Body: data,
	})
}
