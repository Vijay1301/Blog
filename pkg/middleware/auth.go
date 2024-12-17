package middleware

import "github.com/golang-jwt/jwt"

func (t TokenPayload) Valid() error {
	return nil
}

func GenerateJWTTokens(payload TokenPayload) (string, string, error) {

	accessToken, err := GenerateJWTToken(payload, "JWT_SECRET")
	if err != nil {
		return "", "", err
	}
	refreshToken, err := GenerateJWTToken(payload, "JWT_REFRESH_SECRET")
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// GenerateJWTToken change by Andrews
func GenerateJWTToken(payload TokenPayload, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
