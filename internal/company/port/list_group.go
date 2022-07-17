package port

import (
	"context"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"
	"github.com/tttinh/glowing-couscous/internal/company/domain"
)

// ListGroup: lists all groups of a department.
func (s *Server) ListGroup(
	ctx context.Context,
	req *genpb.CompanyListGroupReq,
) (*genpb.CompanyListGroupRes, error) {
	userId := getUserId(ctx)
	groups, err := s.app.Queries.ListGroup.Handle(ctx, userId, req.DepartmentId)
	if err != nil {
		return nil, err
	}

	var result []*genpb.CompanyGroup
	for _, g := range groups {
		info := getTeamInfo(g)
		result = append(result, info)
	}

	return &genpb.CompanyListGroupRes{Result: result}, nil
}

func getTeamInfo(g *domain.Group) *genpb.CompanyGroup {
	return &genpb.CompanyGroup{
		Id:          g.Id,
		Name:        g.Name,
		Note:        g.Note,
		MemberCount: int32(g.EmployeeCount),
	}
}
