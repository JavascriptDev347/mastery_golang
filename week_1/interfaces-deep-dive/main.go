package main

import "fmt"

type Notifier interface {
	Notify(message string) error
}

type EmailSender struct {
	smtpHost string
}

func (e *EmailSender) Notify(message string) error {
	fmt.Printf("Sending email: %s\n", message)
	return nil
}

type SMSSender struct {
	phoneNumber string
}

func (s *SMSSender) Notify(message string) error {
	fmt.Printf("Sending SMS: %s\n", message)
	return nil
}

func notify(notifier Notifier, message string) {
	err := notifier.Notify(message)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Hech qanday "implements" yozilmagan, lekin ikkalasi ham Notifier
func SendAlert(n Notifier, msg string) error {
	return n.Notify(msg)
}
func main() {
	SendAlert(&EmailSender{}, "Server down!")
	SendAlert(&SMSSender{}, "Server down!")
}
