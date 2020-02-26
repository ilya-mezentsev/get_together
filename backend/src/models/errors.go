package models

type (
  ErrorWrapper struct {
    originalError, externalError error
  }
)

func NewErrorWrapper(originalError, externalError error) ErrorWrapper {
  return ErrorWrapper{
    originalError: originalError,
    externalError: externalError,
  }
}

func (e ErrorWrapper) OriginalError() error {
  return e.originalError
}

func (e ErrorWrapper) ExternalError() error {
  return e.externalError
}
