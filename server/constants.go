package server

import (
	"flag"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"

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

var wasInit = false

func InitConst() {
	if wasInit {
		return
	}
	configName := flag.String("properties", "settings.properties", "File with settings properties.")
	flag.Parse()

	v := viper.New()
	file, err := os.Open(*configName)
	if err != nil {
		panic(err)
	}

	// defaults
	v.SetDefault("server.name", ServerName)
	v.SetDefault("server.port", ServerPort)
	v.SetDefault("server.Debug", ServerDebug)
	v.SetDefault("database.file", DatabaseFile)
	v.SetDefault("admin.email", AdminEmail)
	v.SetDefault("admin.password", AdminPassword)

	// read
	err = v.ReadConfig(file)
	if err != nil { // Handle errors reading the config file
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

	wasInit = true
}
