package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

// CreateJWTToken for JWT token
func CreateJWTToken(UserID uint64) (*RedisTokenDetails, error) {

	var err error
	td := &RedisTokenDetails{}
	//td.AtExpires = time.Now().Add(time. * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()
	td.UserID = UserID

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = UserID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	saveErr := CreateRedisAuth(UserID, td)
	if saveErr != nil {
		return nil, saveErr
	}

	return td, nil
}

//ExtractToken from header
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

//VerifyToken to verify
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//TokenValid to check validity
func TokenValid(r *http.Request) (jwt.MapClaims, bool) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}

	return nil, false
}
