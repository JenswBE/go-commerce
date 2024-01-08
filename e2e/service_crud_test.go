//go:build e2e

package e2e

import (
	"context"
	"time"

	"github.com/samber/lo"
	"github.com/tebeka/selenium"
)

func (s *E2ETestSuite) TestServiceCategoryCRUD() {
	// No service categories should exist - GUI
	s.swdMustGetAdmin("service_categories")
	tableText := lo.Must(s.swdMustFindElement(selenium.ByCSSSelector, "table tbody tr td").Text())
	s.Require().Equal("Geen dienstsoorten gevonden", tableText, "Test should have been started with an empty DB")

	// No service categories should exist - API
	ctx := context.Background()
	rspServiceCategoriesList, rspRaw, err := s.apiClient.ServicesAPI.ListServiceCategories(ctx).Execute()
	s.Require().NoError(err, extractHTTPBody(s.T(), rspRaw))
	s.Require().Empty(rspServiceCategoriesList.GetServiceCategories(), "Test should have been started with an empty DB")

	// Move to add service categories page - GUI
	lo.Must0(s.swdMustFindElement(selenium.ByCSSSelector, "a.btn-success").Click())
	s.Require().Equal(s.adminURL("service_categories/new/"), lo.Must(s.swd.CurrentURL()))

	// Add new service category - GUI
	lo.Must0(s.swdMustFindElement(selenium.ByID, "inputName").SendKeys("Service category 1"))
	lo.Must0(s.swdMustFindElement(selenium.ByCSSSelector, "button.btn-success").Click())

	// Validate service category is added - GUI
	s.Require().Equal(s.adminURL("service_categories"), lo.Must(s.swd.CurrentURL()))
	serviceCategories := s.swdMustFindElements(selenium.ByCSSSelector, "table tbody tr")
	s.Require().Len(serviceCategories, 1, "Expected to find 1 element")
	serviceCategoryColumns := lo.Must(serviceCategories[0].FindElements(selenium.ByTagName, "td"))
	s.Require().Len(serviceCategoryColumns, 4, "Expected to find 4 columns for a service category")
	s.Require().Equal("Service category 1", lo.Must(serviceCategoryColumns[0].Text()))

	// Service category should exist - API
	rspServiceCategoriesList, rspRaw, err = s.apiClient.ServicesAPI.ListServiceCategories(ctx).Execute()
	s.Require().NoError(err, extractHTTPBody(s.T(), rspRaw))
	s.Require().Len(rspServiceCategoriesList.GetServiceCategories(), 1)
	serviceCategory := rspServiceCategoriesList.GetServiceCategories()[0]
	s.Require().Equal("Service category 1", serviceCategory.Name)

	// Delete service category - GUI
	lo.Must(serviceCategories[0].FindElement(selenium.ByCSSSelector, "button.btn-danger")).Click()
	s.Require().Contains(lo.Must(s.swd.AlertText()), "dienstsoort")
	s.Require().Contains(lo.Must(s.swd.AlertText()), "verwijderen")
	lo.Must0(s.swd.AcceptAlert())
	time.Sleep(100 * time.Millisecond)

	// Check service category deleted - GUI
	s.Require().Equal(s.adminURL("service_categories"), lo.Must(s.swd.CurrentURL()))
	tableText = lo.Must(s.swdMustFindElement(selenium.ByCSSSelector, "table tbody tr td").Text())
	s.Require().Equal("Geen dienstsoorten gevonden", tableText, "Service category should have been deleted")

	// Check service category deleted - API
	rspServiceCategoriesList, rspRaw, err = s.apiClient.ServicesAPI.ListServiceCategories(ctx).Execute()
	s.Require().NoError(err, extractHTTPBody(s.T(), rspRaw))
	s.Require().Empty(rspServiceCategoriesList.GetServiceCategories(), "Service category should have been deleted")
}

func (s *E2ETestSuite) TestServiceCRUD() {
	// Add new service category - GUI
	s.swdMustGetAdmin("service_categories/new/")
	lo.Must0(s.swdMustFindElement(selenium.ByID, "inputName").SendKeys("Service category 1"))
	lo.Must0(s.swdMustFindElement(selenium.ByCSSSelector, "button.btn-success").Click())

	// No services should exist - API
	ctx := context.Background()
	rspServiceCategoriesList, rspRaw, err := s.apiClient.ServicesAPI.ListServiceCategories(ctx).Execute()
	s.Require().NoError(err, extractHTTPBody(s.T(), rspRaw))
	serviceCategories := rspServiceCategoriesList.GetServiceCategories()
	s.Require().Len(serviceCategories, 1, "Service category should have been created")
	serviceCategory := serviceCategories[0]
	serviceCategoryID := lo.Must1(decodeBase58UUID(serviceCategory.Id))
	s.Require().Empty(serviceCategory.Services, "Service category should have been created without services")

	// Move to services page - GUI
	lo.Must0(s.swdMustFindElement(selenium.ByLinkText, "0").Click())
	s.Require().Equal(s.adminURL("service_categories", serviceCategoryID, "services"), lo.Must(s.swd.CurrentURL()))

	// No services should exist - GUI
	tableText := lo.Must(s.swdMustFindElement(selenium.ByCSSSelector, "table tbody tr td").Text())
	s.Require().Equal("Geen diensten gevonden", tableText)

	// Move to add service page - GUI
	lo.Must0(s.swdMustFindElement(selenium.ByCSSSelector, "a.btn-success").Click())
	s.Require().Equal(s.adminURL("service_categories", serviceCategoryID, "services/new/"), lo.Must(s.swd.CurrentURL()))

	// Add new service - GUI
	lo.Must0(s.swdMustFindElement(selenium.ByID, "inputName").SendKeys("Service 1"))
	lo.Must0(s.swdMustFindElement(selenium.ByID, "inputPrice").SendKeys("10.99"))
	lo.Must0(s.swdMustFindElement(selenium.ByID, "inputDescription").SendKeys("Test desc 1"))
	lo.Must0(s.swdMustFindElement(selenium.ByCSSSelector, "button.btn-success").Click())

	// Validate service is added - GUI
	s.Require().Equal(s.adminURL("service_categories", serviceCategoryID, "services/"), lo.Must(s.swd.CurrentURL()))
	services := s.swdMustFindElements(selenium.ByCSSSelector, "table tbody tr")
	s.Require().Len(services, 1, "Expected to find 1 element")
	serviceColumns := lo.Must(services[0].FindElements(selenium.ByTagName, "td"))
	s.Require().Len(serviceColumns, 5, "Expected to find 5 columns for an service")
	s.Require().Equal("Service 1", lo.Must(serviceColumns[0].Text()))

	// Service should exist - API
	rspServiceCategoriesList, rspRaw, err = s.apiClient.ServicesAPI.ListServiceCategories(ctx).Execute()
	s.Require().NoError(err, extractHTTPBody(s.T(), rspRaw))
	s.Require().Len(rspServiceCategoriesList.GetServiceCategories(), 1)
	serviceCategoryRetry := rspServiceCategoriesList.GetServiceCategories()[0]
	s.Require().Len(serviceCategoryRetry.GetServices(), 1)
	service := serviceCategoryRetry.GetServices()[0]
	s.Require().Equal("Service 1", service.Name)
	s.Require().EqualValues(1099, service.Price)
	s.Require().Equal("Test desc 1", service.Description)

	// Delete service - GUI
	lo.Must(services[0].FindElement(selenium.ByCSSSelector, "button.btn-danger")).Click()
	s.Require().Contains(lo.Must(s.swd.AlertText()), "dienst")
	s.Require().Contains(lo.Must(s.swd.AlertText()), "verwijderen")
	lo.Must0(s.swd.AcceptAlert())
	time.Sleep(100 * time.Millisecond)

	// Check service deleted - GUI
	s.Require().Equal(s.adminURL("service_categories", serviceCategoryID, "services/"), lo.Must(s.swd.CurrentURL()))
	tableText = lo.Must(s.swdMustFindElement(selenium.ByCSSSelector, "table tbody tr td").Text())
	s.Require().Equal("Geen diensten gevonden", tableText, "Service should have been deleted")

	// Check service deleted - API
	rspServiceCategoriesList, rspRaw, err = s.apiClient.ServicesAPI.ListServiceCategories(ctx).Execute()
	s.Require().NoError(err, extractHTTPBody(s.T(), rspRaw))
	s.Require().Len(rspServiceCategoriesList.GetServiceCategories(), 1)
	serviceCategoryRetry = rspServiceCategoriesList.GetServiceCategories()[0]
	s.Require().Empty(serviceCategoryRetry.Services, "Service should have been deleted")
}
