package domain

import (
	"context"
)

type Company struct {
	Id     string
	Code   string
	Name   string
	Slogan string
	Logo   string
}

type CompanyRepository interface {
	CreateCompany(ctx context.Context, c *Company) (*Company, error)
	GetCompanyByCode(ctx context.Context, code string) (*Company, error)
}

func NewCompanyForCreate(code, name, slogan, logo string) (*Company, error) {
	return &Company{
		Code:   code,
		Name:   name,
		Slogan: slogan,
		Logo:   logo,
	}, nil
}

func NewCompanyForUpdate(id, code, name, slogan, logo string) (*Company, error) {
	return &Company{
		Id:     id,
		Code:   code,
		Name:   name,
		Slogan: slogan,
		Logo:   logo,
	}, nil
}
