package port

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/app/command"

	"github.com/golang/protobuf/ptypes/empty"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"
)

// AddCEO: create a new CEO for a company.
func (s *Server) AddCEO(
	ctx context.Context,
	req *genpb.CompanyAddCEOReq,
) (*empty.Empty, error) {
	cmd := command.AddCEO{
		EmployeeId: req.EmployeeId,
		CompanyId:  req.CompanyId,
	}

	err := s.app.Commands.AddCEO.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
