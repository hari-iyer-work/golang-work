package apis

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type apiCall interface {
	unmarshalJson(body []byte) (interface{}, error)
}

type Params struct {
	paramName  string
	paramValue string
}
type CallToApi struct {
	url string
}

// http calls
func (p CallToApi) MakeCall(urlString string, queryParams []Params, pathParams []Params) ([]byte, error) {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	parameter := url.Values{}

	for _, param := range queryParams {
		parameter.Add(param.paramName, param.paramValue)
	}
	parsedUrl.RawQuery = parameter.Encode()
	// ToDo: add path params
	response, err := http.Get(parsedUrl.String())

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

func (p *Params) SetName(paramName string) {
	p.paramName = paramName

}

func (p *Params) SetValue(paramValue string) {
	p.paramValue = paramValue
}

// grpc calls
