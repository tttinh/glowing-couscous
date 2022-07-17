package port

import (
	"context"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"
	"github.com/tttinh/glowing-couscous/internal/company/domain"
)

// ListDepartment: list all departments of a company.
func (s *Server) ListDepartment(
	ctx context.Context,
	req *genpb.CompanyListDepartmentReq,
) (*genpb.CompanyListDepartmentRes, error) {
	departments, err := s.app.Queries.ListDepartment.Handle(ctx, req.CompanyId)
	if err != nil {
		return nil, err
	}

	var result []*genpb.CompanyDepartment
	for _, d := range departments {
		info := getDepartmentInfo(d)
		result = append(result, info)
	}

	return &genpb.CompanyListDepartmentRes{Result: result}, nil
}

func getDepartmentInfo(b *domain.Department) *genpb.CompanyDepartment {
	return &genpb.CompanyDepartment{
		Id:          b.Id,
		Name:        b.Name,
		Note:        b.Note,
		MemberCount: int32(b.EmployeeCount),
	}
}
