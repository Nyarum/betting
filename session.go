package betting

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

type Client struct {
	ApiKey     string
	SessionKey string
}

type Session struct {
	SessionToken string
	LoginStatus  ELoginStatus
}

func (c *Client) GetSession(pemCert, keyCert, login, password string) error {
	var session *Session = &Session{}

	cert, err := tls.LoadX509KeyPair(pemCert, keyCert)
	if err != nil {
		return err
	}

	client := fasthttp.Client{TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}}

	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	req.SetRequestURI(CertURL)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Application", c.ApiKey)
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

	switch session.LoginStatus {
	case LS_SUCCESS:
		c.SessionKey = session.SessionToken
	default:
		err = errors.New(string(session.LoginStatus))
	}

	return err
}

type KeepAlive struct {
	Token   string
	Product string
	Status  string
	Error   string
}

// KeepAlive for support connect, session key will available for 20 minutes
func (c *Client) KeepAlive() error {
	var keepAlive *KeepAlive = &KeepAlive{}

	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	req.SetRequestURI(KeepAliveURL)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Application", c.ApiKey)
	req.Header.Set("X-Authentication", c.SessionKey)
	req.Header.SetMethod("POST")

	err := fasthttp.Do(req, resp)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Body(), keepAlive)
	if err != nil {
		return err
	}

	switch keepAlive.Status {
	case "SUCCESS":
		c.SessionKey = keepAlive.Token
	default:
		err = errors.New(keepAlive.Error)
	}

	return err
}
