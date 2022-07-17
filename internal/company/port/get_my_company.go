package port

import (
	"context"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"
)

// GetMyCompany: Get a company profile that this user belongs to.
func (s *Server) GetMyCompany(
	ctx context.Context,
	_ *genpb.CompanyGetMyCompanyReq,
) (*genpb.CompanyGetMyCompanyRes, error) {
	userId := getUserId(ctx)
	profile, err := s.app.Queries.GetMyCompany.Handle(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &genpb.CompanyGetMyCompanyRes{
		Id:     profile.Id,
		Code:   profile.Code,
		Name:   profile.Name,
		Slogan: profile.Slogan,
		Logo:   profile.Logo,
	}, nil
}
