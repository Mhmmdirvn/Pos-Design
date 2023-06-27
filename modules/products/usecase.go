package products

import (
	"context"
	"time"
)

type UseCase struct {
	Repo Repository
}

func (usecase UseCase) GetAllProducts() ([]Product, error) {
	products, err := usecase.Repo.GetAllProducts()
	return products, err
}

func (usecase UseCase) GetProductById(ctx context.Context) (*Product, error) {
	idPrms := ctx.Value("idPrms")
	id := idPrms.(int)
	product, err := usecase.Repo.GetProductById(id)
	return product, err
}

func (usecase UseCase) CreateProduct(product *Product) error {
	err := usecase.Repo.CreateProduct(product)
	return err
}

func (usecase UseCase) UpdateProductById(ctx context.Context, product *Product) error {
	idPrms := ctx.Value("idPrms")
	id := idPrms.(int)
	products, err := usecase.Repo.GetProductById(id)
	if err != nil {
		return ErrProductIdNotFound
	}

	if products.DeletedAt != nil {
		return ErrPoductHasBeenRemoved
	}

	product.Id = id

	if err := usecase.Repo.Updates(product); err != nil {
		return ErrChangedStatus
	}

	return err
}

func (usecase UseCase) SoftDelete(ctx context.Context, status string) (*Product, error) {
	idPrms := ctx.Value("idPrms")
	id := idPrms.(int)
	product, err := usecase.Repo.GetProductById(id)
	if err != nil {
		return nil, err
	}

	if status == "active" {
		if product.DeletedAt == nil {
			return nil, ErrProductNotDeleted
		} else if product.DeletedAt != nil {
			product.DeletedAt = nil
		}
	} else if status == "inactive" {
		if product.DeletedAt == nil {
			deleteAt := time.Now()
			product.DeletedAt = &deleteAt
		} else if product.DeletedAt != nil {
			return nil, ErrProductAlreadyDeleted
		}
	} else {
		return nil, ErrInvalidStatus
	}

	if err := usecase.Repo.Save(product); err != nil {
		return nil, ErrChangedStatus
	}

	return product, nil
}

func (usecase UseCase) DeleteProductById(ctx context.Context) error {
	idPrms := ctx.Value("idPrms")
	id := idPrms.(int)
	err := usecase.Repo.DeleteProductById(id)
	return err
}