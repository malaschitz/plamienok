package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	rnd "math/rand"
	"regexp"
	"strconv"
	"strings"
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
func RemoveDiacritics(s string) string {
	// Replace some common accent characters
	b := bytes.NewBufferString("")
	for _, c := range s {
		// Check transliterations first
		if val, ok := transliterations[c]; ok {
			b.WriteString(val)
		} else {
			b.WriteRune(c)
		}
	}
	return strings.ToLower(b.String())
}

func init() {
	rnd.Seed(time.Now().UnixNano())
}

var transliterations = map[rune]string{
	'À': "A",
	'Á': "A",
	'Â': "A",
	'Ã': "A",
	'Ä': "A",
	'Å': "AA",
	'Æ': "AE",
	'Ç': "C",
	'Č': "C",
	'Ď': "D",
	'È': "E",
	'É': "E",
	'Ê': "E",
	'Ë': "E",
	'Ì': "I",
	'Í': "I",
	'Î': "I",
	'Ï': "I",
	'Ð': "D",
	'Ł': "L",
	'Ľ': "L",
	'Ĺ': "L",
	'Ñ': "N",
	'Ň': "N",
	'Ò': "O",
	'Ó': "O",
	'Ô': "O",
	'Õ': "O",
	'Ö': "O",
	'Ø': "OE",
	'Ŕ': "R",
	'Ř': "R",
	'Š': "S",
	'Ť': "T",
	'Ù': "U",
	'Ú': "U",
	'Ü': "U",
	'Û': "U",
	'Ý': "Y",
	'Þ': "Th",
	'ß': "ss",
	'à': "a",
	'á': "a",
	'â': "a",
	'ã': "a",
	'ä': "a",
	'å': "aa",
	'æ': "ae",
	'ç': "c",
	'č': "c",
	'ď': "d",
	'è': "e",
	'é': "e",
	'ê': "e",
	'ë': "e",
	'ì': "i",
	'í': "i",
	'î': "i",
	'ï': "i",
	'ð': "d",
	'ł': "l",
	'ľ': "l",
	'ĺ': "l",
	'ñ': "n",
	'ń': "n",
	'ň': "n",
	'ò': "o",
	'ó': "o",
	'ô': "o",
	'õ': "o",
	'ō': "o",
	'ö': "o",
	'ø': "oe",
	'ŕ': "r",
	'ř': "r",
	'ś': "s",
	'š': "s",
	'ť': "t",
	'ù': "u",
	'ú': "u",
	'û': "u",
	'ū': "u",
	'ü': "u",
	'ý': "y",
	'þ': "th",
	'ÿ': "y",
	'ż': "z",
	'Œ': "OE",
	'œ': "oe",
	'Ž': "Z",
	'ž': "z",
}
