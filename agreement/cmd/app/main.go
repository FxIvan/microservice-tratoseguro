package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	mongodb "github.com/fxivan/microservicio/agreement/pkg/models/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	errorLog     *log.Logger
	infoLog      *log.Logger
	ctpyFileName *mongodb.CtpyFileNameModel
	agreement    *mongodb.AgreementModel
	users        *mongodb.UserSignupModel
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//mongo_user := os.Getenv("MONGO_USER")
	//mongo_password := os.Getenv("MONGO_PASSWORD")

	mongo_user := "admtratoseguro210814"
	mongo_password := "LkdU7ZDADARiFEtZiKJUjUeg5Swfyq9dA7qwkqjerkpQZwEvUs"

	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4000, "HTTP server network port")
	mongoURI := flag.String("mongoURI", fmt.Sprintf("mongodb://%s:%s@host.docker.internal:%d/%s?authSource=admin", mongo_user, mongo_password, 27019, "agreement"), "Mongo Connection Uri")
	//mongoURIuser := flag.String("mongoURIuser", fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin", mongo_user, mongo_password, "0.0.0.0", 27018, "user"), "Mongo Connection Uri for User Database")
	mongoURIuser := flag.String("mongoURIuser", fmt.Sprintf("mongodb://%s:%s@host.docker.internal:%d/%s?authSource=admin", mongo_user, mongo_password, 27018, "user"), "Mongo Connection Uri for User Database")
	mongoDBctpyFileName := flag.String("nameDBctpyFilename", "CounterpartFilename", "Name DB")
	mongoDBagrDefinition := flag.String("nameDBagrDefinition", "AgreementDefinition", "Name DB")
	mongoDBuser := flag.String("mongoDatabase", "users", "MongoDB database")
	flag.Parse()

	infoLog.Println("Variable configuration | Port | URI | NameDB")
	infoLog.Println("DB init")

	co := options.Client().ApplyURI(*mongoURI)
	client, err := mongo.NewClient(co)
	if err != nil {
		errorLog.Println(err)
		return
	}

	coUser := options.Client().ApplyURI(*mongoURIuser)
	clientUser, err := mongo.NewClient(coUser)
	if err != nil {
		errorLog.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = clientUser.Connect(ctx)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			errorLog.Println(err)
			return
		}

		if err := clientUser.Disconnect(ctx); err != nil {
			errorLog.Println(err)
			return
		}

	}()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		ctpyFileName: &mongodb.CtpyFileNameModel{
			C: client.Database(*mongoDBctpyFileName).Collection("counterpart_filename"),
		},
		agreement: &mongodb.AgreementModel{
			C: client.Database(*mongoDBagrDefinition).Collection("agreement_definition"),
		},
		users: &mongodb.UserSignupModel{
			C: clientUser.Database(*mongoDBuser).Collection("users"),
		},
	}

	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)

	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", serverURI)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
