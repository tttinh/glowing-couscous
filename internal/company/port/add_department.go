package port

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/app/command"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"
)

// AddDepartment: Create a new department.
func (s *Server) AddDepartment(
	ctx context.Context,
	req *genpb.CompanyAddDepartmentReq,
) (*genpb.CompanyAddDepartmentRes, error) {
	userId := getUserId(ctx)
	cmd := command.AddDepartment{
		OperatorId: userId,
		Name:       req.Name,
		Note:       req.Note,
	}

	id, err := s.app.Commands.AddDepartment.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &genpb.CompanyAddDepartmentRes{Id: id}, nil
}
