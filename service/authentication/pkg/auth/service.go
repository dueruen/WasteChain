package auth

import (
	"errors"
	"log"

	b64 "encoding/base64"

	pb "github.com/dueruen/WasteChain/service/authentication/gen/proto"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateCredentials(req *pb.CreateCredentialsRequest) (res *pb.CreateCredentialsResponse)
	Login(req *pb.LoginRequest) (res *pb.LoginResponse)
	Validate(req *pb.ValidateRequest) (res *pb.ValidateResponse)
}

type Repository interface {
	SaveCredentials(userID, hashedPassword, username string) error
	GetPassword(username string) (string, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (srv *service) CreateCredentials(req *pb.CreateCredentialsRequest) (res *pb.CreateCredentialsResponse) {
	hashedPassword := hashAndSalt(req.Credentials.Password)
	err := srv.repo.SaveCredentials(req.Credentials.UserID, hashedPassword, req.Credentials.Username)
	if err != nil {
		return &pb.CreateCredentialsResponse{
			Error: err.Error(),
		}
	}
	return &pb.CreateCredentialsResponse{}
}

func (srv *service) Login(req *pb.LoginRequest) (res *pb.LoginResponse) {
	hashedPassword, err := srv.repo.GetPassword(req.Username)
	if err != nil {
		return &pb.LoginResponse{
			Error: err.Error(),
		}
	}
	if !comparePasswords(hashedPassword, req.Password) {
		return &pb.LoginResponse{
			Error: errors.New("Not valid").Error(),
		}
	}

	return &pb.LoginResponse{
		Token: generateToken(req.Username),
	}
}

func (srv *service) Validate(req *pb.ValidateRequest) (res *pb.ValidateResponse) {
	username := decodeToken(req.Token)
	_, err := srv.repo.GetPassword(username)
	if err != nil {
		return &pb.ValidateResponse{
			Error: errors.New("Not valid").Error(),
			Valid: false,
		}
	}

	return &pb.ValidateResponse{
		Valid: true,
	}
}

func generateToken(input string) string {
	return b64.StdEncoding.EncodeToString([]byte(input))
}

func decodeToken(token string) string {
	sDec, _ := b64.StdEncoding.DecodeString(token)
	return string(sDec)
}

func hashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPassword string, plainPassword string) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPassword))
	if err != nil {
		return false
	}

	return true
}
