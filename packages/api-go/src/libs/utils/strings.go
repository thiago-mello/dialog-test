package utils

// StringPointer takes a string value and returns a pointer to that string.
// This is useful when you need to pass a string pointer but only have a string value.
func StringPointer(value string) *string {
	return &value
}
