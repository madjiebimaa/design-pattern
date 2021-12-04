package repositories

import (
	"context"
	"database/sql"
	"go-design-pattern/repository-pattern/helpers"
	"go-design-pattern/repository-pattern/models"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

func (rp *ProductRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, product models.Product) models.Product {
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

func (rp *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product models.Product) models.Product {
	SQL := "SET products name = ?, price = ?, quantity = ?, created_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Quantity, product.CreatedAt, product.Id)
	helpers.PanicIfError(err)

	return product
}

func (rp *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []models.Product {
	SQL := "SELECT id, name, price, quantity, created_at FROM products"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.PanicIfError(err)
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.CreatedAt)
		products = append(products, product)
	}

	return products
}
