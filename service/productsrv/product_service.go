package productsrv

import (
	"context"

	"github.com/faridlan/olshop-rest-api/model/web/productweb"
)

type ProductService interface {
	Create(ctx context.Context, request productweb.ProductRequest) productweb.ProductResponse
	Update(ctx context.Context, request productweb.ProductRequest) productweb.ProductResponse
	Delete(ctx context.Context, productId string)
	FindById(ctx context.Context, productId string) productweb.ProductResponse
	FindAll(ctx context.Context) []productweb.ProductResponse
}
