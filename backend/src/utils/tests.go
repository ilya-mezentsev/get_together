package utils

import (
  "fmt"
  "os"
)

type Expectation struct {
  Expected, Got interface{}
}

func SkipInShortMode() {
  if os.Getenv("SHORT_MODE") == "1" {
    fmt.Println("skipping DB tests in short mode")
    os.Exit(0)
  }
}

func Assert(condition bool, onFalseFn func()) {
  if !condition {
    onFalseFn()
  }
}

func AssertIsNil(v interface{}, onFalseFn func(string)) {
  if v != nil {
    onFalseFn(
      GetExpectationString(
        Expectation{Got: v}))
  }
}

func AssertErrorsEqual(expectedErr, actualErr error, onFalseFn func(string)) {
  Assert(
    expectedErr != nil && actualErr != nil && expectedErr.Error() == actualErr.Error(),
    func() {
      onFalseFn(GetExpectationString(Expectation{Expected: expectedErr, Got: actualErr}))
    },
  )
}

func GetExpectationString(e Expectation) string {
  return fmt.Sprintf("expected: %v, got: %v\n", e.Expected, e.Got)
}
