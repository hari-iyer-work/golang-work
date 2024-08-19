package apis

import (
	"encoding/json"  
	"net/http"
)

type apicall interface{ 
	makeCall(url String) json
	unMarshalJson(body json) apicallStruct any
}

type CallToApi struct{
	url String 
	makeCall(url String) json 
}

func (p CallToApi) makeCall(url String) (json,error) {
response, err :=http.get()
if err!=nil{
	return nil, err
}
if response.StatusCode!= http.StatusOK {

}
 return 
 defer response.Body.Close()
}




