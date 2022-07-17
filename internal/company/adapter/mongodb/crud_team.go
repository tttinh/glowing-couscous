package mongodb

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Repository) CreateGroup(ctx context.Context, g *domain.Group) (*domain.Group, error) {
	g.Id = ""
	m, err := GroupFromDomain(g)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.CreateGroup")
	}

	m.SetCreatedAt()
	rs, err := r.groups.InsertOne(ctx, m)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.CreateGroup: collection.InsertOne")
	}

	m.Id = rs.InsertedID.(primitive.ObjectID)
	group, err := GroupToDomain(m)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.CreateGroup")
	}

	return group, nil
}

func (r Repository) ListGroup(ctx context.Context, departmentId string) ([]*domain.Group, error) {
	cursor, err := r.groups.Find(ctx, bson.M{"department_id": departmentId})
	if err != nil {
		return nil, errors.Wrap(err, "Repository.ListGroup: Find")
	}

	var groupModels []*Group
	if err = cursor.All(ctx, &groupModels); err != nil {
		return nil, errors.Wrap(err, "Repository.ListGroup: cursor.All")
	}

	var results []*domain.Group
	for _, m := range groupModels {
		g, err := GroupToDomain(m)
		if err != nil {
			return nil, errors.Wrap(err, "Repository.ListGroup: GroupToDomain")
		}
		results = append(results, g)
	}

	return results, nil
}

func (r Repository) GetGroupById(ctx context.Context, id string) (*domain.Group, error) {
	groupId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var g *Group
	err = r.groups.FindOne(ctx, bson.M{"_id": groupId}).Decode(&g)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "Repository.GetGroupById: FindOne")
	}

	group, err := GroupToDomain(g)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.GetGroupById: GroupToDomain")
	}

	return group, nil
}
