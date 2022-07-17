package mongodb

import (
	"context"
	"time"

	"github.com/tttinh/glowing-couscous/internal/company/adapter/client"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

const (
	companyCollectionName    string = "companies"
	departmentCollectionName string = "departments"
	groupCollectionName      string = "groups"
	employeeCollectionName   string = "employees"
)

type Repository struct {
	userClient  *client.User
	db          *mongo.Database
	companies   *mongo.Collection
	departments *mongo.Collection
	groups      *mongo.Collection
	employees   *mongo.Collection
}

func NewRepository(userClient *client.User, db *mongo.Database) *Repository {
	if db == nil {
		panic("nil *mongo.Database")
	}

	return &Repository{
		userClient:  userClient,
		db:          db,
		companies:   db.Collection(companyCollectionName),
		departments: db.Collection(departmentCollectionName),
		groups:      db.Collection(groupCollectionName),
		employees:   db.Collection(employeeCollectionName),
	}
}

type TxnFunc func(sc mongo.SessionContext) (interface{}, error)

func (r Repository) WithTnx(callback TxnFunc) (interface{}, error) {
	var (
		ses mongo.Session
		ctx context.Context
		err error
	)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if ses, err = r.db.Client().StartSession(); err != nil {
		return nil, err
	}
	defer ses.EndSession(ctx)

	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	return ses.WithTransaction(ctx, callback, txnOpts)
}
