package command

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/domain"
	er "github.com/tttinh/glowing-couscous/pkg/errors"

	ce "github.com/tttinh/glowing-couscous/common/errors"
)

type AddCompany struct {
	Code   string
	Name   string
	Slogan string
	Logo   string
}

type AddCompanyHandler struct {
	repo domain.CompanyRepository
}

func NewAddCompanyHandler(repo domain.CompanyRepository) AddCompanyHandler {
	if repo == nil {
		panic("nil repo")
	}

	return AddCompanyHandler{repo}
}

func (h AddCompanyHandler) Handle(ctx context.Context, cmd AddCompany) (string, error) {
	c, err := h.repo.GetCompanyByCode(ctx, cmd.Code)
	if err != nil {
		return "", err
	}

	if c != nil {
		return "", ce.New(ce.ErrBadRequest, er.ErrBankAlreadyExisted.Error())
	}

	company, err := domain.NewCompanyForCreate(
		cmd.Code,
		cmd.Name,
		cmd.Slogan,
		cmd.Logo,
	)
	if err != nil {
		return "", err
	}

	company, err = h.repo.CreateCompany(ctx, company)
	if err != nil {
		return "", err
	}

	return company.Id, nil
}
