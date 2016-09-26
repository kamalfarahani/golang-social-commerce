package converters

import (
	"../models"
	"../viewmodels"
)

func ConvertCollectionToView(collection models.Collection) viewmodels.Collection {
	vmCollection := viewmodels.Collection{
		ImgUrl: collection.ImgUrl,
		Name:   collection.Name,
		ID:     collection.ID,
	}

	for _, product := range collection.Products {
		vmCollection.Products = append(
			vmCollection.Products, ConvertProductToView(product))
	}

	return vmCollection
}

func ConvertCollectionsToViews(collections []models.Collection) []viewmodels.Collection {
	vmCollections := []viewmodels.Collection{}

	for _, collection := range collections {
		vmCollections = append(
			vmCollections, ConvertCollectionToView(collection))
	}

	return vmCollections
}
