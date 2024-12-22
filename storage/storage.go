package storage

import (
	"app/models"
	"context"
)

type StorageI interface {
	CloseDb()
	User() UserRepoI
	Post() PostRepoI
}

type UserRepoI interface {
	Create(context.Context, *models.CreateUser) (string, error)
	GetById(context.Context, *models.UserPrimaryKey) (*models.User, error)
	GetByLoginAndPassword(context.Context, *models.Login) (*models.User, error)
	GetList(context.Context, *models.GetListUserRequest) (*models.GetListUserResponse, error)
	Update(context.Context, *models.UpdateUser) (int64, error)
	Delete(context.Context, *models.UserPrimaryKey) (int64, error)
}

type PostRepoI interface {
	Create(ctx context.Context, req *models.CreatePost) (string, error)
	GetById(ctx context.Context, req *models.PostPrimaryKey) (*models.Post, error)
	GetList(ctx context.Context, req *models.GetListPostRequest) (*models.GetListPostResponse, error)
	Update(ctx context.Context, req *models.UpdatePost) (int64, error)
	UpdateStatus(ctx context.Context, req *models.UpdateStatus) error
	Delete(ctx context.Context, req *models.PostPrimaryKey) (int64, error)
}
