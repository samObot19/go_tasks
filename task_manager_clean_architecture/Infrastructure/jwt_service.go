package Infrastructure

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/task_manager/Domain"
)

type JWTService struct {
    SigningKey []byte
}

func NewJWTService(signingKey string) *JWTService {
	return &JWTService{
		SigningKey: []byte(signingKey),
    	}
}


func (s *JWTService) GenerateToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.UserName,
		"role":   	user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
    	}

    	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    	return token.SignedString(s.SigningKey)
}

func (s *JWTService) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
        	}
        	return s.SigningKey, nil
    	})
    	return token, err
}


