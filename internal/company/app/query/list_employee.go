package query

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/domain"
	er "github.com/tttinh/glowing-couscous/pkg/errors"

	ce "github.com/tttinh/glowing-couscous/common/errors"
)

type ListEmployeeRepo interface {
	GetEmployeeById(ctx context.Context, id string) (*domain.Employee, error)
	ListEmployeeByCompany(ctx context.Context, companyId string) ([]*domain.Employee, error)
}

type ListEmployeeHandler struct {
	repo ListEmployeeRepo
}

func NewListEmployeeHandler(repo ListEmployeeRepo) ListEmployeeHandler {
	if repo == nil {
		panic("nil repo")
	}

	return ListEmployeeHandler{repo}
}

func (h ListEmployeeHandler) Handle(ctx context.Context, employeeId string) (results []*domain.Employee, err error) {
	// User must be existed
	e, err := h.repo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return nil, err
	}

	if e == nil {
		return nil, ce.New(ce.ErrBadRequest, er.ErrUserNotFound.Error())
	}

	results, err = h.repo.ListEmployeeByCompany(ctx, e.CompanyId)
	if err != nil {
		return nil, err
	}

	return results, nil
}
