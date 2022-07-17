package query

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/domain"
	er "github.com/tttinh/glowing-couscous/pkg/errors"

	ce "github.com/tttinh/glowing-couscous/common/errors"
)

type ListGroupRepo interface {
	GetEmployeeById(ctx context.Context, id string) (*domain.Employee, error)
	GetDepartmentById(ctx context.Context, id string) (*domain.Department, error)
	ListGroup(ctx context.Context, departmentId string) ([]*domain.Group, error)
}

type ListGroupHandler struct {
	repo ListGroupRepo
}

func NewListGroupHandler(repo ListGroupRepo) ListGroupHandler {
	if repo == nil {
		panic("nil repo")
	}

	return ListGroupHandler{repo}
}

func (h ListGroupHandler) Handle(ctx context.Context, employeeId, departmentId string) (results []*domain.Group, err error) {
	// User must be existed
	e, err := h.repo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return nil, err
	}

	if e == nil {
		return nil, ce.New(ce.ErrBadRequest, er.ErrUserNotFound.Error())
	}

	// Department must be existed
	d, err := h.repo.GetDepartmentById(ctx, departmentId)
	if err != nil {
		return nil, err
	}

	if d == nil {
		return nil, ce.New(ce.ErrBadRequest, er.ErrBankBranchIdInvalid.Error())
	}

	idCEO := e.Role == domain.EmployeeRole_CEO && e.CompanyId == d.CompanyId
	isHead := e.Role == domain.EmployeeRole_HEAD && e.DepartmentId == d.Id

	if !idCEO && !isHead {
		return nil, ce.New(ce.ErrBadRequest, er.ErrBankerUnauthorized.Error())
	}

	results, err = h.repo.ListGroup(ctx, departmentId)
	if err != nil {
		return nil, err
	}

	return results, nil
}
