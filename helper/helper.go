package helper

import (
	"PAPER-WALLET-SERVICE-CORE/internal/domain"
	"PAPER-WALLET-SERVICE-CORE/internal/dto"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"net/http"
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

	updatedDate, err := time.Parse(time.RFC3339, record[8])
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
		ID:          record[0],
		Name:        record[1],
		Currency:    record[2],
		Scale:       scale,
		Balance:     balance,
		CreatedDate: createdDate,
		CreatedBy:   record[5],
		UpdatedDate: updatedDate,
		UpdatedBy:   record[7],
		Version:     version,
		IsDeleted:   isDeleted,
	}, nil
}

func uint32FromString(s string) (uint32, error) {
	var val uint32
	_, err := fmt.Sscanf(s, "%d", &val)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func intFromString(s string) (int, error) {
	var val int
	_, err := fmt.Sscanf(s, "%d", &val)
	return val, err
}

func GetMandatoryRequest(e echo.Context) dto.MandatoryRequest {
	var header = e.Request().Header

	return dto.MandatoryRequest{
		ChannelID:     getHeaderOrDefault(header, "X-Channel-Id", ""),
		RequestID:     getHeaderOrDefault(header, "X-Request-Id", ""),
		ServiceID:     getHeaderOrDefault(header, "X-Service-Id", ""),
		Username:      getHeaderOrDefault(header, "X-Username", ""),
		Language:      getHeaderOrDefault(header, "Accept-Language", "en"),
		UserAgent:     getHeaderOrDefault(header, "User-Agent", ""),
		Authorization: getHeaderOrDefault(header, "Authorization", ""),
		AppVersion:    getHeaderOrDefault(header, "X-App-Version", ""),
	}
}

func getHeaderOrDefault(header http.Header, key, defaultValue string) string {
	var val = header.Get(key)

	if len(val) == 0 {
		return defaultValue
	}

	return val
}
