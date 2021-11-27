package config

import (
	presenter "github.com/JenswBE/go-commerce/api/presenter/config"
	"github.com/gin-gonic/gin"
)

func (h *ConfigHandler) getConfig(c *gin.Context) {
	c.JSON(200, presenter.ConfigFromEntity(h.presenter, h.config))
}
