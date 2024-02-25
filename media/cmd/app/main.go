package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fxivan/microservicio/media/pkg/models/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	photos   *mongodb.PhotosModel
	files    *mongodb.FilesModel
}

func main() {

	//Generamos Logs para mostrar en la Shell
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mongo_user := os.Getenv("MONGO_USER")
	mongo_password := os.Getenv("MONGO_PASSWORD")

	//Configuramos las variables para levantar el Server y DB
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4000, "HTTP server network port")
	mongoURI := flag.String("mongoURI", fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin", mongo_user, mongo_password, "localhost", 27018, "media"), "Mongo Connection Uri")
	mongoDB := flag.String("mongodb", "media", "Photos and Files Into MONGO")
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
		errorLog.Println(err)
		return
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
		photos: &mongodb.PhotosModel{
			C: client.Database(*mongoDB).Collection("photos"),
		},
		files: &mongodb.FilesModel{
			C: client.Database(*mongoDB).Collection(("files")),
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
	errorLog.Println(err)

}
