package code

import (
  "fmt"
  "mock"
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

  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  utils.Assert(mock.TestToken == encrypted, func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: mock.TestToken, Got: encrypted}))
    t.Fail()
  })
}

func TestCoder_Decrypt(t *testing.T) {
  expected := map[string]interface{}{
    "id": 1,
  }
  decrypted, err := coder.Decrypt(mock.TestToken)

  utils.AssertIsNil(err, func(exp string) {
    t.Log(exp)
    t.Fail()
  })
  utils.Assert(fmt.Sprintf("%v", decrypted["id"]) == fmt.Sprintf("%v", expected["id"]), func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: expected["user_id"], Got: decrypted["user_id"]}))
    t.Fail()
  })
  utils.Assert(decrypted["role"] == expected["role"], func() {
    t.Log(
      utils.GetExpectationString(
        utils.Expectation{Expected: expected["role"], Got: decrypted["role"]}))
    t.Fail()
  })
}

func TestCoder_DecryptErrorEmpty(t *testing.T) {
  _, err := coder.Decrypt("")

  utils.Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}

func TestCoder_DecryptErrorIncorrectFormat(t *testing.T) {
  _, err := coder.Decrypt("etJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ8.eyJyb2xlIjoiYWRtaW4iLCJ1c2VyX2lkIjoxfQ.")

  utils.Assert(err != nil, func() {
    t.Log("should be error")
    t.Fail()
  })
}
