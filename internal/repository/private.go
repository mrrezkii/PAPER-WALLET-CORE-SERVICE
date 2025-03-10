package repository

import (
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
	"PAPER-WALLET-SERVICE-CORE/internal/handler"
	"PAPER-WALLET-SERVICE-CORE/shared"
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func (u userRepository) readCSVFile() ([][]string, error) {
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

	return records, nil
}

func (u userRepository) writeCSVFile(records [][]string) error {
	file, err := os.OpenFile(u.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(records)
	if err != nil {
		return fmt.Errorf("unable to write to CSV: %v", err)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("error flushing CSV writer: %v", err)
	}

	return nil
}

func (u userRepository) findUsersByFilter(records [][]string, filter map[string]interface{}) ([]domain.User, error) {
	var users []domain.User
	for _, record := range records[1:] {
		user, err := shared.MapRecordToUser(record)
		if err != nil {
			continue
		}

		if shared.MatchFilter(user, filter) {
			users = append(users, user)
		}
	}
	return users, nil
}

func (u userRepository) softDeleteUserRecord(ctx context.Context, records [][]string, user *domain.User) ([][]string, error) {
	mandatoryRequest := handler.MandatoryRequest(ctx)
	var updated bool
	var updatedRecords [][]string

	updatedRecords = append(updatedRecords, records[0])

	for _, record := range records[1:] {
		existingUser, err := shared.MapRecordToUser(record)
		if err != nil {
			continue
		}

		if existingUser.ID == user.ID {
			existingUser.IsDeleted = 1
			existingUser.UpdatedBy = mandatoryRequest.Username
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
				"1", // IsDeleted
			})
			updated = true
		} else {
			updatedRecords = append(updatedRecords, record)
		}
	}

	if !updated {
		return nil, fmt.Errorf("user with ID %s not found", user.ID)
	}

	return updatedRecords, nil
}
