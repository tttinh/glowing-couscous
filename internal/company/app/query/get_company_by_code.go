package query

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/domain"
)

type GetCompanyByCodeHandler struct {
	repo domain.CompanyRepository
}

func NewGetBankByCodeHandler(repo domain.CompanyRepository) GetCompanyByCodeHandler {
	if repo == nil {
		panic("nil repo")
	}

	return GetCompanyByCodeHandler{repo}
}

func (h GetCompanyByCodeHandler) Handle(ctx context.Context, code string) (p *domain.Company, err error) {
	return h.repo.GetCompanyByCode(ctx, code)
}
