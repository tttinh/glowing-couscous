package port

import (
	"context"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"
)

// GetCompanyByCode: get a company profile by its code.
func (s *Server) GetCompanyByCode(
	ctx context.Context,
	req *genpb.GetCompanyByCodeReq,
) (*genpb.GetCompanyByCodeRes, error) {
	c, err := s.app.Queries.GetCompanyByCode.Handle(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	if c == nil {
		return &genpb.GetCompanyByCodeRes{}, nil
	}

	return &genpb.GetCompanyByCodeRes{
		Id:     c.Id,
		Code:   c.Code,
		Name:   c.Name,
		Slogan: c.Slogan,
		Logo:   c.Logo,
	}, nil
}
