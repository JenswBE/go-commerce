package config

import (
	"github.com/JenswBE/go-commerce/api/presenter"
	"github.com/JenswBE/go-commerce/fixtures"
	"github.com/JenswBE/go-commerce/utils/shortid"
)

func setupHandlerTest() *ConfigHandler {
	presenter := presenter.New(shortid.NewFakeService())
	handler := NewConfigHandler(presenter, fixtures.Config())
	return handler
}
