package repository

import (
	"PAPER-WALLET-SERVICE-CORE/config"
	"PAPER-WALLET-SERVICE-CORE/helper"
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
	"context"
	"fmt"
	"time"
)

type (
	UserRepository interface {
		Find(ctx context.Context, filter map[string]interface{}) ([]domain.User, error)
		FindOne(ctx context.Context, filter map[string]interface{}) (*domain.User, error)
		Upsert(ctx context.Context, user *domain.User) error
		SoftDelete(ctx context.Context, user *domain.User) error
		HardDelete(ctx context.Context, user *domain.User) error
	}

	userRepository struct {
		filePath string
	}
)

func UserRepositoryNew(config *config.Config) UserRepository {
	return &userRepository{filePath: config.CSVFilePath}
}

func (u userRepository) Find(ctx context.Context, filter map[string]interface{}) ([]domain.User, error) {
	records, err := u.readCSVFile()
	if err != nil {
		return nil, err
	}

	return u.findUsersByFilter(records, filter)
}

func (u userRepository) FindOne(ctx context.Context, filter map[string]interface{}) (*domain.User, error) {
	records, err := u.readCSVFile()
	if err != nil {
		return nil, err
	}

	users, err := u.findUsersByFilter(records, filter)
	if err != nil || len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}

func (u userRepository) upsertUserRecord(records [][]string, user *domain.User) ([][]string, bool) {
	var updated bool
	var updatedRecords [][]string

	for _, record := range records[1:] {
		existingUser, err := helper.MapRecordToUser(record)
		if err != nil {
			continue
		}

		if existingUser.ID == user.ID {
			updatedRecords = append(updatedRecords, []string{
				user.ID,
				user.Name,
				user.Currency,
				fmt.Sprintf("%d", user.Scale),
				user.Balance.String(),
				user.CreatedBy,
				user.CreatedDate.Format(time.RFC3339),
				user.UpdatedBy,
				user.UpdatedDate.Format(time.RFC3339),
				fmt.Sprintf("%d", user.Version),
				fmt.Sprintf("%d", user.IsDeleted),
			})
			updated = true
		} else {
			updatedRecords = append(updatedRecords, record)
		}
	}

	if !updated {
		updatedRecords = append(updatedRecords, []string{
			user.ID,
			user.Name,
			user.Currency,
			fmt.Sprintf("%d", user.Scale),
			user.Balance.String(),
			user.CreatedBy,
			user.CreatedDate.Format(time.RFC3339),
			user.UpdatedBy,
			user.UpdatedDate.Format(time.RFC3339),
			fmt.Sprintf("%d", user.Version),
			fmt.Sprintf("%d", user.IsDeleted),
		})
	}

	return updatedRecords, updated
}

func (u userRepository) Upsert(ctx context.Context, user *domain.User) error {
	records, err := u.readCSVFile()
	if err != nil {
		return err
	}

	updatedRecords, _ := u.upsertUserRecord(records, user)

	if err := u.writeCSVFile(updatedRecords); err != nil {
		return err
	}

	return nil
}

func (u userRepository) SoftDelete(ctx context.Context, user *domain.User) error {
	records, err := u.readCSVFile()
	if err != nil {
		return err
	}

	updatedRecords, err := u.softDeleteUserRecord(records, user)
	if err != nil {
		return err
	}

	if err := u.writeCSVFile(updatedRecords); err != nil {
		return err
	}

	return nil
}

func (u userRepository) HardDelete(ctx context.Context, user *domain.User) error {
	records, err := u.readCSVFile()
	if err != nil {
		return err
	}

	var updatedRecords [][]string

	for _, record := range records[1:] {
		existingUser, err := helper.MapRecordToUser(record)
		if err != nil {
			continue
		}

		if existingUser.ID == user.ID {
			continue
		}

		updatedRecords = append(updatedRecords, record)
	}

	if len(updatedRecords) == len(records)-1 {
		return fmt.Errorf("user with ID %s not found", user.ID)
	}

	if err := u.writeCSVFile(updatedRecords); err != nil {
		return err
	}

	return nil
}
