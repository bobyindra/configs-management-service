package encryption_test

import (
	"testing"

	"github.com/bobyindra/configs-management-service/module/configuration/internal/encryption"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncryption_GeneratePassword(t *testing.T) {
	t.Parallel()

	t.Run("Generate Password Success", func(t *testing.T) {
		t.Parallel()

		// Given
		encryption := encryption.NewEncryption()
		password := "Test Password"

		// When
		encrypPass, err := encryption.GeneratePassword(password)

		// Then
		assert.NotEmpty(t, encrypPass, "should contain generated password")
		assert.Nil(t, err, "error should be empty")
	})

	t.Run("Generate Password Error", func(t *testing.T) {
		t.Parallel()

		// Given
		encryption := encryption.NewEncryption()
		password := "Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."

		// When
		encrypPass, err := encryption.GeneratePassword(password)

		// Then
		assert.Nil(t, encrypPass, "encrypted password should be empty")
		require.Error(t, err, "error should not be empty")
	})

}

func TestEncryption_ComparePassword(t *testing.T) {
	t.Run("Compare Match Password", func(t *testing.T) {
		t.Parallel()

		// Given
		encryption := encryption.NewEncryption()
		plainPassword := "Test Password"
		encryptedPassword := "$2a$10$lWySZfvaI35s0lawVhiTROYjbjOOU5ANZJ9YGMC2p0rowWuqJ4uUW"

		// When
		err := encryption.ComparePassword(encryptedPassword, plainPassword)

		// Then
		assert.Nil(t, err, "error should be empty")
	})

	t.Run("Compare Unmatch Password ", func(t *testing.T) {
		t.Parallel()

		// Given
		encryption := encryption.NewEncryption()
		plainPassword := "New Password"
		encryptedPassword := "$2a$10$lWySZfvaI35s0lawVhiTROYjbjOOU5ANZJ9YGMC2p0rowWuqJ4uUW"

		// When
		err := encryption.ComparePassword(encryptedPassword, plainPassword)

		// Then
		require.Error(t, err, "error should not be empty")
	})
}
