package repository

import (
	"PAPER-WALLET-SERVICE-CORE/config"
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
	"context"
)

type (
	UserRepository interface {
		Find(ctx context.Context, filter interface{}) ([]domain.User, error)
		FindOne(ctx context.Context, filter interface{}) (*domain.User, error)
		Upsert(ctx context.Context, user *domain.User) error
		Delete(ctx context.Context, user *domain.User) error
	}

	userRepository struct {
		filePath string
	}
)

func UserRepositoryNew(config *config.Config) UserRepository {
	return &userRepository{filePath: config.CSVFilePath}
}

func (u userRepository) Find(ctx context.Context, filter interface{}) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindOne(ctx context.Context, filter interface{}) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Upsert(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Delete(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}
