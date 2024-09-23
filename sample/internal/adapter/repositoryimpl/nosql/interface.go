package nosql

import "context"

type (
	NoSQL interface {
		Store(ctx context.Context, collection string, data interface{}) error
		Update(ctx context.Context, collection string, query interface{}, update interface{}) error
		FindAll(ctx context.Context, collection string, query interface{}, result interface{}) error
		FindOne(ctx context.Context, collection string, query interface{}, projection interface{}, result interface{}) error
		StartSession() (Session, error)
	}
	Session interface {
		WithTransaction(context.Context, func(context.Context) error) error
		EndSession(context.Context)
	}
)
