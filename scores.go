package betting

import (
	"github.com/valyala/fasthttp"
	"github.com/pquerna/ffjson/ffjson"
	"fmt"
)

type Scores struct {
	*Client
}

type ScoresResponse struct {
	Result interface{} `json:"result,omitempty"`
	Error  ScoresError `json:"error,omitempty"`
}

type ScoresError struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (e ScoresError) Error() string {
	return fmt.Sprintf("Error with code %d and message %s", e.Code, e.Message)
}

func (s *Scores) Request(reqStruct interface{}, method ScoresListMethod, params *ScoresParams) (err error) {

	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()

	req.SetRequestURI(ScoresURL)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Application", s.ApiKey)
	req.Header.Set("X-Authentication", s.SessionKey)
	req.Header.SetMethod("POST")

	reqBody, err := ffjson.Marshal(&map[string]interface{}{
		"jsonrpc": "2.0",
		"method": "ScoresAPING/v1.0/" + method,
		"params": &params,
	})
	if err != nil {
		return
	}

	req.SetBody(reqBody)

	err = fasthttp.Do(req, res)
	if err != nil {
		return
	}

	scoresResponse := ScoresResponse{Result: reqStruct}

	err = ffjson.Unmarshal(res.Body(), &scoresResponse)
	if err != nil {
		return
	}

	if scoresResponse.Error.Code != 0 {
		err = scoresResponse.Error
	}

	return
}
