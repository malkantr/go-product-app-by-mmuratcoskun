package service

import (
	"errors"
	"fmt"
	"product-app/domain"
	"product-app/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

func (fakeRepository *FakeProductRepository) GetAllProducts() []domain.Product {
	return fakeRepository.products
}

func (fakeRepository *FakeProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	// todo: something something!!
	filteredProducts := []domain.Product{}

	for _, product := range fakeRepository.products {
		if product.Store == storeName {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func (fakeRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	fakeRepository.products = append(fakeRepository.products, domain.Product{
		Id:       int64(len(fakeRepository.products)) + 1,
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})
	return nil
}
func (fakeRepository *FakeProductRepository) GetById(productId int64) (domain.Product, error) {
	for _, product := range fakeRepository.products {
		if product.Id == productId {
			return product, nil
		}
	}
	return domain.Product{}, errors.New(fmt.Sprintf("Product not found with id %d", productId))
}
func (fakeRepository *FakeProductRepository) DeleteById(productId int64) error {
	return nil
}

func (fakeRepository *FakeProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	return nil
}
