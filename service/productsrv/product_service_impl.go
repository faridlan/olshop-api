package productsrv

import (
	"context"
	"database/sql"

	"github.com/faridlan/olshop-rest-api/helper"
	"github.com/faridlan/olshop-rest-api/helper/model"
	"github.com/faridlan/olshop-rest-api/model/domain"
	"github.com/faridlan/olshop-rest-api/model/web/productweb"
	"github.com/faridlan/olshop-rest-api/repository/mysql/product"
)

type ProductServiceImpl struct {
	ProductRepo product.ProductRepository
	DB          *sql.DB
}

func NewProductService(productrepo product.ProductRepository, DB *sql.DB) ProductService {
	return &ProductServiceImpl{
		ProductRepo: productrepo,
		DB:          DB,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request productweb.ProductRequest) productweb.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		Id:         request.Id,
		Name:       request.Name,
		Price:      request.Price,
		Quantity:   request.Quantity,
		CategoryId: request.CategoryId,
		Desription: request.Desription,
		ImageUrl:   request.ImageUrl,
		CreatedAt:  request.CreatedAt,
	}

	product = service.ProductRepo.Save(ctx, tx, product)
	return model.ToProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request productweb.ProductRequest) productweb.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepo.FindById(ctx, tx, request.Id)
	helper.PanicIfErr(err)

	product.Id = request.Id
	product.Name = request.Name
	product.Price = request.Price
	product.Quantity = request.Quantity
	product.Desription = request.Desription
	product.ImageUrl = request.ImageUrl
	product.UpdatedAt = request.UpdatedAt

	product = service.ProductRepo.Update(ctx, tx, product)

	return model.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepo.FindById(ctx, tx, productId)
	helper.PanicIfErr(err)

	service.ProductRepo.Delete(ctx, tx, product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId string) productweb.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepo.FindById(ctx, tx, productId)
	helper.PanicIfErr(err)

	return model.ToProductResponse(product)

}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []productweb.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepo.FindAll(ctx, tx)

	return model.ToProductResponses(products)
}
