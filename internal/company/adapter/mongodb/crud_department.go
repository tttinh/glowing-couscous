package mongodb

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Repository) CreateDepartment(ctx context.Context, d *domain.Department) (*domain.Department, error) {
	d.Id = ""
	m, err := DepartmentFromDomain(d)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.CreateDepartment")
	}

	m.SetCreatedAt()
	rs, err := r.departments.InsertOne(ctx, m)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.CreateDepartment: collection.InsertOne")
	}

	m.Id = rs.InsertedID.(primitive.ObjectID)
	department, err := DepartmentToDomain(m)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.CreateDepartment")
	}

	return department, nil
}

func (r Repository) GetDepartmentById(ctx context.Context, id string) (*domain.Department, error) {
	departmentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var d *Department
	err = r.departments.FindOne(ctx, bson.M{"_id": departmentId}).Decode(&d)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "Repository.GetDepartmentById: FindOne")
	}

	department, err := DepartmentToDomain(d)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.GetDepartmentById: DepartmentToDomain")
	}

	return department, nil
}

func (r Repository) ListDepartment(ctx context.Context, companyId string) ([]*domain.Department, error) {
	cursor, err := r.departments.Find(ctx, bson.M{"company_id": companyId})
	if err != nil {
		return nil, errors.Wrap(err, "Repository.ListDepartment: Find")
	}

	var departmentModels []*Department
	if err = cursor.All(ctx, &departmentModels); err != nil {
		return nil, errors.Wrap(err, "Repository.ListDepartment: cursor.All")
	}

	var results []*domain.Department
	for _, m := range departmentModels {
		d, err := DepartmentToDomain(m)
		if err != nil {
			return nil, errors.Wrap(err, "Repository.ListDepartment: DepartmentToDomain")
		}
		results = append(results, d)
	}

	return results, nil
}
