package repository

import (
	"PAPER-WALLET-SERVICE-CORE/helper"
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
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

func (u userRepository) softDeleteUserRecord(records [][]string, user *domain.User) ([][]string, error) {
	var updated bool
	var updatedRecords [][]string

	updatedRecords = append(updatedRecords, records[0]) // Keep header row

	for _, record := range records[1:] {
		existingUser, err := helper.MapRecordToUser(record)
		if err != nil {
			continue
		}

		if existingUser.ID == user.ID {
			existingUser.IsDeleted = 1
			existingUser.UpdatedBy = "system"
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
