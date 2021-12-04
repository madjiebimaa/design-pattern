package facades

import (
	"context"
	"database/sql"
	"go-design-pattern/facade-pattern/models/web"
	"go-design-pattern/facade-pattern/repositories"
	"go-design-pattern/repository-pattern/helpers"
)

type ProductFacadeImpl struct {
	DB                *sql.DB
	ProductRepository *repositories.ProductRepositoryImpl
}

func NewProductFacade(db *sql.DB, productRepository *repositories.ProductRepositoryImpl) *ProductFacadeImpl {
	return &ProductFacadeImpl{
		DB:                db,
		ProductRepository: productRepository,
	}
}

func (fc *ProductFacadeImpl) BuyProduct(ctx context.Context, request web.ProductToBuyRequest) web.ProductToBuyResponse {
	tx, err := fc.DB.Begin()
	helpers.PanicIfError(err)

	// Create bill transaction

	// Update quantity value in product
	request.Product.Quantity -= request.Product.Quantity - request.Quantity
	product := fc.ProductRepository.Update(ctx, tx, request.Product)
	// Customer pay the product to seller

	return web.ProductToBuyResponse{
		Product: product,
	}
}
