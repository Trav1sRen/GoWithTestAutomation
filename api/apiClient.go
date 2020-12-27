package api

import (
	"GoWithTestAutomation/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type RequestClient struct {
	resBody    string
	statusCode int
}

type SOAPClient struct {
	RequestClient
}

type RESTClient struct {
	RequestClient
}

func (sc *SOAPClient) DispatchReq(url string, headers map[string]string, requestBody string) (err error) {
	log.Print("*********************** REQUEST START ***********************")
	log.Printf("POST to <%s>", url)
	log.Printf("Headers: {%s}", utils.MapToString(headers))
	log.Print("Request Body: \n", requestBody)

	var res *http.Response
	if res, err = http.Post(url, "application/soap+xml; charset=utf-8", strings.NewReader(requestBody)); err != nil {
		return
	}
	defer res.Body.Close()

	sc.statusCode = res.StatusCode
	log.Printf("Response status: <%d>", res.StatusCode)

	var data []byte
	if data, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}

	resBody := string(data)
	sc.resBody = resBody
	log.Print("Response Body: \n", resBody)
	return
}

func (rc *RESTClient) DispatchReq() {

}
