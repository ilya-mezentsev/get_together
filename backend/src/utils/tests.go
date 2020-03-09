package utils

import (
  "bytes"
  "fmt"
  "github.com/gorilla/mux"
  "io"
  "net/http"
  "net/http/httptest"
  "os"
  "testing"
)

type Expectation struct {
  Expected, Actual interface{}
}

type RequestData struct {
  Router *mux.Router
  Method, Endpoint, Data string
  Cookie *http.Cookie
}

func MakeRequest(rd RequestData) io.ReadCloser {
  srv := httptest.NewServer(rd.Router)
  defer srv.Close()

  req := getHttpRequest(rd.Method, fmt.Sprintf("%s/%s", srv.URL, rd.Endpoint), rd.Data)
  req.AddCookie(rd.Cookie)
  return doRequestAndGetBody(req)
}

func getHttpRequest(method, url, data string) *http.Request {
  req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
  if err != nil {
    panic(err)
  }

  req.Header.Set("Content-Type", "application/json; charset=utf-8")
  return req
}

func doRequestAndGetBody(req *http.Request) io.ReadCloser {
  resp, err := (&http.Client{}).Do(req)
  if err != nil {
    panic(err)
  }

  return resp.Body
}

func SkipInShortMode() {
  if os.Getenv("SHORT_MODE") == "1" {
    fmt.Println("skipping DB tests in short mode")
    os.Exit(0)
  }
}

func AssertTrue(actual bool, t *testing.T) {
  if !actual {
    logExpectationAndFail(true, actual, t)
  }
}

func AssertFalse(actual bool, t *testing.T) {
  if actual {
    logExpectationAndFail(false, actual, t)
  }
}

func AssertEqual(expected, actual interface{}, t *testing.T) {
  if expected != actual {
    logExpectationAndFail(expected, actual, t)
  }
}

func AssertNil(v interface{}, t *testing.T) {
  if v != nil {
    logExpectationAndFail(nil, v, t)
  }
}

func AssertNotNil(v interface{}, t *testing.T) {
  if v == nil {
    logExpectationAndFail("not nil", v, t)
  }
}

func AssertErrorsEqual(expectedErr, actualErr error, t *testing.T) {
  if expectedErr != actualErr {
    logExpectationAndFail(expectedErr, actualErr, t)
  }
}

func logExpectationAndFail(expected, actual interface{}, t *testing.T) {
  t.Log(
    GetExpectationString(
      Expectation{Expected: expected, Actual: actual}))
  t.Fail()
}

func GetExpectationString(e Expectation) string {
  return fmt.Sprintf("expected: %v, got: %v\n", e.Expected, e.Actual)
}
