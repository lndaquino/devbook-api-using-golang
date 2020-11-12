package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken returns a jwt token
func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix() // 6 hours of expiration time
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions) // setting signing method and permissions
	return token.SignedString([]byte(config.SecretKey))             // signing token with secret
}

// ValidateToken validates if token is valid
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, keyVerification) // check if the signing method is the same to unsing
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Expired token")
}

// GetUserID extracts userID from jwt token
func GetUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, keyVerification) // check if the signing method is the same to unsing
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		parsedUserID := fmt.Sprintf("%.0f", permissions["userID"]) // claim is saved as a float
		userID, err := strconv.ParseUint(parsedUserID, 10, 64)     // converters string o uint
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New("Invalid token")
}

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	token := strings.Split(bearerToken, " ")

	if len(token) == 2 {
		return token[1]
	}

	return ""
}

// check if the signing method is the same used to generate token
func keyVerification(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Signing method not allowed! %v", token.Header["alg"])
	}

	return []byte(config.SecretKey), nil
}
