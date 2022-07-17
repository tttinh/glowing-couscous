package mongodb

import (
	"context"

	"github.com/google/uuid"
	"github.com/tttinh/glowing-couscous/internal/company/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Repository) CreateEmployee(ctx context.Context, e *domain.Employee) error {
	_, err := uuid.Parse(e.Id)
	if err != nil {
		return errors.Wrap(err, "Repository.CreateEmployee: Invalid user id")
	}

	m, err := EmployeeFromDomain(e)
	if err != nil {
		return errors.Wrap(err, "Repository.CreateEmployee: EmployeeFromDomain")
	}

	m.SetCreatedAt()
	_, err = r.employees.InsertOne(ctx, m)
	if err != nil {
		return errors.Wrap(err, "Repository.CreateEmployee: collection.InsertOne")
	}

	return nil
}

func (r Repository) GetEmployeeById(ctx context.Context, id string) (*domain.Employee, error) {
	var e *Employee
	err := r.employees.FindOne(ctx, bson.M{"_id": id}).Decode(&e)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "Repository.GetEmployeeById: FindOne")
	}

	user, err := r.userClient.GetUserProfile(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.GetEmployeeById: GetUserProfile")
	}

	if user.Id != e.Id {
		return nil, errors.Errorf("Repository.GetEmployeeById: inconsistent id between company and user")
	}

	user.Role = domain.EmployeeRole(e.Role)
	user.CompanyId = e.CompanyId
	user.DepartmentId = e.DepartmentId
	user.GroupId = e.GroupId

	return user, nil
}
