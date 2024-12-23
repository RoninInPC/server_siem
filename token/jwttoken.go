package token

import (
	"github.com/golang-jwt/jwt/v5"
	"server_siem/hostinfo"
)

func MakeToken(info *hostinfo.HostInfo) string {
	payload := jwt.MapClaims{
		"sub": info.HostName,
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(info.HostName)
	info.Token = token
	return token
}
