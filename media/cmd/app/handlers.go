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
	fileFront, handlerFront, errFront := r.FormFile("frontDocument")
	fileBack, handlerBack, errBack := r.FormFile("backDocument")

	if errFront != nil || errBack != nil {
		app.errorLog.Println(errFront, errBack)
		responseError := &response.Response{
			Status:  false,
			Message: "Error en la request",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}
	defer fileFront.Close()
	defer fileBack.Close()

	if !strings.HasPrefix(handlerFront.Header.Get("Content-type"), "image/") || !strings.HasPrefix(handlerBack.Header.Get("Content-type"), "image/") {
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
	if handlerFront.Size > maxFileSize || handlerBack.Size > maxFileSize {
		app.errorLog.Println("El tamaño del archivo excede el límite permitido")
		responseError := &response.Response{
			Status:  false,
			Message: "El tamaño del archivo excede el límite permitido",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	switch filepath.Ext(handlerFront.Filename) {
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

	switch filepath.Ext(handlerBack.Filename) {
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

	randomNameFront := uuid.New().String() + filepath.Ext(handlerFront.Filename)
	randomNameBack := uuid.New().String() + filepath.Ext(handlerBack.Filename)

	fFront, errFront := os.Create(filepath.Join("uploads", randomNameFront))
	fBack, errBack := os.Create(filepath.Join("uploads", randomNameBack))

	if errFront != nil || errBack != nil {
		app.errorLog.Println("Error al guardar errFront:", errFront)
		app.errorLog.Println("Error al guardar errBack:", errBack)
		responseError := &response.Response{
			Status:  false,
			Message: "Error al guardar",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	defer fFront.Close()
	defer fBack.Close()

	_, errFront = io.Copy(fFront, fileFront)
	_, errBack = io.Copy(fBack, fileBack)
	if errFront != nil || errBack != nil {
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

	structJSONImgFront := &models.ModelPhoto{
		UserId:    ID,
		NameImg:   randomNameFront,
		Field:     "documentFront",
		Size:      handlerFront.Size,
		CreatedAt: time.Now(),
	}

	structJSONImgBack := &models.ModelPhoto{
		UserId:    ID,
		NameImg:   randomNameBack,
		Field:     "documentBack",
		Size:      handlerBack.Size,
		CreatedAt: time.Now(),
	}

	resUploadImageFront, statusFront := app.photos.UploadImage(structJSONImgFront)
	resUploadImageBack, statusBack := app.photos.UploadImage(structJSONImgBack)

	if statusFront == false || statusBack == false {
		app.errorLog.Println(resUploadImageFront, resUploadImageBack)
		responseError := &response.Response{
			Status:  false,
			Message: "Error al cargar la imagen",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

	_, _ = io.WriteString(w, "File "+randomNameFront+" and "+randomNameFront+" Uploaded successfully")
}

func (app *application) uploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("document")
	if err != nil {
		app.errorLog.Println("Error al obtener el documento")
		errorMsg := fmt.Sprintf("Error al cargar el contrato: %s", err)
		responseError := &response.Response{
			Status:  false,
			Message: errorMsg,
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

	_, exist := app.files.FindFile(randomName)
	if exist == true {
		app.errorLog.Println("Vuelve a cargar la image, hubo un error")
		responseError := &response.Response{
			Status:  false,
			Message: "Vuelve a cargar el contrato, hubo un error",
			Code:    400,
		}
		response.HttpResponseError(w, responseError)
		return
	}

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
