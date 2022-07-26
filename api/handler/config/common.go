package config

import (
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/config"
	"github.com/gin-gonic/gin"
)

const pathPrefixConfig = "/config"

type ConfigHandler struct {
	presenter *presenter.Presenter
	config    config.Config
}

func NewConfigHandler(p *presenter.Presenter, config config.Config) *ConfigHandler {
	return &ConfigHandler{
		presenter: p,
		config:    config,
	}
}

func (h *ConfigHandler) RegisterRoutes(rg *gin.RouterGroup) {
	groupConfig := rg.Group(pathPrefixConfig)
	groupConfig.GET("/", h.getConfig)
}
