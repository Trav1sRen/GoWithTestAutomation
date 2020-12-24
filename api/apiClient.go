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

func (*RequestClient) sendSoapRequest(url string, headers map[string]string, requestBody string) error {
	log.Print("*********************** REQUEST START ***********************")
	log.Printf("POST to <%s>", url)
	log.Printf("Headers: {%s}", utils.MapToString(headers))
	log.Print("Request Body: \n", requestBody)

	res, err := http.Post(url, "application/soap+xml; charset=utf-8", strings.NewReader(requestBody))
	if err != nil {
		log.Fatal("Http post error: ", err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("SOAP request failed with the status: %d", res.StatusCode)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error when reading the response body stream: ", err)
		return err
	}

	log.Print("Response Body: \n", string(data))
	return nil
}

func (*RequestClient) sendRestRequest() {

}
