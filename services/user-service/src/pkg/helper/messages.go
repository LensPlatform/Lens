package helper

import (
	"fmt"
)

func WelcomeMessage(firstname, lastname string) string {
	name := fmt.Sprintf("%s %s", firstname, lastname)
	message := "Dear " + name + " \n Thank you for opening a new account on the CUBE platform. We look " +
		"forward to providing you with solutions and support to help you reach your goals" +
		"\n" +
		"\n" + "Thank you for being our customer"
	return message
}
