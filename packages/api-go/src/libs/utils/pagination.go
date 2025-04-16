package utils

import "github.com/google/uuid"

// CalculatePageSize returns the page size for pagination.
// If the input pageSize is 0, it returns a default value of 10.
// Otherwise, it returns the input pageSize value.
func CalculatePageSize(pageSize int32) int32 {
	if pageSize == 0 {
		return 10
	}
	return pageSize
}

// StringPointerToUuid converts a string pointer containing a UUID string to a UUID pointer.
// If the input string pointer is nil, returns nil.
// If the string cannot be parsed as a valid UUID, returns nil.
// Otherwise returns a pointer to the parsed UUID.
func StringPointerToUuid(value *string) *uuid.UUID {
	if value == nil {
		return nil
	}

	uuidValue, err := uuid.Parse(*value)
	if err != nil {
		return nil
	}
	return &uuidValue
}
