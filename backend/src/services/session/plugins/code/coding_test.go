package code

import (
	"fmt"
	"mock/services"
	"os"
	"testing"
	"utils"
)

var coder Coder

func init() {
	coderKey := os.Getenv("CODER_KEY")
	if coderKey == "" {
		fmt.Println("CODER_KEY env var is not set")
		os.Exit(1)
	}

	coder = NewCoder(coderKey)
}

func TestCoder_Encrypt(t *testing.T) {
	encrypted, err := coder.Encrypt(map[string]interface{}{
		"id": 1,
	})

	utils.AssertNil(err, t)
	utils.AssertEqual(services.TestToken, encrypted, t)
}

func TestCoder_Decrypt(t *testing.T) {
	expected := map[string]interface{}{
		"id": 1,
	}
	decrypted, err := coder.Decrypt(services.TestToken)

	utils.AssertNil(err, t)
	utils.AssertEqual(fmt.Sprintf("%v", decrypted["id"]), fmt.Sprintf("%v", expected["id"]), t)
	utils.AssertEqual(decrypted["role"], expected["role"], t)
}

func TestCoder_DecryptErrorEmpty(t *testing.T) {
	_, err := coder.Decrypt("")

	utils.AssertNotNil(err, t)
}

func TestCoder_DecryptErrorIncorrectFormat(t *testing.T) {
	_, err := coder.Decrypt("etJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ8.eyJyb2xlIjoiYWRtaW4iLCJ1c2VyX2lkIjoxfQ.")

	utils.AssertNotNil(err, t)
}
