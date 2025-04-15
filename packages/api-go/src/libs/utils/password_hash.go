package utils

import (
	"github.com/alexedwards/argon2id"
)

// Hashes a plaintext password, returning the hash as a string
func HashPassword(password string) (string, error) {
	argon2Config := &argon2id.Params{
		Memory:      uint32(19456),
		Iterations:  uint32(2),
		Parallelism: uint8(1),
		SaltLength:  uint32(16),
		KeyLength:   uint32(32),
	}

	return argon2id.CreateHash(password, argon2Config)
}

// Compares a plain text password to a hashed password, returning true if they match, or false otherwise
func VerifyPasswordHash(password string, reference_hash string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, reference_hash)
	if err != nil {
		return false
	}

	return match
}
