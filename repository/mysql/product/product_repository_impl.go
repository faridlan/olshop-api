package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/olshop-rest-api/helper"
	"github.com/faridlan/olshop-rest-api/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "insert into products(id, name, price, quantity, category_id, description, iamge_url, created_at) values (?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, product.Id, product.Name, product.Price, product.Quantity, product.CategoryId, product.Desription, product.ImageUrl, product.CreatedAt)
	helper.PanicIfErr(err)

	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "update products set name = ?, price = ?, quantity = ?, category_id = ?, description = ?, iamge_url = ?, updated_at where id = ?) values (?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.Quantity, product.CategoryId, product.Desription, product.ImageUrl, product.UpdatedAt, product.Id)
	helper.PanicIfErr(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "delete products where id = ? values (?)"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfErr(err)

}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId string) (domain.Product, error) {
	SQL := "select id, name, price, quantity, category_id, description, iamge_url, created_at, updated_at from products where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfErr(err)

	defer rows.Close()

	product := domain.Product{}
	if rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.CategoryId, &product.Desription, &product.ImageUrl, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicIfErr(err)

		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "select id, name, price, quantity, category_id, description, iamge_url, created_at, updated_at from products"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)

	defer rows.Close()

	products := []domain.Product{}
	for rows.Next() {
		product := domain.Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.CategoryId, &product.Desription, &product.ImageUrl, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicIfErr(err)

		products = append(products, product)
	}
	return products
}
