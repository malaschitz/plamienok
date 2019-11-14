module github.com/malaschitz/plamienok

replace gopkg.in/russross/blackfriday.v2 v2.0.1 => github.com/russross/blackfriday/v2 v2.0.1

require (
	github.com/asdine/storm v2.1.2+incompatible
	github.com/facebookgo/ensure v0.0.0-20160127193407-b4ab57deab51 // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/subset v0.0.0-20150612182917-8dac2c3c4870 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gobuffalo/envy v1.6.11 // indirect
	github.com/google/uuid v1.1.0 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.2.8
	github.com/mailgun/mailgun-go v1.1.1
	github.com/matcornic/hermes v1.2.0
	github.com/matcornic/hermes/v2 v2.0.1
	github.com/oklog/ulid v1.3.1
	github.com/pkg/errors v0.8.0
	github.com/spf13/viper v1.2.1
	github.com/texttheater/golang-levenshtein v0.0.0-20180516184445-d188e65d659e
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v0.0.0-20170224212429-dcecefd839c4 // indirect
	go.etcd.io/bbolt v1.3.0
	golang.org/x/crypto v0.0.0-20181203042331-505ab145d0a9
	golang.org/x/sys v0.0.0-20181213081344-73d4af5aa059 // indirect
	golang.org/x/text v0.3.0
	google.golang.org/api v0.0.0-20181213000619-1327b224df06
	google.golang.org/appengine v1.3.0 // indirect
)

go 1.13
