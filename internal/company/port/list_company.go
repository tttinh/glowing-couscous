package port

import (
	"context"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"
	"github.com/tttinh/glowing-couscous/internal/company/domain"
)

// ListCompany: lists all companies.
func (s *Server) ListCompany(
	ctx context.Context,
	_ *genpb.CompanyListCompanyReq,
) (*genpb.CompanyListCompanyRes, error) {
	companies, err := s.app.Queries.ListCompany.Handle(ctx)
	if err != nil {
		return nil, err
	}

	var result []*genpb.CompanyProfile
	for _, c := range companies {
		info := getCompanyProfile(c)
		result = append(result, info)
	}

	return &genpb.CompanyListCompanyRes{Result: result}, nil
}

func getCompanyProfile(c *domain.Company) *genpb.CompanyProfile {
	return &genpb.CompanyProfile{
		Id:     c.Id,
		Code:   c.Code,
		Name:   c.Name,
		Slogan: c.Slogan,
		Logo:   c.Logo,
	}
}
