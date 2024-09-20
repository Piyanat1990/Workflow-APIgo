package user

import (
	"errors"
	"log"

	"github.com/Piyanat1990/workflow/internal/auth"
	"github.com/Piyanat1990/workflow/internal/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	Repository Repository
	secret     string
}

func NewService(db *gorm.DB, secret string) Service {
	return Service{
		Repository: NewRepository(db),
		secret:     secret,
	}
}

// return stirng token,
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.UhG3lFvqII3ReCCIRzxl6g-VwyNSGmWZy_Q22_MPNxQ

func (service Service) Login(req model.RequestLogin) (string, error) {
	// TODO: Check username and password here
	//Check user in database
	user,err := service.Repository.FindOneByUsername(req.Username)
	if err!=nil{
		return"",errors.New("Invalid user or password")
	}

// req.Password //req password  (plaint text)
// user.Password //hashed password (valid)

if ok:= checkPasswordHash(req.Password,user.Password); !ok{
	return "",errors.New("Invalid user or password")
}

	// TODO: Create token here

	token,err := auth.CreateToken(user.Username,service.secret)
	if err!=nil{
		log.Println("Fail to create token")
		return "",errors.New("something went wrong")
	}


	return token, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
