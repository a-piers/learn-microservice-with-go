package controllers

import (
	"encoding/json"
	"fmt"
	"lib/cmd/api"
	"lib/database"
	"lib/models"
	"lib/utils"
	"net/http"
	"reflect"
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

// Register		 RegisterAccount godoc
// @Summary      Create a account
// @Description  Register and create account
// @Tags         Register
// @Accept       json
// @Produce      json
// @Param        userModelArgs body models.UserRegisterArgs true "UserRegister"
// @Success      200  {object}  models.UserRegisterResult
// @Router       /register [post]
func (controller BaseController) Register(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	userRegisterArgs := new(models.UserRegisterArgs)
	userRegisterResult := new(models.UserRegisterResult)

	if err := json.NewDecoder(r.Body).Decode(userRegisterArgs); err != nil {
		userRegisterResult.Result.Success = false
		userRegisterResult.Result.ErrorCode = utils.ERR0303
		userRegisterResult.Result.ErrorDescription = utils.ERR0303.ToDescription()
		userRegisterResult.Result.ErrorException = utils.ExceptionToString(err)

		return api.WriteJSON(w, http.StatusOK, userRegisterResult)
	}

	if ok := utils.ValidateCheckSpaceCharacter(
		userRegisterArgs.FirstName,
		userRegisterArgs.LastName,
		userRegisterArgs.Email,
		userRegisterArgs.Password,
		userRegisterArgs.ValidatePassword,
		userRegisterArgs.PhoneNumber,
		userRegisterArgs.BirthDate.GoString(),
	); !ok {
		userRegisterResult.Result.Success = false
		userRegisterResult.Result.ErrorCode = utils.ERR0304
		userRegisterResult.Result.ErrorDescription = utils.ERR0304.ToDescription()

		return api.WriteJSON(w, http.StatusOK, userRegisterResult)
	}

	if ok := utils.IsStringEqual(userRegisterArgs.Password, userRegisterArgs.ValidatePassword); !ok {
		userRegisterResult.Result.Success = false
		userRegisterResult.Result.ErrorCode = utils.ERR0406
		userRegisterResult.Result.ErrorDescription = utils.ERR0406.ToDescription()

		return api.WriteJSON(w, http.StatusOK, userRegisterResult)
	}

	if ok := utils.ValidateEmail(userRegisterArgs.Email); !ok {
		userRegisterResult.Result.Success = false
		userRegisterResult.Result.ErrorCode = utils.ERR0304
		userRegisterResult.Result.ErrorDescription = utils.ERR0304.ToDescription()

		return api.WriteJSON(w, http.StatusOK, userRegisterResult)
	}

	if ok := utils.ValidatePassword(userRegisterArgs.Password); !ok {
		userRegisterResult.Result.Success = false
		userRegisterResult.Result.ErrorCode = utils.ERR0304
		userRegisterResult.Result.ErrorDescription = utils.ERR0304.ToDescription()

		return api.WriteJSON(w, http.StatusOK, userRegisterResult)
	}

	checking_models := database.Model[models.UserModel]{
		Stg: controller.Storage.GetCursor(),
	}
	result, err := checking_models.Get(
		fmt.Sprintf(
			"email = %s OR phone_number = %s",
			userRegisterArgs.Email, userRegisterArgs.PhoneNumber),
	)
	if err != nil || !reflect.DeepEqual(result, models.UserModel{}) {
		userRegisterResult.Result.Success = false
		userRegisterResult.Result.ErrorCode = utils.ERR0404
		userRegisterResult.Result.ErrorDescription = utils.ERR0404.ToDescription()
		userRegisterResult.Result.ErrorException = utils.ExceptionToString(err)

		return api.WriteJSON(w, http.StatusOK, userRegisterResult)
	}

	hashed_password, err := utils.HashPassword(userRegisterArgs.Password)
	if err != nil {
		userRegisterResult.Result.Success = false
		userRegisterResult.Result.ErrorCode = utils.ERR0405
		userRegisterResult.Result.ErrorDescription = utils.ERR0405.ToDescription()
		userRegisterResult.Result.ErrorException = utils.ExceptionToString(err)

		return api.WriteJSON(w, http.StatusOK, userRegisterResult)
	}

	to_register := new(models.UserModel)
	to_register.Id = utils.NewID()
	to_register.FirstName = userRegisterArgs.FirstName
	to_register.LastName = userRegisterArgs.LastName
	to_register.Email = userRegisterArgs.Email
	to_register.PhoneNumber = userRegisterArgs.PhoneNumber
	to_register.Password = hashed_password

	if err := checking_models.Insert(*to_register); err != nil {
		userRegisterResult.Result.Success = false
		userRegisterResult.Result.ErrorCode = utils.ERR0407
		userRegisterResult.Result.ErrorDescription = utils.ERR0407.ToDescription()
		userRegisterResult.Result.ErrorException = utils.ExceptionToString(err)

		return api.WriteJSON(w, http.StatusOK, userRegisterResult)
	}

	userRegisterResult.Result.Success = true
	userRegisterResult.Result.ErrorCode = ""
	userRegisterResult.Result.ErrorDescription = ""

	return api.WriteJSON(w, http.StatusOK, userRegisterResult)
}
