package command

import (
	"context"

	ce "github.com/tttinh/glowing-couscous/common/errors"
	"github.com/tttinh/glowing-couscous/internal/company/domain"
	er "github.com/tttinh/glowing-couscous/pkg/errors"
)

type AddCEO struct {
	EmployeeId string
	CompanyId  string
}

type AddCEORepo interface {
	GetCompanyById(ctx context.Context, id string) (*domain.Company, error)
	CreateEmployee(ctx context.Context, u *domain.Employee) error
}

type AddCEOHandler struct {
	repo AddCEORepo
}

func NewAddCEOHandler(repo AddCEORepo) AddCEOHandler {
	if repo == nil {
		panic("nil repo")
	}

	return AddCEOHandler{repo}
}

func (h AddCEOHandler) Handle(ctx context.Context, cmd AddCEO) error {
	ceo, err := domain.NewCEOToCreate(cmd.EmployeeId, cmd.CompanyId)
	if err != nil {
		return err
	}

	// Make sure the company is existed.
	c, err := h.repo.GetCompanyById(ctx, ceo.CompanyId)
	if err != nil {
		return err
	}

	if c == nil {
		return ce.New(ce.ErrBadRequest, er.ErrBankIdInvalid.Error())
	}

	err = h.repo.CreateEmployee(ctx, ceo)
	if err != nil {
		return err
	}

	return nil
}
