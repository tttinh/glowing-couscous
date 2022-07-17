package query

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/domain"
)

type ListCompanyRepo interface {
	ListCompany(ctx context.Context) ([]*domain.Company, error)
}

type ListCompanyHandler struct {
	repo ListCompanyRepo
}

func NewListCompanyHandler(repo ListCompanyRepo) ListCompanyHandler {
	if repo == nil {
		panic("nil repo")
	}

	return ListCompanyHandler{repo}
}

func (h ListCompanyHandler) Handle(ctx context.Context) (results []*domain.Company, err error) {
	results, err = h.repo.ListCompany(ctx)
	if err != nil {
		return nil, err
	}

	return results, nil
}
