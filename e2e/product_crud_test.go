// +build e2e

package e2e

import (
	"context"

	"github.com/JenswBE/go-commerce/fixtures"
	"github.com/stretchr/testify/require"
)

func (s *E2ETestSuite) TestProductCRUD() {
	// Fixtures should contain 3 products
	rspProductList, rspRaw, err := s.publicClient.ProductsApi.ListProducts(context.Background()).Execute()
	require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	require.Len(s.T(), rspProductList.GetProducts(), 3)

	product := fixtures.ProductOpenAPI()
	rspProduct, rspRaw, err := s.authClient.ProductsApi.AddProduct(context.Background()).Product(*product).Execute()
	require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	require.NotNil(s.T(), rspProduct)
}
