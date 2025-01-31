package helper

import (
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
	"PAPER-WALLET-SERVICE-CORE/internal/dto"
	"fmt"
	"github.com/shopspring/decimal"
	"reflect"
	"time"
)

func MatchFilter(obj interface{}, filter map[string]interface{}) bool {
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for key, value := range filter {
		field := val.FieldByName(key)

		if !field.IsValid() {
			continue
		}

		if !field.IsZero() && field.Interface() != value {
			return false
		}
	}
	return true
}

func MapRecordToUser(record []string) (domain.User, error) {
	createdDate, err := time.Parse(time.RFC3339, record[6])
	if err != nil {
		return domain.User{}, fmt.Errorf("invalid createdDate: %v", err)
	}

	updatedDate, err := time.Parse(time.RFC3339, record[7])
	if err != nil {
		return domain.User{}, fmt.Errorf("invalid updatedDate: %v", err)
	}

	balance, err := decimal.NewFromString(record[4])
	if err != nil {
		return domain.User{}, fmt.Errorf("invalid balance: %v", err)
	}

	scale, err := uint32FromString(record[3])
	if err != nil {
		return domain.User{}, fmt.Errorf("invalid scale: %v", err)
	}

	version, err := intFromString(record[9])
	if err != nil {
		return domain.User{}, fmt.Errorf("invalid version: %v", err)
	}

	isDeleted, err := intFromString(record[10])
	if err != nil {
		return domain.User{}, fmt.Errorf("invalid isDeleted: %v", err)
	}

	return domain.User{
		Name:     record[1],
		Currency: record[2],
		Scale:    scale,
		Balance:  balance,
		BaseTableFields: dto.BaseTableFields{
			ID:          record[0],
			CreatedDate: createdDate,
			CreatedBy:   record[5],
			UpdatedDate: updatedDate,
			UpdatedBy:   record[8],
			Version:     version,
			IsDeleted:   isDeleted,
		},
	}, nil
}

func uint32FromString(s string) (uint32, error) {
	val, err := fmt.Sscanf(s, "%d", new(uint32))
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

func intFromString(s string) (int, error) {
	var val int
	_, err := fmt.Sscanf(s, "%d", &val)
	return val, err
}
