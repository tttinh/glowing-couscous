package domain

import (
	"context"

	ce "github.com/tttinh/glowing-couscous/common/errors"
	er "github.com/tttinh/glowing-couscous/pkg/errors"
)

type Group struct {
	Id            string
	CompanyId     string
	DepartmentId  string
	Name          string
	Note          string
	EmployeeCount int
}

type GroupRepository interface {
	CreateGroup(ctx context.Context, g *Group) (*Group, error)
	GetGroupById(ctx context.Context, id string) (*Group, error)
}

func NewGroupToCreate(
	companyId,
	departmentId,
	name,
	note string,
) (*Group, error) {
	if name == "" {
		return nil, ce.New(ce.ErrBadRequest, er.ErrNameRequired.Error())
	}

	if departmentId == "" {
		return nil, ce.New(ce.ErrBadRequest, er.ErrBankBranchIdRequired.Error())
	}

	if companyId == "" {
		return nil, ce.New(ce.ErrBadRequest, er.ErrBankIdRequired.Error())
	}

	return &Group{
		CompanyId:    companyId,
		DepartmentId: departmentId,
		Name:         name,
		Note:         note,
	}, nil
}
