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

type params struct {
	paramName  string
	paramValue string
}
type callToApi struct {
	url string
}

// http calls
func (p callToApi) makeCall(urlString string, queryParams []params, pathParams []params) ([]byte, error) {
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

// grpc calls
