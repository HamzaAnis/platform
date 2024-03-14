package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (j *jwtImpl) GenerateToken(userID int64) (*string, error) {
	presentTime := time.Now().UTC()
	claims := &CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  presentTime.Unix(),
			ExpiresAt: presentTime.Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		log.Errorf("unable to generate token: %v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	bearerToken := fmt.Sprintf("Bearer %s", tokenString)
	return &bearerToken, nil
}
