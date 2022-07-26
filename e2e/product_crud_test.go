//go:build e2e

package e2e

// var imageConfigs = []string{"100_100_FIT"} // See config file

func (s *E2ETestSuite) TestProductCRUD() {
	// // Get products before AddProduct
	// ctx := context.Background()
	// rspProductList, rspRaw, err := s.publicClient.ProductsApi.ListProducts(ctx).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	// productCountBeforeAdd := len(rspProductList.GetProducts())

	// // Add a new product
	// product := fixtures.ProductOpenAPI()
	// product.CategoryIds = nil
	// product.ManufacturerId = nil
	// rspProduct, rspRaw, err := s.authClient.ProductsApi.AddProduct(ctx).Product(*product).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	// require.NotNil(s.T(), rspProduct)
	// productID := rspProduct.GetId()

	// // Get resolved product for later comparison
	// rspResolvedProduct, rspRaw, err := s.publicClient.ProductsApi.GetProduct(ctx, productID).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	// require.Equal(s.T(), productID, rspResolvedProduct.GetId())

	// // Refetch products after AddProduct
	// rspProductListAfterAdd, rspRaw, err := s.publicClient.ProductsApi.ListProducts(ctx).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	// require.Len(s.T(), rspProductListAfterAdd.GetProducts(), productCountBeforeAdd+1)

	// // Add initial images
	// file1, err := os.Open("assets/image-1.png")
	// defer file1.Close()
	// require.NoError(s.T(), err)
	// file2, err := os.Open("assets/image-2.png")
	// defer file2.Close()
	// require.NoError(s.T(), err)
	// rspProductImages, rspRaw, err := s.authClient.ProductsApi.AddProductImages(ctx, productID).File([]*os.File{file1, file2}).Img(imageConfigs).Execute()
	// require.Len(s.T(), rspProductImages.GetImages(), 2)

	// // Refetch product after AddProductImages
	// rspResolvedProductRefetch, rspRaw, err := s.publicClient.ProductsApi.GetProduct(ctx, productID).Img(imageConfigs).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	// require.Len(s.T(), rspResolvedProductRefetch.ImageUrls, 2)
	// require.Empty(s.T(), cmp.Diff(rspResolvedProduct, rspResolvedProductRefetch, ignoreResolvedProductFields("ImageUrls")))

	// // Add more images
	// file3, err := os.Open("assets/image-3.png")
	// defer file3.Close()
	// require.NoError(s.T(), err)
	// rspProductImages, rspRaw, err = s.authClient.ProductsApi.AddProductImages(ctx, productID).File([]*os.File{file3}).Img(imageConfigs).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	// require.Len(s.T(), rspProductImages.GetImages(), 3)

	// // Refetch product after AddProductImages
	// rspResolvedProductRefetch, rspRaw, err = s.publicClient.ProductsApi.GetProduct(ctx, productID).Img(imageConfigs).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	// require.Len(s.T(), rspResolvedProductRefetch.ImageUrls, 3)
	// require.Empty(s.T(), cmp.Diff(rspResolvedProduct, rspResolvedProductRefetch, ignoreResolvedProductFields("ImageUrls")))

	// // Refetch product images
	// rspProductImagesRefetch, rspRaw, err := s.publicClient.ProductsApi.ListProductImages(ctx, productID).Img(imageConfigs).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	// require.Equal(s.T(), rspProductImages, rspProductImagesRefetch)

	// // Delete an image
	// rspRaw, err = s.authClient.ProductsApi.DeleteProductImage(ctx, productID, rspProductImagesRefetch.Images[1].Id).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))

	// // Refetch product images
	// rspProductImagesRefetch, rspRaw, err = s.publicClient.ProductsApi.ListProductImages(ctx, productID).Img(imageConfigs).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	// expected := openapi.NewImageList([]openapi.Image{rspProductImages.Images[0], rspProductImages.Images[2]})
	// require.Equal(s.T(), expected, rspProductImagesRefetch)

	// // Update product
	// updatedProduct := *rspProduct
	// updatedProduct.Name = "New name"
	// rspUpdatedProduct, rspRaw, err := s.authClient.ProductsApi.UpdateProduct(ctx, productID).Product(updatedProduct).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	// require.Empty(s.T(), cmp.Diff(&updatedProduct, rspUpdatedProduct, ignoreProductFields()))

	// // Refetch product after UpdateProduct
	// rspResolvedProductRefetch, rspRaw, err = s.publicClient.ProductsApi.GetProduct(ctx, productID).Img(imageConfigs).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	// require.Len(s.T(), rspResolvedProductRefetch.ImageUrls, 2)
	// rspResolvedProduct.Name = updatedProduct.Name
	// require.Empty(s.T(), cmp.Diff(rspResolvedProduct, rspResolvedProductRefetch, ignoreResolvedProductFields("ImageUrls")))

	// // Delete product
	// rspRaw, err = s.authClient.ProductsApi.DeleteProduct(ctx, productID).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))

	// // Refetch products after DeleteProduct
	// rspProductListAfterDelete, rspRaw, err := s.publicClient.ProductsApi.ListProducts(ctx).Execute()
	// require.NoError(s.T(), err, extractHTTPBody(s.T(), rspRaw))
	// require.Len(s.T(), rspProductListAfterDelete.GetProducts(), productCountBeforeAdd)

	// // Get product should return 404
	// _, rspRaw, err = s.publicClient.ProductsApi.GetProduct(ctx, productID).Img(imageConfigs).Execute()
	// require.NotNil(s.T(), rspRaw, err.Error())
	// require.Equal(s.T(), http.StatusNotFound, rspRaw.StatusCode)
}

// func ignoreFields(typ interface{}, fields ...string) cmp.Option {
// 	ignoredFields := make([]string, 0, len(fields)+2)
// 	ignoredFields = append(ignoredFields, "CreatedAt", "UpdatedAt")
// 	ignoredFields = append(ignoredFields, fields...)
// 	return cmpopts.IgnoreFields(typ, ignoredFields...)
// }

// func ignoreProductFields(fields ...string) cmp.Option {
// 	return ignoreFields(openapi.Product{}, fields...)
// }

// func ignoreResolvedProductFields(fields ...string) cmp.Option {
// 	return ignoreFields(openapi.ResolvedProduct{}, fields...)
// }
