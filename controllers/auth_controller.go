package controllers

import (
	"encoding/json"
	"fmt"
	"lib/cmd/api"
	"lib/database"
	"lib/models"
	"lib/utils"
	"net/http"
)

// Login		 LoginAccount godoc
// @Summary      Login to your account
// @Description  Login with username and password
// @Tags         Login
// @Accept       json
// @Produce      json
// @Param        userModelArgs body models.UserLoginArgs true "UserLogin"
// @Success      200  {object}  models.UserLoginResult
// @Router       /login [post]
func (controller BaseController) Login(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	userLoginResult := new(models.UserLoginResult)
	userLoginArgs := new(models.UserLoginArgs)

	if err := json.NewDecoder(r.Body).Decode(userLoginArgs); err != nil {
		userLoginResult.Result.Success = false
		userLoginResult.Result.ErrorCode = utils.ERR0303
		userLoginResult.Result.ErrorDescription = utils.ERR0303.ToDescription()
		userLoginResult.Result.ErrorException = utils.ExceptionToString(err)

		return api.WriteJSON(w, http.StatusOK, userLoginResult)
	}

	if ok := utils.ValidateCheckSpaceCharacter(userLoginArgs.Email, userLoginArgs.Password); !ok {
		userLoginResult.Result.Success = false
		userLoginResult.Result.ErrorCode = utils.ERR0304
		userLoginResult.Result.ErrorDescription = utils.ERR0304.ToDescription()

		return api.WriteJSON(w, http.StatusOK, userLoginResult)
	}

	if ok := utils.ValidateEmail(userLoginArgs.Email); !ok {
		userLoginResult.Result.Success = false
		userLoginResult.Result.ErrorCode = utils.ERR0304
		userLoginResult.Result.ErrorDescription = utils.ERR0304.ToDescription()

		return api.WriteJSON(w, http.StatusOK, userLoginResult)
	}

	if ok := utils.ValidatePassword(userLoginArgs.Password); !ok {
		userLoginResult.Result.Success = false
		userLoginResult.Result.ErrorCode = utils.ERR0304
		userLoginResult.Result.ErrorDescription = utils.ERR0304.ToDescription()

		return api.WriteJSON(w, http.StatusOK, userLoginResult)
	}

	model := database.Model[models.UserModel]{
		Stg: controller.Storage.GetCursor(),
	}
	result, err := model.Get(fmt.Sprintf("email = %s", userLoginArgs.Email))
	if err != nil {
		userLoginResult.Result.Success = false
		userLoginResult.Result.ErrorCode = utils.ERR0402
		userLoginResult.Result.ErrorDescription = utils.ERR0402.ToDescription()
		userLoginResult.Result.ErrorException = utils.ExceptionToString(err)

		return api.WriteJSON(w, http.StatusOK, userLoginResult)
	}

	if ok := utils.CompareHashAndPassword(result.Password, userLoginArgs.Password); !ok {
		userLoginResult.Result.Success = false
		userLoginResult.Result.ErrorCode = utils.ERR0403
		userLoginResult.Result.ErrorDescription = utils.ERR0403.ToDescription()
		userLoginResult.Result.ErrorException = utils.ExceptionToString(err)

		return api.WriteJSON(w, http.StatusOK, userLoginResult)
	}

	token, err := utils.CreateJSONWebToken()
	if err != nil {
		userLoginResult.Result.Success = false
		userLoginResult.Result.ErrorCode = utils.ERR0405
		userLoginResult.Result.ErrorDescription = utils.ERR0405.ToDescription()
		userLoginResult.Result.ErrorException = utils.ExceptionToString(err)

		return api.WriteJSON(w, http.StatusOK, userLoginResult)
	}

	userLoginResult.Id = result.Id
	userLoginResult.AuthenticationToken = token
	userLoginResult.UserInfos = map[string]string{
		"full_name":    result.FirstName + " " + result.LastName,
		"phone_number": result.PhoneNumber,
		"email":        result.Email,
	}

	userLoginResult.Result.Success = true
	userLoginResult.Result.ErrorCode = ""
	userLoginResult.Result.ErrorDescription = ""
	userLoginResult.Result.ErrorException = ""

	return api.WriteJSON(w, http.StatusOK, userLoginResult)
}

func (controller BaseController) Register(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	return api.WriteJSON(w, http.StatusOK, nil)
}
