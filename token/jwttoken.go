package token

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"server_siem/hostinfo"
)

func MakeToken(info *hostinfo.HostInfo) string {
	payload := jwt.MapClaims{
		"sub": info.HostName,
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(info.HostName))
	if err != nil {
		log.Println(err.Error())
	}
	info.Token = token
	return token
}
