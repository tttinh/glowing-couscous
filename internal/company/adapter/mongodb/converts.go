package mongodb

import (
	"github.com/tttinh/glowing-couscous/internal/company/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CompanyToDomain(c *Company) (*domain.Company, error) {
	return &domain.Company{
		Id:     c.Id.Hex(),
		Code:   c.Code,
		Name:   c.Name,
		Slogan: c.Slogan,
		Logo:   c.Logo,
	}, nil
}

func CompanyFromDomain(c *domain.Company) (*Company, error) {
	company := &Company{
		Code:   c.Code,
		Name:   c.Name,
		Slogan: c.Slogan,
		Logo:   c.Logo,
	}

	if c.Id != "" {
		objId, err := primitive.ObjectIDFromHex(c.Id)
		if err != nil {
			return nil, err
		}

		company.Id = objId
	}

	return company, nil
}

func EmployeeToDomain(e *Employee) (*domain.Employee, error) {
	return &domain.Employee{
		Id:           e.Id,
		Role:         domain.EmployeeRole(e.Role),
		CompanyId:    e.CompanyId,
		DepartmentId: e.DepartmentId,
		GroupId:      e.GroupId,
	}, nil
}

func EmployeeFromDomain(e *domain.Employee) (*Employee, error) {
	return &Employee{
		Id:           e.Id,
		Role:         string(e.Role),
		CompanyId:    e.CompanyId,
		DepartmentId: e.DepartmentId,
		GroupId:      e.GroupId,
	}, nil
}

func DepartmentToDomain(d *Department) (*domain.Department, error) {
	return &domain.Department{
		Id:            d.Id.Hex(),
		Name:          d.Name,
		Note:          d.Note,
		CompanyId:     d.CompanyId,
		EmployeeCount: d.EmployeeCount,
	}, nil
}

func DepartmentFromDomain(d *domain.Department) (*Department, error) {
	department := &Department{
		Name:      d.Name,
		Note:      d.Note,
		CompanyId: d.CompanyId,
	}

	if d.Id != "" {
		objId, err := primitive.ObjectIDFromHex(d.Id)
		if err != nil {
			return nil, err
		}

		department.Id = objId
	}

	return department, nil
}

func GroupToDomain(g *Group) (*domain.Group, error) {
	return &domain.Group{
		Id:            g.Id.Hex(),
		Name:          g.Name,
		Note:          g.Note,
		CompanyId:     g.CompanyId,
		DepartmentId:  g.DepartmentId,
		EmployeeCount: g.EmployeeCount,
	}, nil
}

func GroupFromDomain(g *domain.Group) (*Group, error) {
	group := &Group{
		Name:         g.Name,
		Note:         g.Note,
		CompanyId:    g.CompanyId,
		DepartmentId: g.DepartmentId,
	}

	if g.Id != "" {
		objId, err := primitive.ObjectIDFromHex(g.Id)
		if err != nil {
			return nil, err
		}

		group.Id = objId
	}

	return group, nil
}
