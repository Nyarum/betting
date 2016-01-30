package betting

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

type Session struct {
	SessionToken string
	LoginStatus  string
}

func (b *Betting) GetSession(pemCert, keyCert, login, password string) error {
	var session *Session = &Session{}

	cert, err := tls.LoadX509KeyPair(pemCert, keyCert)
	if err != nil {
		return err
	}

	client := fasthttp.Client{TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}}

	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	req.SetRequestURI(CertURL)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Application", b.ApiKey)
	req.Header.SetMethod("POST")

	bufferString := bytes.NewBuffer([]byte{})
	bufferString.WriteString(`username=`)
	bufferString.WriteString(login)
	bufferString.WriteString(`&password=`)
	bufferString.WriteString(password)

	req.SetBody(bufferString.Bytes())

	err = client.Do(req, resp)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Body(), session)
	if err != nil {
		return err
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

	return nil
}
