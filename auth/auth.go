package auth

import (
	"time"

	"goback/core/output"
	"goback/utils"
)

// Token which is returned in response to an authorization request
type Token struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
	ExpiresIn    int64  `json:"expires"`
}

func Login(identificator string, password string) {
	//Претворимся что пользователь прошёл авторизацию

	authToken, err := startSession(1)

	if err != nil {
		output.Error()
	}

	//Возвращаем токен
}

func startSession(userID uint) (Token, error) {
	accessToken := utils.RandStr(64)

	refreshToken := utils.RandStr(32)

	expiresIn := time.Now().Add(time.Hour).Unix()

	//Записываем это в базу

	//Возвращаем токен
	authToket := Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}

	return authToket, nil
}
