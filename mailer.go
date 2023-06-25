package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"gopkg.in/gomail.v2"
)

var Config struct {
	Host     string
	Port     int
	User     string
	Password string
}

func ReportScheduleAppointment() error {

	HTML, err := downloadHTMLFile("https://email.zkaia.com/temps/filler.html")
	if err != nil {
		return err
	}

	audience, err := ProcessAudience("https://email.zkaia.com/temps/audience.csv")
	if err != nil {
		return err
	}

	for _, person := range audience {
		mailer := gomail.NewMessage()
		htm, _ := generateHTMLWithUserData(string(HTML), *person)
		mailer.SetBody("text/html", htm)
		mailer.SetHeader("From", Config.User)

		mailer.SetHeader("To", person.Email)
		mailer.SetHeader("Subject", fmt.Sprintf("%s %s- ZKAIA send you an email", person.First, person.Last))

		dialer := gomail.NewDialer(Config.Host, Config.Port, Config.User, Config.Password)

		if err := dialer.DialAndSend(mailer); err != nil {
			return err
		}

		mailer.Reset()
	}

	return nil

}

func generateHTMLWithUserData(htmlContent string, data Audience) (string, error) {
	tmpl, err := template.New("htmlTemplate").Parse(htmlContent)
	if err != nil {
		return "", err
	}

	var renderedHTML strings.Builder
	err = tmpl.Execute(&renderedHTML, data)
	if err != nil {
		return "", err
	}

	return renderedHTML.String(), nil
}

func downloadHTMLFile(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
