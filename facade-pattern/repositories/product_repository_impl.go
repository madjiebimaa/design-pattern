package repositories

import (
	"context"
	"database/sql"
	"go-design-pattern/facade-pattern/models/domain"
	"go-design-pattern/repository-pattern/helpers"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

func (rp *ProductRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO products (name, price, quantity, created_at) VALUES (?, ?, ?, ?)"
	res, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Quantity, product.CreatedAt)
	helpers.PanicIfError(err)

	id, err := res.LastInsertId()
	helpers.PanicIfError(err)
	product.Id = uint(id)

	return product
}

func (rp *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productId uint) {
	SQL := "DELETE products WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, productId)
	helpers.PanicIfError(err)
}

func (rp *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "SET products name = ?, price = ?, quantity = ?, created_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Quantity, product.CreatedAt, product.Id)
	helpers.PanicIfError(err)

	return product
}

func (rp *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "SELECT id, name, price, quantity, created_at FROM products"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.CreatedAt)
		products = append(products, product)
	}

	return products
}
