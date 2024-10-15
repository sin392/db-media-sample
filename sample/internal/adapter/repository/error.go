package repository

import (
	"errors"

	appErrors "github.com/sin392/db-media-sample/sample/internal/errors"
)

type DatabaseError struct {
	ErrType ErrorType
	err     error
}

type ErrorType int

const (
	// 入力パラメータの不正.
	InvalidParameterError ErrorType = iota
	// データ不整合.
	DataIntegrityError
	// データが見つからない.
	NotFoundError
	// データが既に存在する.
	ConflictError
	// 接続に失敗.
	ConnectionError
	// その他.
	UnknownError
)

func (e *DatabaseError) Error() string {
	if e == nil {
		return ""
	}
	return e.err.Error()
}

func (e *DatabaseError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

func (e *DatabaseError) GetType() ErrorType {
	if e == nil {
		return UnknownError
	}
	return e.ErrType
}

func NewDatabaseError(errType ErrorType, message string) *DatabaseError {
	return &DatabaseError{ErrType: errType, err: errors.New(message)}
}

func classifyError(err error) error {
	var dbErr *DatabaseError
	if errors.As(err, &dbErr) {
		switch dbErr.GetType() {
		case NotFoundError:
			return appErrors.NewApplicationError(appErrors.NotFoundError, err.Error())
		case InvalidParameterError:
			// DBのパラメータ不正はユーザ側で復旧できるエラーじゃないのでInternalErrorを返す
			return appErrors.NewApplicationError(appErrors.InternalError, err.Error())
		case ConflictError:
			return appErrors.NewApplicationError(appErrors.InternalError, err.Error())
		case ConnectionError:
			// データベースの接続失敗はユーザ側で復旧できるエラーじゃないのでInternalErrorを返す
			return appErrors.NewApplicationError(appErrors.InternalError, err.Error())
		default:
			return appErrors.NewApplicationError(appErrors.UnknownError, err.Error())
		}
	}
	return nil
}
