package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/fxivan/microservicio/auth/pkg/functions"
	"github.com/fxivan/microservicio/auth/pkg/models"
	"github.com/fxivan/microservicio/auth/pkg/response"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	var m models.UserSignup

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		app.errorLog.Println(err)
		responseError := &response.Response{
			Status:  false,
			Message: "Error al decodificar el json",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	responseInsert, status := app.users.RegisterUser(&m)

	if status == false {
		responseError := &response.Response{
			Status:  false,
			Message: responseInsert,
			Code:    400,
		}
		app.errorLog.Println(err)
		response.HttpResponseError(w, responseError)
		return
	}

	responseSucc := &response.Response{
		Status:  true,
		Message: responseInsert,
		Code:    200,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseSucc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (app *application) signin(w http.ResponseWriter, r *http.Request) {
	var m models.UserLogin

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	userSingIn, err := app.users.FindUsername(m.Username)
	if err != nil {
		app.errorLog.Println(err)
		responseError := &response.Response{
			Status:  false,
			Message: "Necessary to register first",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}
	result := functions.CheckPasswordMatch(userSingIn.Password, m.Password)
	if result == false {

		resposeError := &response.Response{
			Status:  false,
			Message: "Error, contraseña o usuario incorrecto",
			Code:    400,
		}
		response.HttpResponseError(w, resposeError)
		return
	}

	appEnv, err := godotenv.Read(".env")
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	JWTExpirationMs := appEnv["JWTExpirationMs"]
	JWTSecret := appEnv["JWTSecret"]

	expireTimeMs, _ := strconv.Atoi(JWTExpirationMs)

	type JWTCustomClaims struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		ID       string `bson:"_id"`
		jwt.RegisteredClaims
	}

	claims := JWTCustomClaims{
		Username: userSingIn.Username,
		Email:    userSingIn.Email,
		ID:       userSingIn.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireTimeMs) * time.Millisecond)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtString, err := token.SignedString([]byte(JWTSecret))

	if err != nil {
		app.errorLog.Println(err)
		return
	}

	response := &response.Response{
		Status:  true,
		Message: jwtString,
		Code:    200,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
