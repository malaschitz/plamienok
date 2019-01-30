package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	rnd "math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/oklog/ulid"
)

var emailCheck *regexp.Regexp

func LogErr(err error) {
	if err != nil {
		log.Info(err)
	}
}

func PanicErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func IsEmail(email string) bool {
	if emailCheck == nil {
		emailCheck = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	}
	return emailCheck.MatchString(email)
}

func UUID() (uuid string) {
	u := ulid.MustNew(ulid.Now(), rand.Reader).String()
	uuid = u[:5] + "-" + u[5:10] + "-" + u[10:14] + "-" + u[14:18] + "-" + u[18:22] + "-" + u[22:]
	return
}

func RandomCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(rnd.Intn(10))
	}
	return code
}

func JsonPrint(o interface{}) {
	j, err := json.MarshalIndent(o, "#", "   ")
	if err != nil {
		fmt.Println("###", err)
	} else {
		fmt.Println(string(j))
	}
}

func init() {
	rnd.Seed(time.Now().UnixNano())
}
