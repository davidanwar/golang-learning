package service

import (
	"context"
	"golang-restapi/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryRequest) web.CategoryResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
