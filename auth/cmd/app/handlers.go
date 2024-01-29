package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fxivan/microservicio/auth/pkg/functions"
	"github.com/fxivan/microservicio/auth/pkg/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func (app *application) insert(w http.ResponseWriter, r *http.Request) {
	var m models.UserSignup
	fmt.Println("Insertando Usuario", m)
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		app.errorLog.Println(err)
	}

	insertResult, err := app.users.InsertRegisterUser(&m)
	if err != nil {
		app.errorLog.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(insertResult)

}

func (app *application) signin(w http.ResponseWriter, r *http.Request) {
	var m models.UserLogin

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.errorLog.Println(err)
	}
	userSingIn, err := app.users.FindUserEmail(m.Username)

	result := functions.CheckPasswordMatch(userSingIn.Password, m.Password)
	if result == false {
		http.Error(w, "Contraseña incorrecta", 405)
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

	response := map[string]interface{}{
		"status":  "success",
		"message": jwtString,
		"code":    200,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
