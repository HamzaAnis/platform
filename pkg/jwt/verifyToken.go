package jwt

import (
	"context"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (j *jwtImpl) VerifyToken(ctx context.Context) (*CustomClaims, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		keyFunc := func(token *jwt.Token) (interface{}, error) {
			return []byte(j.secret), nil
		}

		if vals, ok := md["authorization"]; ok && len(vals) > 0 {
			token := strings.TrimPrefix(vals[0], "Bearer ")
			jwtToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, keyFunc)
			v, _ := err.(*jwt.ValidationError)
			if err != nil {
				if v.Errors != jwt.ValidationErrorExpired {
					return nil, status.Errorf(codes.Unauthenticated, "invalid token")
				}
				return nil, status.Errorf(codes.Unauthenticated, err.Error())
			}

			claims, ok := jwtToken.Claims.(*CustomClaims)
			if (!ok) || (!jwtToken.Valid && v.Errors != jwt.ValidationErrorExpired) {
				return nil, status.Errorf(codes.Unauthenticated, "missing claims")
			}

			return claims, nil
		}
	}
	return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
}
