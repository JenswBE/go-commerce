//go:build e2e

package e2e

import (
	"context"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"github.com/tebeka/selenium"
)

func (s *E2ETestSuite) TestEventCRUD() {
	// No events should exist - GUI
	s.swdMustGetAdmin("events")
	tableText := lo.Must(s.swdMustFindElement(selenium.ByCSSSelector, "table tbody tr td").Text())
	require.Equal(s.T(), "Geen evenementen gevonden", tableText, "Test should have been started with an empty DB")

	// No events should exist - API
	ctx := context.Background()
	rspEventsList, rspRaw, err := s.apiClient.EventsApi.ListEvents(ctx).Execute()
	require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	require.Empty(s.T(), rspEventsList.GetEvents(), "Test should have been started with an empty DB")

	// Move to add event page - GUI
	lo.Must0(s.swdMustFindElement(selenium.ByCSSSelector, "a.btn-success").Click())
	require.Equal(s.T(), s.adminURL("events/new"), lo.Must(s.swd.CurrentURL()))

	// Add new event - GUI
	lo.Must0(s.swdMustFindElement(selenium.ByID, "inputName").SendKeys("Test event 1"))
	now := time.Now()
	start := now.AddDate(0, 0, 1)
	lo.Must0(s.swdMustFindElement(selenium.ByID, "inputStart").SendKeys(start.Format("2006-01-02")))
	end := now.AddDate(0, 0, 8)
	lo.Must0(s.swdMustFindElement(selenium.ByID, "inputEnd").SendKeys(end.Format("2006-01-02")))
	lo.Must0(s.swdMustFindElement(selenium.ByCSSSelector, "button.btn-success").Click())

	// Validate event is added - GUI
	require.Equal(s.T(), s.adminURL("events"), lo.Must(s.swd.CurrentURL()))
	events := s.swdMustFindElements(selenium.ByCSSSelector, "table tbody tr")
	require.Len(s.T(), events, 1, "Expected to find 1 element")
	eventColumns := lo.Must(events[0].FindElements(selenium.ByTagName, "td"))
	require.Len(s.T(), eventColumns, 4, "Expected to find 4 columns for an event")
	require.Equal(s.T(), "Test event 1", lo.Must(eventColumns[2].Text()))

	// Event should exist - API
	rspEventsList, rspRaw, err = s.apiClient.EventsApi.ListEvents(ctx).Execute()
	require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	require.Len(s.T(), rspEventsList.GetEvents(), 1)
	event := rspEventsList.GetEvents()[0]
	require.Equal(s.T(), "Test event 1", event.Name)
	require.NotNil(s.T(), event.WholeDay)
	require.True(s.T(), *event.WholeDay, "Expected event to be the whole day")
	require.Equal(s.T(), lo.T3(start.Date()), lo.T3(event.Start.Date()))
	require.Equal(s.T(), lo.T3(end.Date()), lo.T3(event.End.Date()))

	// Delete event - GUI
	lo.Must(events[0].FindElement(selenium.ByCSSSelector, "button.btn-danger")).Click()
	require.Contains(s.T(), lo.Must(s.swd.AlertText()), "evenement")
	require.Contains(s.T(), lo.Must(s.swd.AlertText()), "verwijderen")
	lo.Must0(s.swd.AcceptAlert())

	// Check event deleted - GUI
	require.Equal(s.T(), s.adminURL("events"), lo.Must(s.swd.CurrentURL()))
	tableText = lo.Must(s.swdMustFindElement(selenium.ByCSSSelector, "table tbody tr td").Text())
	require.Equal(s.T(), "Geen evenementen gevonden", tableText, "Event should have been deleted")

	// Check event deleted - API
	rspEventsList, rspRaw, err = s.apiClient.EventsApi.ListEvents(ctx).Execute()
	require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	require.Empty(s.T(), rspEventsList.GetEvents(), "Event should have been deleted")
}
