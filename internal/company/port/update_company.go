package port

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"
	"github.com/tttinh/glowing-couscous/internal/company/app/command"
)

// UpdateCompany: update a company profile.
func (s *Server) UpdateCompany(
	ctx context.Context,
	req *genpb.CompanyUpdateReq,
) (*empty.Empty, error) {
	userId := getUserId(ctx)
	cmd := command.UpdateCompany{
		OperatorId: userId,
		Code:       req.Code,
		Name:       req.Name,
		Slogan:     req.Slogan,
		Logo:       req.Logo,
	}

	err := s.app.Commands.UpdateCompany.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
