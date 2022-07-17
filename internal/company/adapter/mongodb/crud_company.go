package mongodb

import (
	"context"

	"github.com/tttinh/glowing-couscous/internal/company/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Repository) CreateCompany(ctx context.Context, c *domain.Company) (*domain.Company, error) {
	c.Id = ""
	m, err := CompanyFromDomain(c)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.CreateCompany: BankFromDomain")
	}

	m.SetCreatedAt()
	rs, err := r.companies.InsertOne(ctx, m)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.CreateCompany: collection.InsertOne")
	}

	m.Id = rs.InsertedID.(primitive.ObjectID)
	company, err := CompanyToDomain(m)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.CreateCompany: CompanyToDomain")
	}

	return company, nil
}

func (r Repository) UpdateCompany(ctx context.Context, b *domain.Company) (bool, error) {
	m, err := CompanyFromDomain(b)
	if err != nil {
		return false, errors.Wrap(err, "Repository.UpdateCompany: BankFromDomain")
	}

	filter := bson.M{"_id": m.Id}
	_, err = r.companies.UpdateOne(ctx, filter, bson.M{"$set": m})
	if err != nil {
		return false, errors.Wrap(err, "Repository.UpdateCompany: collection.UpdateOne")
	}

	return true, nil
}

func (r Repository) GetCompanyByCode(ctx context.Context, codeName string) (*domain.Company, error) {
	var c *Company
	err := r.companies.FindOne(ctx, bson.M{"code": codeName}).Decode(&c)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "Repository.GetCompanyByCode: FindOne")
	}

	company, err := CompanyToDomain(c)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.GetCompanyByCode: CompanyToDomain")
	}

	return company, nil
}

func (r Repository) GetCompanyById(ctx context.Context, id string) (*domain.Company, error) {
	var c *Company
	companyId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = r.companies.FindOne(ctx, bson.M{"_id": companyId}).Decode(&c)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "Repository.GetCompanyById: FindOne")
	}

	company, err := CompanyToDomain(c)
	if err != nil {
		return nil, errors.Wrap(err, "Repository.GetCompanyById: CompanyToDomain")
	}

	return company, nil
}

func (r Repository) ListCompany(ctx context.Context) ([]*domain.Company, error) {
	cursor, err := r.companies.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "Repository.ListCompany: Find")
	}

	var companyModels []*Company
	if err = cursor.All(ctx, &companyModels); err != nil {
		return nil, errors.Wrap(err, "Repository.ListCompany: cursor.All")
	}

	var results []*domain.Company
	for _, m := range companyModels {
		c, err := CompanyToDomain(m)
		if err != nil {
			return nil, errors.Wrap(err, "Repository.ListCompany: CompanyToDomain")
		}
		results = append(results, c)
	}

	return results, nil
}
