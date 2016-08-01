package converters

import (
	"kamal/server/social-commerce/server/src/models"
	"kamal/server/social-commerce/server/src/viewmodels"
)

func ConvertProductToView(product models.Product) viewmodels.Product {
	vmProduct := viewmodels.Product{
		Name:        product.Name,
		Description: product.Description,
		ImgUrl:      product.ImgUrl,
		Price:       product.Price,
		ID:          product.ID,
	}
	return vmProduct
}

func ConvertProductsToViews(products []models.Product) []viewmodels.Product {
	vmProducts := []viewmodels.Product{}

	for _, product := range products {
		vmProducts = append(
			vmProducts, ConvertProductToView(product))
	}

	return vmProducts
}
