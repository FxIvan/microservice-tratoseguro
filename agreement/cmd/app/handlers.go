package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fxivan/microservicio/agreement/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (app *application) searchCTPY(w http.ResponseWriter, r *http.Request) {

	var m models.SearchUser
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		panic(err)
		return
	}

	cursor, err := app.users.C.Find(r.Context(), bson.M{})
	if err != nil {
		fmt.Print(err)
		panic(err)
		return
	}
	defer cursor.Close(r.Context())

	var users []bson.M
	if err := cursor.All(r.Context(), &users); err != nil {
		fmt.Print(err)
		panic(err)
		return
	}

	fmt.Println(users)

}
