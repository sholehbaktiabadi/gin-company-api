package repository

import (
	"github.com/sholehbaktiabadi/go-api/entity"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	InsertCompany(company entity.Company) entity.Company
	UpdateCompany(company entity.Company) entity.Company
	AllCompany() []entity.Company
	DeleteCompany(company entity.Company)
	FindCompanyByID(companyID uint64) entity.Company
}

type companyConnection struct {
	companyConnection *gorm.DB
}

func NewCompanyRespository(dbConnection *gorm.DB) CompanyRepository {
	return &companyConnection{
		companyConnection: dbConnection,
	}
}

func (db *companyConnection) InsertCompany(company entity.Company) entity.Company {
	db.companyConnection.Save(&company)
	db.companyConnection.Preload("User").Find(&company)
	return company
}

func (db *companyConnection) UpdateCompany(company entity.Company) entity.Company {
	db.companyConnection.Save(&company)
	db.companyConnection.Preload("User").Find(&company)
	return company
}

func (db *companyConnection) DeleteCompany(company entity.Company) {
	db.companyConnection.Delete(&company)
}

func (db *companyConnection) FindCompanyByID(companyID uint64) entity.Company {
	var company entity.Company
	db.companyConnection.Preload("User").Find(&companyID, company)
	return company
}

func (db *companyConnection) AllCompany() []entity.Company {
	var company []entity.Company
	db.companyConnection.Preload("User").Find(&company)
	return company
}
