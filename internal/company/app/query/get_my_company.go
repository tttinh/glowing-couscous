package query

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/domain"
	er "github.com/tttinh/glowing-couscous/pkg/errors"

	ce "github.com/tttinh/glowing-couscous/common/errors"
)

type GetMyCompanyRepo interface {
	GetEmployeeById(ctx context.Context, id string) (*domain.Employee, error)
	GetCompanyById(ctx context.Context, id string) (*domain.Company, error)
}

type GetMyCompanyHandler struct {
	repo GetMyCompanyRepo
}

func NewGetMyCompanyHandler(repo GetMyCompanyRepo) GetMyCompanyHandler {
	if repo == nil {
		panic("nil repo")
	}

	return GetMyCompanyHandler{repo}
}

func (h GetMyCompanyHandler) Handle(ctx context.Context, employeeId string) (c *domain.Company, err error) {
	e, err := h.repo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return nil, err
	}

	if e == nil {
		return nil, ce.New(ce.ErrBadRequest, er.ErrUserNotFound.Error())
	}

	if e.CompanyId == "" {
		return nil, ce.New(ce.ErrBadRequest, er.ErrBankIdInvalid.Error())
	}

	c, err = h.repo.GetCompanyById(ctx, e.CompanyId)
	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, ce.New(ce.ErrBadRequest, er.ErrBankNotFound.Error())
	}

	return c, nil
}
