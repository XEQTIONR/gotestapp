package mail

import (
	"fmt"
	"net/smtp"
	"os"
)

func Send(toAddress, mailSubject, body string) {
	fromAddress := getSenderAddress()
	fromName := getSenderName()
	password := getSenderPassword()
	host := getHost()
	to := []string{toAddress}
	port := getPort()

	content := []byte(fmt.Sprintf("To: %s\r\nFrom: %s<%s>\r\nSubject: %s\r\n\r\n %s",
		toAddress, fromName, fromAddress, mailSubject, body))

	auth := smtp.PlainAuth("", fromAddress, password, host)
	sendAddress := fmt.Sprintf("%v:%v", host, port)
	err := smtp.SendMail(sendAddress, auth, fromAddress, to, content)

	if err != nil {
		fmt.Printf("Error Sending Mail: %v\n", err)
	}
}

func getSenderAddress() string {
	return os.Getenv("EMAIL_SENDER_ADDRESS")
}

func getSenderName() string {
	return os.Getenv("EMAIL_SENDER_NAME")
}

func getSenderPassword() string {
	return os.Getenv("EMAIL_APP_PASSWORD")
}

func getHost() string {
	return os.Getenv("EMAIL_SMTP_HOST")
}

func getPort() string {
	return os.Getenv("EMAIL_SMTP_PORT")
}
