package categorysrv

import (
	"context"
	"database/sql"

	"github.com/faridlan/olshop-rest-api/model/web/categoryweb"
	"github.com/faridlan/olshop-rest-api/repository/mysql/category"
)

type CategoryServiceImpl struct {
	CategoryRepo category.CategoryRepository
	DB           *sql.DB
}

func NewCategoryService(categoryRepo category.CategoryRepository, Db *sql.DB) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepo: categoryRepo,
		DB:           Db,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request categoryweb.CategoryRequest) categoryweb.CategoryResponse {
	panic("not implemented") // TODO: Implement
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request categoryweb.CategoryRequest) categoryweb.CategoryResponse {
	panic("not implemented") // TODO: Implement
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId string) categoryweb.CategoryResponse {
	panic("not implemented") // TODO: Implement
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId string) categoryweb.CategoryResponse {
	panic("not implemented") // TODO: Implement
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []categoryweb.CategoryResponse {
	panic("not implemented") // TODO: Implement
}
