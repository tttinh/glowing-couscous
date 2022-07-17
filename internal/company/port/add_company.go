package port

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/app/command"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"
)

// AddCompany: create a new company.
func (s *Server) AddCompany(
	ctx context.Context,
	req *genpb.CompanyAddReq,
) (*genpb.CompanyAddRes, error) {
	cmd := command.AddCompany{
		Code:   req.Code,
		Name:   req.Name,
		Slogan: req.Slogan,
		Logo:   req.Logo,
	}
	id, err := s.app.Commands.AddCompany.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &genpb.CompanyAddRes{Id: id}, nil
}
