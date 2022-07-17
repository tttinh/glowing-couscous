package query

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/domain"
)

type ListDepartmentRepo interface {
	ListDepartment(ctx context.Context, companyId string) ([]*domain.Department, error)
}

type ListDepartmentHandler struct {
	repo ListDepartmentRepo
}

func NewListDepartmentHandler(repo ListDepartmentRepo) ListDepartmentHandler {
	if repo == nil {
		panic("nil repo")
	}

	return ListDepartmentHandler{repo}
}

func (h ListDepartmentHandler) Handle(ctx context.Context, companyId string) (results []*domain.Department, err error) {
	results, err = h.repo.ListDepartment(ctx, companyId)
	if err != nil {
		return nil, err
	}

	return results, nil
}
