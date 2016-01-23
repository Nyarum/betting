package betting

import (
	"encoding/json"
	"errors"

	"github.com/andelf/go-curl"
)

type Session struct {
	SessionToken string
	LoginStatus  string
}

func (b *Betting) GetSession(pemCert, keyCert, login, password string) error {
	var (
		session *Session = &Session{}
		err     error
		easy    *curl.CURL = curl.EasyInit()
	)
	defer easy.Cleanup()

	if easy != nil {
		easy.Setopt(curl.OPT_URL, CertURL)
		easy.Setopt(curl.OPT_POST, true)
		easy.Setopt(curl.OPT_VERBOSE, false)

		easy.Setopt(curl.OPT_SSLCERT, pemCert)
		easy.Setopt(curl.OPT_SSLKEY, keyCert)
		easy.Setopt(curl.OPT_SSL_VERIFYHOST, 2)
		easy.Setopt(curl.OPT_SSL_VERIFYPEER, 1)
		easy.Setopt(curl.OPT_HTTPHEADER, []string{
			"Content-Type: application/x-www-form-urlencoded",
			"X-Application: " + b.ApiKey,
		})

		easy.Setopt(curl.OPT_POSTFIELDS, `username=`+login+`&password=`+password)
		bodyFunc := func(buf []byte, userdata interface{}) bool {
			err = json.Unmarshal(buf, session)

			return true
		}

		easy.Setopt(curl.OPT_WRITEFUNCTION, bodyFunc)

		err = easy.Perform()
	}

	if session != nil {
		switch session.LoginStatus {
		case "SUCCESS":
			b.SessionKey = session.SessionToken
		case "":
			err = errors.New("Error in certificates")
		default:
			err = errors.New(session.LoginStatus)
		}
	}

	return err
}
