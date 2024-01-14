package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	mongodb "github.com/fxivan/microservicio/transactions/pkg/models/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	transactions *mongodb.TransactionModel
}


func main() {
	//Linea de comando
	//Ejemplo de como se ejecuta los flag
	//go run main.go -serverAddr=":8080"
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4000, "HTTP server network port")
	mongoURI := flag.String("mongoURI", "mongodb://root:secret@localhost:27017/transactions?authSource=admin", "MongoDB connection URI")
	mongoDatabase := flag.String("mongoDatabase", "transactions", "MongoDB database")	
	//enableCredentials := flag.Bool("enableCredentials", false, "Enable HTTP basic authentication")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	co := options.Client().ApplyURI(*mongoURI)

	client , err := mongo.NewClient(co)
	if err != nil{
		errorLog.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil{
		errorLog.Fatal(err)
	}

	defer func(){
		if err = client.Disconnect(ctx); err != nil{
			errorLog.Fatal(err)
		}
	}()
 
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		transactions: &mongodb.TransactionModel{
			C: client.Database(*mongoDatabase).Collection("transactions"),
		},
	}

	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	srv := &http.Server{
		Addr: serverURI,
		ErrorLog: errorLog,
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", serverURI)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}