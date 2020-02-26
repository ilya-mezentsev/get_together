package interfaces

type (
  ErrorWrapper interface {
    OriginalError() error
    ExternalError() error
  }
)
