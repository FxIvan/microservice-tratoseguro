package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	_, _ = io.WriteString(w, "File "+randomName+" Uploaded successfully")
}
