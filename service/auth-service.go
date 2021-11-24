package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/sholehbaktiabadi/go-api/dto"
	"github.com/sholehbaktiabadi/go-api/entity"
	"github.com/sholehbaktiabadi/go-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDto) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepos repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepos,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparePasswod := ComparePassword(v.Password, []byte(password))
		if v.Email == email && comparePasswod {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDto) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed mas %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func ComparePassword(hash_password string, plain_password []byte) bool {
	byteHash := []byte(hash_password)
	err := bcrypt.CompareHashAndPassword(byteHash, plain_password)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
