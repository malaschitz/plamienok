package server

import (
	"fmt"
	"os"

	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"
)

const Version = "0.1"
const AppName = "plamienok"

var DatabaseFile = "plamienok.db"
var ServerName = "localhost"
var ServerPort = "1324"
var ServerDebug = false
var AdminEmail = "admin@mailinator.com"
var AdminPassword = ""

var MailgunDomain = ""
var MailgunApikey = ""
var MailgunPublicApiKey = ""

func InitConst() {
	configName := "settings"
	if len(os.Args) > 1 {
		configName = os.Args[1]
	}

	v := viper.New()
	v.SetConfigName(configName)
	v.AddConfigPath(".")
	v.SetConfigType("properties")

	// defaults
	v.SetDefault("server.name", ServerName)
	v.SetDefault("server.port", ServerPort)
	v.SetDefault("server.Debug", ServerDebug)
	v.SetDefault("database.file", DatabaseFile)
	v.SetDefault("admin.email", AdminEmail)
	v.SetDefault("admin.password", AdminPassword)

	// read
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	ServerName = v.GetString("server.name")
	ServerPort = v.GetString("server.port")
	ServerDebug = v.GetBool("server.debug")
	DatabaseFile = v.GetString("database.file")
	AdminEmail = v.GetString("admin.email")
	AdminPassword = v.GetString("admin.password")
	MailgunDomain = v.GetString("mailgun.domain")
	MailgunApikey = v.GetString("mailgun.apikey")
	MailgunPublicApiKey = v.GetString("mailgun.publicApiKey")

	// setup
	if ServerDebug {
		log.SetLevel(log.DEBUG)
	}
}
