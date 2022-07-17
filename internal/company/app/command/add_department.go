package command

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/domain"
	er "github.com/tttinh/glowing-couscous/pkg/errors"

	ce "github.com/tttinh/glowing-couscous/common/errors"
)

type AddDepartment struct {
	OperatorId string
	Name       string
	Note       string
}

type AddDepartmentRepo interface {
	GetEmployeeById(ctx context.Context, id string) (*domain.Employee, error)
	CreateDepartment(ctx context.Context, d *domain.Department) (*domain.Department, error)
}

type AddDepartmentHandler struct {
	repo AddDepartmentRepo
}

func NewCreateBranchHandler(repo AddDepartmentRepo) AddDepartmentHandler {
	if repo == nil {
		panic("nil repo")
	}

	return AddDepartmentHandler{repo}
}

func (h AddDepartmentHandler) Handle(ctx context.Context, cmd AddDepartment) (string, error) {
	// Creator must be existed.
	e, err := h.repo.GetEmployeeById(ctx, cmd.OperatorId)
	if err != nil {
		return "", err
	}

	if e == nil {
		return "", ce.New(ce.ErrBadRequest, er.ErrUserNotFound.Error())
	}

	// Only CEO can do this.
	if e.Role != domain.EmployeeRole_CEO {
		return "", ce.New(ce.ErrBadRequest, er.ErrBankerUnauthorized.Error())
	}

	d, err := domain.NewDepartmentToCreate(e.CompanyId, cmd.Name, cmd.Note)
	if err != nil {
		return "", err
	}

	d, err = h.repo.CreateDepartment(ctx, d)
	if err != nil {
		return "", err
	}

	return d.Id, nil
}
