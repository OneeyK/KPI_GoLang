package auth

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/OneeyK/KPI_GoLang/pkg/models"
	"github.com/OneeyK/KPI_GoLang/pkg/service/user"
	"github.com/golang-jwt/jwt"
)

const (
	signinkKey = "qwerty123"
	salt       = "123dfhsdghkvkjfxgkjgskeh"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id,omitempty"`
}

type authorizationRepository interface {
	CreateUser(user models.User) (int, error)
	GetUser(name, password string) (models.User, error)
}

type AuthService struct {
	authorizationRepository authorizationRepository
	User                    user.User
}

func NewAuthService(authorizationRepository authorizationRepository, user user.User) Auth {
	return &AuthService{authorizationRepository: authorizationRepository, User: user}
}

type Auth interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(name, password string) (string, error)
	ParseToken(token string) (int, error)
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.authorizationRepository.CreateUser(user)
}

func (s *AuthService) GenerateToken(name, password string) (string, error) {
	user, err := s.authorizationRepository.GetUser(name, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signinkKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing number")
		}
		return []byte(signinkKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
