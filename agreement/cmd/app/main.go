package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fxivan/microservicio/agreement/pkg/models/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	ctpyFileName  *mongodb.CtpyFileNameModel
	agrDefinition *mongodb.AgrDefinitionModel
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mongo_user := os.Getenv("MONGO_USER")
	mongo_password := os.Getenv("MONGO_PASSWORD")

	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4000, "HTTP server network port")
	mongoURI := flag.String("mongoURI", fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin", mongo_user, mongo_password, "mongo", 27017, "agreement"), "Mongo Connection Uri")
	mongoDBctpyFileName := flag.String("nameDBctpyFilename", "CounterpartFilename", "Name DB")
	mongoDBagrDefinition := flag.String("nameDBagrDefinition", "AgreementDefinition", "Name DB")
	flag.Parse()

	infoLog.Println("Variable configuration | Port | URI | NameDB")
	infoLog.Println("DB init")

	co := options.Client().ApplyURI(*mongoURI)
	client, err := mongo.NewClient(co)
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

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
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
		agrDefinition: &mongodb.AgrDefinitionModel{
			C: client.Database(*mongoDBagrDefinition).Collection("agreement_definition"),
		},
	}

	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)

	// Handler para registrar las solicitudes
	handlerWithLogging := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			infoLog.Printf("Request: %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}

	// Agrega el middleware de logging al router
	router := app.routes()
	router.Use(handlerWithLogging)

	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errorLog,
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", serverURI)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
