package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub string
	Exp int
}

/* Payload - https://jwt.io/#debugger-io
type TokenJSON struct {
	At_hash        string
	Sub            string
	Email_verified bool
	Iss            string
	Username       string
	Aud            string
	Event_id       string
	Token_use      string
	Auth_time      int
	Exp            int
	Iat            int
	Jti            string
	Email          string
}*/

func ValidoToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		fmt.Println("El token no es válido")
		return false, nil, "El token no es válido"
	}

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])

	if err != nil {
		fmt.Println("No se puede decodificar la parte del token", err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)

	if err != nil {
		fmt.Println("No se puede decodificar en la estructura JSON", err.Error())
		return false, err, err.Error()
	}

	ahora := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	if tm.Before(ahora) {
		fmt.Println("Fecha expiración token: " + tm.String())
		fmt.Println("Token expirado")
		return false, nil, "Token expirado"
	}

	return true, nil, string(tkj.Sub)

}
