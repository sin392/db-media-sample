package errors

import "errors"

// ユースケース層/アダプター層で扱うエラー
type ApplicationError struct {
	ErrType ErrorType
	err     error
}

type ErrorType int

const (
	// 入力パラメータの不正
	InvalidParameterError ErrorType = iota
	// データが見つからない
	NotFoundError
	// データが既に存在する
	ConflictError
	// 内部エラー
	InternalError
	// その他
	UnknownError
)

func (e *ApplicationError) Error() string {
	if e == nil {
		return ""
	}
	return e.err.Error()
}

func (e *ApplicationError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

func (e *ApplicationError) GetType() ErrorType {
	if e == nil {
		return UnknownError
	}
	return e.ErrType
}

func NewApplicationError(errType ErrorType, message string) *ApplicationError {
	return &ApplicationError{ErrType: errType, err: errors.New(message)}
}
