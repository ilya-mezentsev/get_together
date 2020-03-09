package api

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "io/ioutil"
  "models"
  "net/http"
  "plugins/logger"
)

const (
  StatusOk = "ok"
  StatusError = "error"
)

var r *mux.Router

func init() {
  r = mux.NewRouter()
}

func GetRouter() *mux.Router {
  return r
}

func DecodeRequestBody(r *http.Request, target interface{}) {
  requestBody, err := ioutil.ReadAll(r.Body)
  if err != nil {
    logger.ErrorF("Error while reading request body: %v", err)
    panic(ReadRequestBodyError)
  }

  err = json.Unmarshal(requestBody, target)
  if err != nil {
    logger.WithFields(logger.Fields{
      MessageTemplate: "Error while decoding request body: %v",
      Args: []interface{}{err},
    }, logger.Error)

    panic(CannotDecodeRequestBody)
  }
}

func SendDefaultResponse(w http.ResponseWriter) {
  EncodeAndSendResponse(w, nil)
}

func EncodeAndSendResponse(w http.ResponseWriter, v interface{}) {
  output, _ := json.Marshal(models.SuccessResponse{
    Status: StatusOk,
    Data: v,
  })

  w.Header().Set("content-type", "application/json")
  if _, err := w.Write(output); err != nil {
    logger.WithFields(logger.Fields{
      MessageTemplate: "Error while trying to write response: %v",
      Args: []interface{}{err},
    }, logger.Error)

    panic(CannotWriteResponse)
  }
}

func SendErrorIfPanicked(w http.ResponseWriter) {
  if err := recover(); err != nil {
    logger.WarningF("Panicked: %v", err)

    output, _ := json.Marshal(models.ErrorResponse{
      Status: StatusError,
      ErrorDetail: err.(error).Error(),
    })

    w.Header().Set("content-type", "application/json")
    if _, err = w.Write(output); err != nil {
      logger.WithFields(logger.Fields{
        MessageTemplate: "Error while trying to write response: %v",
        Args: []interface{}{err},
      }, logger.Error)

      http.Error(w, "internal server error", http.StatusInternalServerError)
    }
  }
}
