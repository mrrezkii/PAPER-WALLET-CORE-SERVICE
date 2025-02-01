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
		Create(ctx context.Context, user *domain.User) error
		Update(ctx context.Context, user *domain.User) error
		SoftDelete(ctx context.Context, user *domain.User) error
		HardDelete(ctx context.Context, user *domain.User) error
	}

	userRepository struct {
		filePath string
	}
)

func NewUserRepository(config *config.Config) UserRepository {
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

func (u userRepository) Create(ctx context.Context, user *domain.User) error {
	records, err := u.readCSVFile()
	if err != nil {
		return err
	}

	var exist bool
	var newRecords [][]string

	for _, record := range records[1:] {
		existingUser, err := helper.MapRecordToUser(record)
		if err != nil {
			continue
		}

		if existingUser.ID == user.ID {
			exist = true
		}
	}

	if exist {
		return fmt.Errorf("user with ID %s already exists", user.ID)
	}

	newRecords = append(newRecords, []string{
		user.ID,
		user.Name,
		user.Currency,
		fmt.Sprintf("%d", user.Scale),
		user.Balance.String(),
		"system",
		time.Now().UTC().Format("2006-01-02 15:04:05"),
		"system",
		time.Now().UTC().Format("2006-01-02 15:04:05"),
		fmt.Sprintf("%d", 1),
		fmt.Sprintf("%d", 0),
	})

	err = u.writeCSVFile(newRecords)
	if err != nil {
		return err
	}

	return nil
}

func (u userRepository) Update(ctx context.Context, user *domain.User) error {
	records, err := u.readCSVFile()
	if err != nil {
		return err
	}

	var updated bool
	var updatedRecords [][]string

	updatedRecords = append(updatedRecords, records[0])

	for _, record := range records[1:] {
		existingUser, err := helper.MapRecordToUser(record)
		if err != nil {
			continue
		}

		if existingUser.ID == user.ID {
			if user.Name != "" {
				existingUser.Name = user.Name
			}
			if user.Currency != "" {
				existingUser.Currency = user.Currency
			}
			if user.Scale != 0 {
				existingUser.Scale = user.Scale
			}
			if user.Balance.String() != "" {
				existingUser.Balance = user.Balance
			}
			if user.UpdatedBy != "" {
				existingUser.UpdatedBy = user.UpdatedBy
			}
			existingUser.UpdatedDate = time.Now()
			existingUser.Version++

			updatedRecords = append(updatedRecords, []string{
				existingUser.ID,
				existingUser.Name,
				existingUser.Currency,
				fmt.Sprintf("%d", existingUser.Scale),
				existingUser.Balance.String(),
				existingUser.CreatedBy,
				existingUser.CreatedDate.Format(time.RFC3339),
				existingUser.UpdatedBy,
				existingUser.UpdatedDate.Format(time.RFC3339),
				fmt.Sprintf("%d", existingUser.Version),
				fmt.Sprintf("%d", existingUser.IsDeleted),
			})
			updated = true
		} else {
			updatedRecords = append(updatedRecords, record)
		}
	}

	if !updated {
		return fmt.Errorf("user with ID %s not found", user.ID)
	}

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

	updatedRecords = append(updatedRecords, records[0])

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

	if len(updatedRecords) == len(records) {
		return fmt.Errorf("user with ID %s not found", user.ID)
	}

	if err := u.writeCSVFile(updatedRecords); err != nil {
		return err
	}

	return nil
}
