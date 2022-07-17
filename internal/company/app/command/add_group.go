package command

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/domain"
	er "github.com/tttinh/glowing-couscous/pkg/errors"

	ce "github.com/tttinh/glowing-couscous/common/errors"
)

type AddGroup struct {
	OperatorId   string
	DepartmentId string
	Name         string
	Note         string
}

type AddGroupRepo interface {
	GetEmployeeById(ctx context.Context, id string) (*domain.Employee, error)
	GetDepartmentById(ctx context.Context, id string) (*domain.Department, error)
	CreateGroup(ctx context.Context, g *domain.Group) (*domain.Group, error)
}

type AddGroupHandler struct {
	repo AddGroupRepo
}

func NewAddGroupHandler(repo AddGroupRepo) AddGroupHandler {
	if repo == nil {
		panic("nil repo")
	}

	return AddGroupHandler{repo}
}

func (h AddGroupHandler) Handle(ctx context.Context, cmd AddGroup) (string, error) {
	// Creator must be existed.
	e, err := h.repo.GetEmployeeById(ctx, cmd.OperatorId)
	if err != nil {
		return "", err
	}

	if e == nil {
		return "", ce.New(ce.ErrBadRequest, er.ErrUserNotFound.Error())
	}

	// Department must be existed
	d, err := h.repo.GetDepartmentById(ctx, cmd.DepartmentId)
	if err != nil {
		return "", err
	}

	if d == nil {
		return "", ce.New(ce.ErrBadRequest, er.ErrBankBranchIdInvalid.Error())
	}

	isCEO := e.Role == domain.EmployeeRole_CEO && e.CompanyId == d.CompanyId
	isHead := e.Role == domain.EmployeeRole_HEAD && e.DepartmentId == d.Id

	if !isCEO && !isHead {
		return "", ce.New(ce.ErrBadRequest, er.ErrBankerUnauthorized.Error())
	}

	g, err := domain.NewGroupToCreate(
		d.CompanyId,
		cmd.DepartmentId,
		cmd.Name,
		cmd.Note,
	)
	if err != nil {
		return "", err
	}

	g, err = h.repo.CreateGroup(ctx, g)
	if err != nil {
		return "", err
	}

	return g.Id, nil
}
