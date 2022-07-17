package app

import (
	"github.com/tttinh/glowing-couscous/internal/company/app/command"
	"github.com/tttinh/glowing-couscous/internal/company/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AddCEO        command.AddCEOHandler
	AddCompany    command.AddCompanyHandler
	AddDepartment command.AddDepartmentHandler
	AddGroup      command.AddGroupHandler
	UpdateCompany command.UpdateCompanyHandler
}

type Queries struct {
	GetMyProfile     query.GetMyProfileHandler
	GetMyCompany     query.GetMyCompanyHandler
	GetCompanyByCode query.GetCompanyByCodeHandler
	ListCompany      query.ListCompanyHandler
	ListDepartment   query.ListDepartmentHandler
	ListGroup        query.ListGroupHandler
}
