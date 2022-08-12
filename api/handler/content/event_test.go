package content

import (
	"net/http"
	"testing"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/fixtures"
	"github.com/JenswBE/go-commerce/utils/generics"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_listEvents_Success(t *testing.T) {
	// Setup test
	c, r := handler.SetupGinTest(t, "GET", "?include_past_events=True", nil, nil)
	contentHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("ListEvents", mock.Anything).Return(fixtures.EventSlice(), nil)

	// Call handler
	contentHandler.listEvents(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.EventListOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "ListEvents", true)
}

func Test_getEvent_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.EventID}}
	c, r := handler.SetupGinTest(t, "GET", "", params, nil)
	contentHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("GetEvent", mock.Anything, mock.Anything).Return(fixtures.Event(), nil)

	// Call handler
	contentHandler.getEvent(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.EventOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "GetEvent", generics.Must(entities.NewIDFromString(fixtures.EventID)))
}
