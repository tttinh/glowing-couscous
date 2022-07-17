package port

import (
	"context"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"
	"github.com/tttinh/glowing-couscous/internal/company/app/command"
)

// AddGroup: Create a new group.
func (s *Server) AddGroup(
	ctx context.Context,
	req *genpb.CompanyAddGroupReq,
) (*genpb.CompanyAddGroupRes, error) {
	userId := getUserId(ctx)
	cmd := command.AddGroup{
		OperatorId:   userId,
		DepartmentId: req.DepartmentId,
		Name:         req.Name,
		Note:         req.Note,
	}
	id, err := s.app.Commands.AddGroup.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &genpb.CompanyAddGroupRes{Id: id}, nil
}
