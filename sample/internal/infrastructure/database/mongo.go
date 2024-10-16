package database

import (
	"context"
	"fmt"
	"log"

	// ここのパッケージ名変更していいのでは？
	"github.com/sin392/db-media-sample/sample/internal/adapter/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoHandler struct {
	db     *mongo.Database
	client *mongo.Client
}

func NewMongoHandler(c *config) (repository.NoSQL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.ctxTimeout)
	defer cancel()

	uri := fmt.Sprintf(
		"%s://%s:%s",
		c.driver,
		c.host,
		c.port,
	)
	// uri := fmt.Sprintf(
	// 	"%s://%s:%s@%s:%s/?replicaSet=replicaset",
	// 	c.Driver,
	// 	c.User,
	// 	c.Password,
	// 	c.Host,
	// 	c.Port,
	// )

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &mongoHandler{
		db:     client.Database(c.database),
		client: client,
	}, nil
}

func (mgo mongoHandler) Store(ctx context.Context, collection string, data interface{}) error {
	if _, err := mgo.db.Collection(collection).InsertOne(ctx, data); err != nil {
		return mgo.classifyError(err)
	}

	return nil
}

func (mgo mongoHandler) Update(ctx context.Context, collection string, query interface{}, update interface{}) error {
	if _, err := mgo.db.Collection(collection).UpdateOne(ctx, query, update); err != nil {
		return err
	}

	return nil
}

func (mgo mongoHandler) FindAll(ctx context.Context, collection string, query interface{}, result interface{}) error {
	cur, err := mgo.db.Collection(collection).Find(ctx, query)
	if err != nil {
		return mgo.classifyError(err)
	}

	defer cur.Close(ctx)
	if err = cur.All(ctx, result); err != nil {
		return mgo.classifyError(err)
	}

	if err := cur.Err(); err != nil {
		return mgo.classifyError(err)
	}

	return nil
}

func (mgo mongoHandler) FindOne(
	ctx context.Context,
	collection string,
	query interface{},
	projection interface{},
	result interface{},
) error {
	err := mgo.db.Collection(collection).
		FindOne(
			ctx,
			query,
			options.FindOne().SetProjection(projection),
		).Decode(result)
	if err != nil {
		return mgo.classifyError(err)
	}

	return nil
}

func (mgo *mongoHandler) StartSession() (repository.NoSQLSession, error) {
	session, err := mgo.client.StartSession()
	if err != nil {
		log.Fatal(err)
	}

	return newMongoHandlerSession(session), nil
}

func (m *mongoHandler) classifyError(err error) error {
	switch err {
	case mongo.ErrNoDocuments:
		return repository.NewDatabaseError(repository.NotFoundError, err.Error())
	case mongo.ErrNilDocument:
		return repository.NewDatabaseError(repository.InvalidParameterError, err.Error())
	case mongo.ErrClientDisconnected:
		return repository.NewDatabaseError(repository.ConnectionError, err.Error())
	case mongo.ErrUnacknowledgedWrite:
		return repository.NewDatabaseError(repository.ConflictError, err.Error())
	default:
		return repository.NewDatabaseError(repository.UnknownError, err.Error())
	}
}

type mongoDBSession struct {
	session mongo.Session
}

func newMongoHandlerSession(session mongo.Session) *mongoDBSession {
	return &mongoDBSession{session: session}
}

func (m *mongoDBSession) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		err := fn(sessCtx)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	_, err := m.session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoDBSession) EndSession(ctx context.Context) {
	m.session.EndSession(ctx)
}
