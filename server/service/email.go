package service

import (
	mailgun "github.com/mailgun/mailgun-go"
	"github.com/malaschitz/plamienok/server/constants"
	"github.com/matcornic/hermes/v2"
)

var mg mailgun.Mailgun
var hrms hermes.Hermes

func SendCode(email, name, code6 string) (err error) {
	e := hermes.Email{
		Body: hermes.Body{
			Name:         name,
			FreeMarkdown: hermes.Markdown("There is code for login to voca.ninja:\n\n**" + code6 + "**"),
			Greeting:     "Dobrý deň",
			Signature:    "S pozdravom",
		},
	}
	emailBody, err := hrms.GenerateHTML(e)
	if err != nil {
		return err
	}
	txt, err := hrms.GeneratePlainText(e)

	message := mg.NewMessage(
		constants.AdminEmail,
		"code for a new password",
		txt,
		email)
	message.SetHtml(emailBody)
	_, _, err = mg.Send(message)
	if err != nil {
		return err
	}
	return
}

func InitEmail() {
	mg = mailgun.NewMailgun(
		constants.MailgunDomain,
		constants.MailgunApikey,
		constants.MailgunPublicApiKey)

	hrms = hermes.Hermes{
		Product: hermes.Product{
			Name:        constants.AppName,
			Link:        constants.ServerName,
			Logo:        "https://www.plamienok.sk/img/logo.png?v=0.2",
			Copyright:   "© 2018 https://github.com/malaschitz/plamienok",
			TroubleText: "If the {ACTION}-button is not working for you, just copy and paste the URL below into your web browser.",
		},
	}
}

//Created by Richard Malaschitz
//26/10/2018 00:07
//Copyright (c) 2018. All Rights Reserved.
