package utils

import "github.com/lib/pq"

// IsConstraintViolation checks if a database error matches a specific constraint name
// Parameters:
//   - err: The error to check
//   - constraint: The name of the constraint to match against
//
// Returns:
//   - bool: true if the error is a postgres constraint violation matching the given constraint name,
//     false otherwise
func IsConstraintViolation(err error, constraint string) bool {
	dbError, ok := err.(*pq.Error)
	if !ok {
		return false
	}

	if dbError.Constraint == constraint {
		return true
	}

	return false
}
