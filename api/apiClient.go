package api

import (
	"GoWithTestAutomation/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type RequestClient struct {
	responseBody string
	statusCode   int
	verifySsl    bool
}

func (rc *RequestClient) sendSoapRequest(url string, headers map[string]string, requestBody string) (err error) {
	log.Print("*********************** REQUEST START ***********************")
	log.Printf("POST to <%s>", url)
	log.Printf("Headers: {%s}", utils.MapToString(headers))
	log.Print("Request Body: \n", requestBody)

	var res *http.Response
	res, err = http.Post(url, "application/soap+xml; charset=utf-8", strings.NewReader(requestBody))
	if err != nil {
		log.Print("Http post error: ", err)
		return
	}
	defer res.Body.Close()

	rc.statusCode = res.StatusCode
	log.Printf("Response status: <%d>", res.StatusCode)

	var data []byte
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print("Error when reading the response body stream: ", err)
		return
	}

	resBody := string(data)
	rc.responseBody = resBody
	log.Print("Response Body: \n", resBody)
	return
}

func (*RequestClient) sendRestRequest() {

}
