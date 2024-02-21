package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fxivan/microservicio/media/pkg/models"
	"github.com/fxivan/microservicio/media/pkg/response"
	"github.com/google/uuid"
)

func (app *application) uploadImg(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	fileName := r.FormValue("fileName")
	fmt.Print(fileName)

	if err != nil {
		app.errorLog.Println(err)
		responseError := &response.Response{
			Status:  false,
			Message: "Error en la request",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}
	defer file.Close()

	if !strings.HasPrefix(handler.Header.Get("Content-type"), "image/") {
		app.errorLog.Println("El contenido no es valido")
		responseError := &response.Response{
			Status:  false,
			Message: "El contenido no es valido",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	//Medidas de Seguridad
	/*
		1. Tamaño de Archivo
		2. Restringuir extensiones a ".jpg", ".jpeg", ".png"
		3. Generar nombre aleatorios
	*/

	const maxFileSize = 10 << 20
	if handler.Size > maxFileSize {
		app.errorLog.Println("El tamaño del archivo excede el límite permitido")
		responseError := &response.Response{
			Status:  false,
			Message: "El tamaño del archivo excede el límite permitido",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	switch filepath.Ext(handler.Filename) {
	case ".jpg", ".jpeg", ".png":
		break
	default:
		app.errorLog.Println("El formato del archivo no es valido")
		responseError := &response.Response{
			Status:  false,
			Message: "El formato del archivo no es valido",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}
	randomName := uuid.New().String() + filepath.Ext(handler.Filename)
	f, err := os.Create(filepath.Join("uploads", randomName))
	if err != nil {
		app.errorLog.Println("Error al guardar")
		responseError := &response.Response{
			Status:  false,
			Message: "Error al guardar",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		app.errorLog.Println("Error al guardar")
		responseError := &response.Response{
			Status:  false,
			Message: "Error al guardar",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	ID, ok := r.Context().Value("ID").(string)

	if !ok {
		app.errorLog.Println(ok)
		responseError := &response.Response{
			Status:  false,
			Message: "Error interno",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	structJSONImg := &models.ModelPhoto{
		UserId:    ID,
		NameImg:   randomName,
		Size:      handler.Size,
		CreatedAt: time.Now(),
	}

	resUploadImage, status := app.photos.UploadImage(structJSONImg)

	if status == false {
		app.errorLog.Println(resUploadImage)
		responseError := &response.Response{
			Status:  false,
			Message: "Error al cargar la imagen",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	_, _ = io.WriteString(w, "File "+randomName+" Uploaded successfully")
}

func (app *application) uploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("document")
	if err != nil {
		app.errorLog.Println("Error al obtener el documento")
		responseError := &response.Response{
			Status:  false,
			Message: "Error al cargar el contrato",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}
	defer file.Close()

	if !strings.HasPrefix(handler.Header.Get("Content-Type"), "application/") {
		app.errorLog.Println("Error con el content-type del file")
		responseError := &response.Response{
			Status:  false,
			Message: "Error al cargar el contrato",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	const maxFileSize = 10 << 20

	if handler.Size > maxFileSize {
		app.errorLog.Println("Tamaño del file no permitido")
		responseError := &response.Response{
			Status:  false,
			Message: "Tamaño del file no permitido",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	switch filepath.Ext(handler.Filename) {
	case ".pdf":
		break
	default:
		app.errorLog.Println("El formato del archivo no es valido")
		responseError := &response.Response{
			Status:  false,
			Message: "El formato del archivo no es valido",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	randomName := uuid.New().String() + filepath.Ext(handler.Filename)
	f, err := os.Create(filepath.Join("files", randomName))
	if err != nil {
		app.errorLog.Println("El al cargar el archivo")
		responseError := &response.Response{
			Status:  false,
			Message: "El al cargar el archivo",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		app.errorLog.Println("El al cargar el archivo")
		responseError := &response.Response{
			Status:  false,
			Message: "El al cargar el archivo",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	ID, ok := r.Context().Value("ID").(string)
	if !ok {
		app.errorLog.Println("Error al obtener el ID del usuario")
		responseError := &response.Response{
			Status:  false,
			Message: "ERROR INTERNO",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	structJSONFile := &models.ModelFile{
		UserId:    ID,
		NameFile:  randomName,
		Size:      handler.Size,
		CreatedAt: time.Now(),
	}

	_, status := app.files.AddFile(structJSONFile)
	if status == false {
		app.errorLog.Println("El al cargar el archivo")
		responseError := &response.Response{
			Status:  false,
			Message: "El al cargar el archivo",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	_, _ = io.WriteString(w, "File "+randomName+" Uploaded successfully")
}
