package service

import (
	"fmt"
	"log"

	mailgun "github.com/mailgun/mailgun-go"
	"github.com/malaschitz/plamienok/server"

	hermes "github.com/matcornic/hermes/v2"
)

var mg mailgun.Mailgun
var hrms hermes.Hermes

func SendCode(email, name, code6 string) (err error) {
	fmt.Println(email, name, code6)
	if server.IsProd {
		message := mg.NewMessage(
			"admin@mail.voca.ninja",
			"code for login",
			code6,
			email)
		e := hermes.Email{
			Body: hermes.Body{
				//Title: "Voca Ninja",
				Name:         name,
				FreeMarkdown: hermes.Markdown("There is code for login to voca.ninja:\n\n**" + code6 + "**"),
				//Actions:   []hermes.Action{{}},
				//Outros:    []string{},
				Greeting:  "Hi",
				Signature: "Sincerely",
			},
		}
		emailBody, err := hrms.GenerateHTML(e)
		if err != nil {
			log.Println(err)
		}
		message.SetHtml(emailBody)
		var resp, id string
		resp, id, err = mg.Send(message)
		if err != nil {
			fmt.Println("Problem with email", err, resp, id)
			log.Println(err)
		}
	}
	return
}

func init() {
	mg = mailgun.NewMailgun(
		server.MailgunDomain,
		server.MailgunApikey,
		server.MailgunPublicApiKey)

	hrms = hermes.Hermes{
		Product: hermes.Product{
			Name:        server.AppName,
			Link:        server.ServerName,
			Logo:        "https://www.plamienok.sk/img/logo.png?v=0.2",
			Copyright:   "© 2018 https://github.com/malaschitz/plamienok",
			TroubleText: "If the {ACTION}-button is not working for you, just copy and paste the URL below into your web browser.",
		},
	}
}

//Created by Richard Malaschitz malaschitz@gmail.com
//26/10/2018 00:07
//Copyright (c) 2018. All Rights Reserved.
