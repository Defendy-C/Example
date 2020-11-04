package jwt

import (
	"MediaWeb/common"
	"crypto/sha256"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

// 将secret通过sha256加密当盐
func GenerateVToken(uid int, secret string)(tokenStr string, err error) {
	enSecret := sha256.Sum256([]byte(secret))
	fmt.Println(uid)

	var token *jwt.Token
	now := time.Now()
	expTime := now.Add(common.TimeOut)
	charms := jwt.StandardClaims{
		Id:strconv.Itoa(uid),
		ExpiresAt:expTime.Unix(),
	}

	err = charms.Valid()
	if err != nil {
		return
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, charms)
	tokenStr, err = token.SignedString(enSecret[:])

	return
}

func parseToken(tokenStr string, secret string) (token *jwt.Token, err error) {
	secretSlice := sha256.Sum256([]byte(secret))
	enSecret := secretSlice[:]
	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, err error) {
		return enSecret[:], nil
	})

	if err != nil {
		return
	} else if token == nil || !token.Valid {
		err = fmt.Errorf("invalid token")
		return
	}

	return
}

func GetInvaildToken(tokenStr string, secret string)(res string, err error) {
	token , err := parseToken(tokenStr, secret)
	if err != nil {
		return
	}

	m := token.Claims.(jwt.MapClaims)
	m["exp"] = time.Now().Unix()
	token.Claims = m

	res, err = token.SignedString(secret)

	return
}

func ParseVToken(tokenStr string, secret string) (exp int64, uid int, err error) {

	token, err := parseToken(tokenStr, secret)
	if err != nil {
		return
	}
	m := token.Claims.(jwt.MapClaims)
	exp = int64(m["exp"].(float64))
	uid, err = strconv.Atoi(m["jti"].(string))

	return
}