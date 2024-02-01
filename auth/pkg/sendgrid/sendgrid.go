package sendgrid

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmailSengrid(email string, messageText string) {

	appEnv, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	from := mail.NewEmail("Example User", "service.almendra@gmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", email)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(appEnv["SENDGRID_API_KEY"])
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
