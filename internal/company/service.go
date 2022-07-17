package company

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/tttinh/glowing-couscous/internal/company/adapter/client"
	"github.com/tttinh/glowing-couscous/internal/company/adapter/mongodb"
	"github.com/tttinh/glowing-couscous/internal/company/app"
	"github.com/tttinh/glowing-couscous/internal/company/app/command"
	"github.com/tttinh/glowing-couscous/internal/company/app/query"
	"github.com/tttinh/glowing-couscous/internal/company/port"
)

func New(authClient *client.User, db *mongo.Database) *port.Server {
	repo := mongodb.NewRepository(authClient, db)

	app := app.Application{
		Commands: app.Commands{
			AddCEO:        command.NewAddCEOHandler(repo),
			AddCompany:    command.NewAddCompanyHandler(repo),
			AddDepartment: command.NewCreateBranchHandler(repo),
			AddGroup:      command.NewAddGroupHandler(repo),
			UpdateCompany: command.NewUpdateCompanyHandler(repo),
		},
		Queries: app.Queries{
			GetMyProfile:     query.NewGetMyProfileHandler(repo),
			GetMyCompany:     query.NewGetMyCompanyHandler(repo),
			GetCompanyByCode: query.NewGetBankByCodeHandler(repo),
			ListCompany:      query.NewListCompanyHandler(repo),
			ListDepartment:   query.NewListDepartmentHandler(repo),
			ListGroup:        query.NewListGroupHandler(repo),
		},
	}

	return port.NewServer(app)
}
