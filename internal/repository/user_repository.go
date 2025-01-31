package repository

import (
	"PAPER-WALLET-SERVICE-CORE/config"
	"PAPER-WALLET-SERVICE-CORE/helper"
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
	"context"
	"encoding/csv"
	"fmt"
	"os"
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
	file, err := os.Open(u.filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read CSV: %v", err)
	}

	var users []domain.User

	for _, record := range records[1:] {
		user, err := helper.MapRecordToUser(record)
		if err != nil {
			continue
		}

		if helper.MatchFilter(user, filter) {
			users = append(users, user)
		}
	}

	return users, nil
}

func (u userRepository) FindOne(ctx context.Context, filter map[string]interface{}) (*domain.User, error) {
	file, err := os.Open(u.filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read CSV: %v", err)
	}

	for _, record := range records[1:] {
		user, err := helper.MapRecordToUser(record)
		if err != nil {
			continue
		}

		if helper.MatchFilter(user, filter) {
			return &user, nil
		}
	}

	return nil, nil
}

func (u userRepository) Upsert(ctx context.Context, user *domain.User) error {
	file, err := os.OpenFile(u.filePath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("unable to read CSV: %v", err)
	}

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

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	err = writer.WriteAll(updatedRecords)
	if err != nil {
		return fmt.Errorf("unable to write to CSV: %v", err)
	}

	return nil
}

func (u userRepository) SoftDelete(ctx context.Context, user *domain.User) error {
	file, err := os.OpenFile(u.filePath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("unable to read CSV: %v", err)
	}

	var updated bool
	var updatedRecords [][]string

	for _, record := range records[1:] {
		existingUser, err := helper.MapRecordToUser(record)
		if err != nil {
			continue
		}

		if existingUser.ID == user.ID {
			existingUser.IsDeleted = 1
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
				"1",
			})
			updated = true
		} else {
			updatedRecords = append(updatedRecords, record)
		}
	}

	if !updated {
		return fmt.Errorf("user with ID %s not found", user.ID)
	}

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	err = writer.WriteAll(updatedRecords)
	if err != nil {
		return fmt.Errorf("unable to write to CSV: %v", err)
	}

	return nil
}

func (u userRepository) HardDelete(ctx context.Context, user *domain.User) error {
	file, err := os.OpenFile(u.filePath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("unable to read CSV: %v", err)
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

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	err = writer.WriteAll(updatedRecords)
	if err != nil {
		return fmt.Errorf("unable to write to CSV: %v", err)
	}

	return nil
}
