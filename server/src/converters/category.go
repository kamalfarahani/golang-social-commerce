package converters

import (
	"kamal/server/social-commerce/server/src/models"
	"kamal/server/social-commerce/server/src/viewmodels"
)

func ConvertCategoryToView(category models.Category) viewmodels.Category {
	vmCategory := viewmodels.Category{
		Name:   category.Name,
		ImgUrl: category.ImgUrl,
		ID:     category.ID,
	}

	for _, product := range category.Products {
		vmCategory.Products = append(
			vmCategory.Products, ConvertProductToView(product))
	}

	return vmCategory
}

func ConvertCategoriesToViews(categories []models.Category) []viewmodels.Category {
	vmCategories := []viewmodels.Category{}

	for _, category := range categories {
		vmCategories = append(
			vmCategories, ConvertCategoryToView(category))
	}

	return vmCategories
}
