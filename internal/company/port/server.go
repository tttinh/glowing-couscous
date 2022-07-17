package port

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/app"
	"github.com/tttinh/glowing-couscous/internal/company/domain"

	genpb "github.com/tttinh/glowing-couscous/api/genproto/go"
)

const CtxUserIdKey = "x-user-id"

type Server struct {
	*genpb.UnimplementedCompanyServiceServer
	app app.Application
}

func NewServer(application app.Application) *Server {
	return &Server{app: application}
}

func getUserId(ctx context.Context) string {
	userId, _ := ctx.Value(CtxUserIdKey).(string)
	return userId
}

func roleDomain2Proto(role domain.EmployeeRole) genpb.CompanyRole {
	switch role {
	case domain.EmployeeRole_CEO:
		return genpb.CompanyRole_COMPANY_ROLE_CEO
	case domain.EmployeeRole_HEAD:
		return genpb.CompanyRole_COMPANY_ROLE_HEAD
	case domain.EmployeeRole_LEAD:
		return genpb.CompanyRole_COMPANY_ROLE_LEAD
	case domain.EmployeeRole_STAFF:
		return genpb.CompanyRole_COMPANY_ROLE_STAFF
	default:
		return genpb.CompanyRole_COMPANY_ROLE_INVALID
	}
}
