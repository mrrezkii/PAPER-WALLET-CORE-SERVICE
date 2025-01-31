package usecase

import (
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
	"PAPER-WALLET-SERVICE-CORE/internal/repository"
	"context"
)

type (
	UserUsecase interface {
		Find(ctx context.Context) ([]domain.User, error)
		FindOne(ctx context.Context, userID string) (*domain.User, error)
		Update(ctx context.Context, user *domain.User) error
		SoftDelete(ctx context.Context, userID string) error
		HardDelete(ctx context.Context, userID string) error
	}

	userUsecase struct {
		repository repository.UserRepository
	}
)

func NewUserUsecase(repository repository.UserRepository) UserUsecase {
	return &userUsecase{repository}
}

func (u userUsecase) Find(ctx context.Context) ([]domain.User, error) {
	users, err := u.repository.Find(ctx, map[string]interface{}{
		"IsDeleted": 0,
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u userUsecase) FindOne(ctx context.Context, userID string) (*domain.User, error) {
	user, err := u.repository.FindOne(ctx, map[string]interface{}{
		"UserID":    userID,
		"IsDeleted": 0,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userUsecase) Update(ctx context.Context, user *domain.User) error {
	err := u.repository.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u userUsecase) SoftDelete(ctx context.Context, userID string) error {
	user, err := u.repository.FindOne(ctx, map[string]interface{}{
		"ID":        userID,
		"IsDeleted": 0,
	})
	if err != nil {
		return err
	}

	err = u.repository.SoftDelete(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u userUsecase) HardDelete(ctx context.Context, userID string) error {
	user, err := u.repository.FindOne(ctx, map[string]interface{}{
		"ID":        userID,
		"IsDeleted": 0,
	})
	if err != nil {
		return err
	}

	err = u.repository.HardDelete(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
