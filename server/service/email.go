package service

import (
	"project/utils"
)

type EmailService struct {
}

func (e *EmailService) EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = utils.EmailTest(subject, body)
	return err
}
