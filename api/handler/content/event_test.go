package content

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/JenswBE/go-commerce/api/handler"
	"github.com/JenswBE/go-commerce/entities"
	"github.com/JenswBE/go-commerce/fixtures"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_listEvents_Success(t *testing.T) {
	// Setup test
	c, r := handler.SetupGinTest(t, "GET", "", nil, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("ListEvents", mock.Anything).Return(fixtures.EventSlice(), nil)

	// Call handler
	productHandler.listEvents(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.EventListOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "ListEvents")
}

func Test_getEvent_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.EventID}}
	c, r := handler.SetupGinTest(t, "GET", "", params, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("GetEvent", mock.Anything, mock.Anything).Return(fixtures.Event(), nil)

	// Call handler
	productHandler.getEvent(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.EventOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "GetEvent", uuid.MustParse(fixtures.EventID))
}

func Test_createEvent_Success(t *testing.T) {
	// Setup test
	body, err := json.Marshal(fixtures.EventOpenAPI())
	require.NoError(t, err)
	c, r := handler.SetupGinTest(t, "POST", "", nil, body)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("CreateEvent", mock.Anything).Return(fixtures.Event(), nil)

	// Call handler
	productHandler.createEvent(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.EventOpenAPI(), r)
	require.Equal(t, http.StatusCreated, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "CreateEvent", mock.MatchedBy(func(actual *entities.Event) bool {
		expected := fixtures.Event()
		expected.ID = uuid.Nil
		require.Equal(t, expected, actual)
		return true
	}))
}

func Test_updateEvent_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.EventID}}
	body, err := json.Marshal(fixtures.EventOpenAPI())
	require.NoError(t, err)
	c, r := handler.SetupGinTest(t, "PUT", "", params, body)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("UpdateEvent", mock.Anything).Return(fixtures.Event(), nil)

	// Call handler
	productHandler.updateEvent(c)

	// Assert result
	handler.RequireEqualJSON(t, fixtures.EventOpenAPI(), r)
	require.Equal(t, http.StatusOK, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "UpdateEvent", mock.MatchedBy(func(actual *entities.Event) bool {
		require.Equal(t, fixtures.Event(), actual)
		return true
	}))
}

func Test_deleteEvent_Success(t *testing.T) {
	// Setup test
	params := gin.Params{{Key: "id", Value: fixtures.EventID}}
	c, r := handler.SetupGinTest(t, "DELETE", "", params, nil)
	productHandler, usecaseMock := setupHandlerTest()
	usecaseMock.On("DeleteEvent", mock.Anything).Return(nil)

	// Call handler
	productHandler.deleteEvent(c)

	// Assert result
	require.Empty(t, r.Body.String())
	require.Equal(t, http.StatusNoContent, r.Code)

	// Assert mock calls
	usecaseMock.AssertCalled(t, "DeleteEvent", uuid.MustParse(fixtures.EventID))
}
