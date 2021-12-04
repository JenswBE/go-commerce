package content

import (
	"github.com/JenswBE/go-commerce/api/presenter"
	mocks "github.com/JenswBE/go-commerce/mocks/usecases/content"
	"github.com/JenswBE/go-commerce/utils/sanitizer"
	"github.com/JenswBE/go-commerce/utils/shortid"
)

func setupHandlerTest() (*ContentHandler, *mocks.Usecase) {
	presenter := presenter.New(shortid.NewFakeService(), sanitizer.NewFakeService())
	usecase := &mocks.Usecase{}
	handler := NewContentHandler(presenter, usecase)
	return handler, usecase
}
