package query

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/domain"
	er "github.com/tttinh/glowing-couscous/pkg/errors"

	ce "github.com/tttinh/glowing-couscous/common/errors"
)

type GetMyProfileHandler struct {
	repo domain.EmployeeRepository
}

func NewGetMyProfileHandler(repo domain.EmployeeRepository) GetMyProfileHandler {
	if repo == nil {
		panic("nil repo")
	}

	return GetMyProfileHandler{repo}
}

func (h GetMyProfileHandler) Handle(ctx context.Context, employeeId string) (*domain.Employee, error) {
	e, err := h.repo.GetEmployeeById(ctx, employeeId)
	if err != nil {
		return nil, err
	}

	if e == nil {
		return nil, ce.New(ce.ErrBadRequest, er.ErrUserNotFound.Error())
	}

	return e, nil
}
