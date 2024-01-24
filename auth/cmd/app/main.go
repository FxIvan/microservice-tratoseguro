package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	mongodb "github.com/fxivan/microservicio/auth/pkg/models/mongodb"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	users    *mongodb.UserSignupModel
}

func main() {
	appEnv, err := godotenv.Read(".env")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if err != nil {
		log.Println("No .env file")
	}
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4000, "HTTP server network port")
	mongoURI := flag.String("mongoURI", fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin", appEnv["MONGO_USER"], appEnv["MONGO_PASSWORD"], "localhost", 27018, "user"), "Mongo Connection Uri")
	mondoDatabase := flag.String("mongoDatabase", "users", "MongoDB database")
	flag.Parse()

	co := options.Client().ApplyURI(*mongoURI)

	client, err := mongo.NewClient(co)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		users: &mongodb.UserSignupModel{
			C: client.Database(*mondoDatabase).Collection("users"),
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
