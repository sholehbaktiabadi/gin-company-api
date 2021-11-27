package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/sholehbaktiabadi/go-api/dto"
	"github.com/sholehbaktiabadi/go-api/entity"
	"github.com/sholehbaktiabadi/go-api/repository"
)

type CompanyService interface {
	Insert(company dto.CompanyCreateDTO) entity.Company
	Update(company dto.CompanyUpdateDTO) entity.Company
	Delete(company entity.Company)
	AllCompany() []entity.Company
	FindByID(companyID uint64) entity.Company
	IsAllowedToEdit(userID string, companyID uint64) bool
}

type companyService struct {
	companyRepository repository.CompanyRepository
}

func NewCompanyService(companyRepository repository.CompanyRepository) CompanyService {
	return &companyService{
		companyRepository: companyRepository,
	}
}

func (service *companyService) Insert(company dto.CompanyCreateDTO) entity.Company {
	companyData := entity.Company{}
	err := smapping.FillStruct(&companyData, smapping.MapFields(&company))
	if err != nil {
		log.Fatalf("smapping data failed %v", err)
	}
	createUser := service.companyRepository.InsertCompany(companyData)
	return createUser

}

func (service *companyService) Update(company dto.CompanyUpdateDTO) entity.Company {
	companyData := entity.Company{}
	err := smapping.FillStruct(&companyData, smapping.MapFields(company))
	if err != nil {
		log.Fatalf("smapping data failed %v", err)
	}

	updateUser := service.companyRepository.UpdateCompany(companyData)
	return updateUser
}

func (service *companyService) Delete(company entity.Company) {
	service.companyRepository.DeleteCompany(company)

}

func (service *companyService) AllCompany() []entity.Company {
	return service.companyRepository.AllCompany()
}

func (service *companyService) FindByID(companyID uint64) entity.Company {
	return service.companyRepository.FindCompanyByID(companyID)
}

func (service *companyService) IsAllowedToEdit(userID string, companyID uint64) bool {
	company := service.companyRepository.FindCompanyByID(companyID)
	id := fmt.Sprintf("%v", company.UserID)
	return userID == id

}
