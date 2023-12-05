package middlewares

import (
	"fmt"
	"lib/cmd/api"
	"lib/models"
	"lib/utils"
	"net/http"
	"strings"
)

func ProtectedJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := parseToken(r)
		if err != nil {
			permissionDenied(w)
			return
		}

		validated, err := utils.ValidateJWT(token)
		if err != nil {
			permissionDenied(w)
			return
		}

		if !validated.Valid {
			permissionDenied(w)
			return
		}

		handlerFunc(w, r)
	}
}

func parseToken(r *http.Request) (string, error) {
	authentication_credential := r.Header.Get("Authorization")

	parsed_token := strings.Split(authentication_credential, " ")
	if len(parsed_token) == 2 {
		return parsed_token[1], nil
	}

	return "", fmt.Errorf("error occurred while parsing token")
}

func permissionDenied(w http.ResponseWriter) error {
	result := new(models.Result)
	result.Success = false
	result.ErrorCode = utils.ERR0401
	result.ErrorDescription = utils.ERR0401.ToDescription()

	return api.WriteJSON(w, http.StatusUnauthorized, result)
}
