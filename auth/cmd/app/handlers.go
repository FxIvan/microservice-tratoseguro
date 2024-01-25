package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fxivan/microservicio/auth/pkg/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

	err = bcrypt.CompareHashAndPassword([]byte(userSingIn.Password), []byte(m.Password))
	if err != nil {
		app.errorLog.Println("Contraseña o usuario incorrecto")
		return
	}

	JWTExpirationMs := "86400000"
	JWTSecret := "R1BYcTVXVGNDU2JmWHVnZ1lnN0FKeGR3cU1RUU45QXV4SDJONFZ3ckhwS1N0ZjNCYVkzZ0F4RVBSS1UzRENwRw=="

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

	fmt.Println(jwtString)

}
