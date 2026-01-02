package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// =======================
// JWT CLAIMS
// =======================

type Claims struct {
	UserID uint     `json:"user_id"`
	Scope  []string `json:"scope"`
	jwt.RegisteredClaims
}

// =======================
// GENERATE JWT FOR USER
// =======================

func GenerateUserToken(userID uint) (string, error) {
	expDays := 30
	exp := time.Now().Add(time.Duration(expDays) * 24 * time.Hour)

	claims := Claims{
		UserID: userID,
		Scope:  []string{"foods:read"},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("JWT_ISSUER"),
			Audience:  []string{os.Getenv("JWT_AUDIENCE")},
			Subject:   fmt.Sprintf("user:%d", userID),
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// =======================
// PARSE & VALIDATE JWT
// =======================

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
