package command

import (
	"context"

	ce "github.com/tttinh/glowing-couscous/common/errors"
	"github.com/tttinh/glowing-couscous/internal/company/domain"
	er "github.com/tttinh/glowing-couscous/pkg/errors"
)

type UpdateCompany struct {
	OperatorId string
	Code       string
	Name       string
	Slogan     string
	Logo       string
}

type UpdateCompanyRepo interface {
	GetEmployeeById(ctx context.Context, id string) (*domain.Employee, error)
	UpdateCompany(ctx context.Context, c *domain.Company) (bool, error)
}

type UpdateCompanyHandler struct {
	repo UpdateCompanyRepo
}

func NewUpdateCompanyHandler(repo UpdateCompanyRepo) UpdateCompanyHandler {
	if repo == nil {
		panic("nil repo")
	}

	return UpdateCompanyHandler{repo}
}

func (h UpdateCompanyHandler) Handle(ctx context.Context, cmd UpdateCompany) error {
	e, err := h.repo.GetEmployeeById(ctx, cmd.OperatorId)
	if err != nil {
		return err
	}

	if e == nil {
		return ce.New(ce.ErrBadRequest, er.ErrUserNotFound.Error())
	}

	if e.CompanyId == "" {
		return ce.New(ce.ErrBadRequest, er.ErrBankIdInvalid.Error())
	}

	isOwner := e.Role == domain.EmployeeRole_CEO
	if !isOwner {
		return ce.New(ce.ErrBadRequest, er.ErrBankerUnauthorized.Error())
	}

	c, err := domain.NewCompanyForUpdate(
		e.CompanyId,
		cmd.Code,
		cmd.Name,
		cmd.Slogan,
		cmd.Logo,
	)
	if err != nil {
		return err
	}

	ok, err := h.repo.UpdateCompany(ctx, c)
	if err != nil {
		return err
	}

	if !ok {
		return ce.New(ce.ErrBadRequest, er.ErrBankerUnauthorized.Error())
	}

	return nil
}
