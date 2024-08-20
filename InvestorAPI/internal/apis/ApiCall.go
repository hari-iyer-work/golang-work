package apis

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type ApiCall interface {
	UnmarshalJson(body []byte) (interface{}, error)
}

type CallToApi struct {
	url string
}

func (p CallToApi) MakeCall(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("error: received non-200 status code `%#v`", response)
		log.Println(msg)
		return nil, fmt.Errorf(msg)
	}
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
