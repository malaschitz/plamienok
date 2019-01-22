package constants

import (
	"errors"
	"flag"
	"fmt"
	"time"

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

var TokenDuration = time.Hour * 4

var ERR_NO_USER = errors.New("Používateľ s daným emailom neexistuje.")

var wasInit = false

func InitConst() {
	if wasInit {
		return
	}

	// defaults
	viper.SetDefault("server.name", ServerName)
	viper.SetDefault("server.port", ServerPort)
	viper.SetDefault("server.Debug", ServerDebug)
	viper.SetDefault("database.file", DatabaseFile)
	viper.SetDefault("admin.email", AdminEmail)
	viper.SetDefault("admin.password", AdminPassword)

	// read
	viper.SetConfigType("properties")
	configName := flag.String("properties", "settings.properties", "File with settings properties.")
	flag.Parse()
	viper.SetConfigFile(*configName)

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	ServerName = viper.GetString("server.name")
	ServerPort = viper.GetString("server.port")
	ServerDebug = viper.GetBool("server.debug")
	DatabaseFile = viper.GetString("database.file")
	AdminEmail = viper.GetString("admin.email")
	AdminPassword = viper.GetString("admin.password")
	MailgunDomain = viper.GetString("mailgun.domain")
	MailgunApikey = viper.GetString("mailgun.apikey")
	MailgunPublicApiKey = viper.GetString("mailgun.publicapikey")

	if ServerName == "" || ServerPort == "" || DatabaseFile == "" || AdminEmail == "" ||
		AdminPassword == "" || MailgunDomain == "" || MailgunApikey == "" || MailgunPublicApiKey == "" {
		panic("Bad config")
	}

	// setup
	if ServerDebug {
		log.SetLevel(log.DEBUG)
	}

	wasInit = true

}
