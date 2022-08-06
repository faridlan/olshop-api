package categorysrv

import (
	"context"

	"github.com/faridlan/olshop-rest-api/model/web/categoryweb"
)

type CategoryService interface {
	Create(ctx context.Context, request categoryweb.CategoryRequest) categoryweb.CategoryResponse
	Update(ctx context.Context, request categoryweb.CategoryRequest) categoryweb.CategoryResponse
	Delete(ctx context.Context, categoryId string) categoryweb.CategoryResponse
	FindById(ctx context.Context, categoryId string) categoryweb.CategoryResponse
	FindAll(ctx context.Context) []categoryweb.CategoryResponse
}
