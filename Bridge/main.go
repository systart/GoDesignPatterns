package main

import (
	"fmt"
)

func main() {
	var mb MessagingBridge
	mb.SendMessageByBridge("Test Sms.", SmsSender{})
	mb.SendMessageByBridge("Test Mail.", MailSender{})
}

type MessagingBridge struct{}

func (mb MessagingBridge) SendMessageByBridge(message string, sender IMessageSender) {
	sender.SendMessage(message)
}

type IMessageSender interface {
	SendMessage(message string)
}

type SmsSender struct{}

func (sender SmsSender) SendMessage(message string) {
	fmt.Println("Message: ", message, "\n Sended by SmsSender")
}

type MailSender struct{}

func (sender MailSender) SendMessage(message string) {
	fmt.Println("Message: ", message, "\n Sended by MailSender")
}
