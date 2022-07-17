package domain

import (
	"context"

	ce "github.com/tttinh/glowing-couscous/common/errors"
	er "github.com/tttinh/glowing-couscous/pkg/errors"
)

type Department struct {
	Id            string
	CompanyId     string
	Name          string
	Note          string
	EmployeeCount int
}

type DepartmentRepository interface {
	CreateDepartment(ctx context.Context, d *Department) (*Department, error)
	GetDepartmentById(ctx context.Context, id string) (*Department, error)
	ListDepartment(ctx context.Context, companyId string) ([]*Department, error)
}

func NewDepartmentToCreate(
	companyId,
	name,
	note string,
) (*Department, error) {
	if name == "" {
		return nil, ce.New(ce.ErrBadRequest, er.ErrNameRequired.Error())
	}

	if companyId == "" {
		return nil, ce.New(ce.ErrBadRequest, er.ErrBankIdRequired.Error())
	}

	return &Department{
		CompanyId: companyId,
		Name:      name,
		Note:      note,
	}, nil
}
