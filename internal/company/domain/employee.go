package domain

import (
	"context"
	"time"

	ce "github.com/tttinh/glowing-couscous/common/errors"
	er "github.com/tttinh/glowing-couscous/pkg/errors"
)

type EmployeeRole string

const (
	EmployeeRole_INVALID EmployeeRole = "invalid"
	EmployeeRole_CEO     EmployeeRole = "ceo"
	EmployeeRole_HEAD    EmployeeRole = "head"
	EmployeeRole_LEAD    EmployeeRole = "lead"
	EmployeeRole_STAFF   EmployeeRole = "staff"
)

type Employee struct {
	Id           string
	CompanyId    string
	DepartmentId string
	GroupId      string
	Name         string
	Email        string
	Phone        string
	Avatar       string
	LastLogin    *time.Time
	Role         EmployeeRole
}

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, e *Employee) error
	GetEmployeeById(ctx context.Context, id string) (*Employee, error)
}

func NewCEOToCreate(employeeId, companyId string) (*Employee, error) {
	if employeeId == "" {
		return nil, ce.New(ce.ErrBadRequest, er.ErrBankerIdRequired.Error())
	}

	if companyId == "" {
		return nil, ce.New(ce.ErrBadRequest, er.ErrBankIdRequired.Error())
	}

	return &Employee{
		Id:        employeeId,
		CompanyId: companyId,
		Role:      EmployeeRole_CEO,
	}, nil
}
