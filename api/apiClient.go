package api

import (
	"GoWithTestAutomation/utils"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type RequestClient struct {
	auth
	Url, ResBody string
	StatusCode   int
	InsecureSkip bool
}

type auth struct {
	Username, Password string
}

type SOAPClient struct {
	RequestClient
}

type RESTClient struct {
	Method string
	RequestClient
}

func (sc *SOAPClient) DispatchReq(headers map[string]string, soapAction, soapBody string) (err error) {
	var req *http.Request
	if req, err = http.NewRequest("POST", sc.Url, strings.NewReader(soapBody)); err != nil {
		return
	}

	if sc.auth != struct{ Username, Password string }{} {
		req.SetBasicAuth(sc.Username, sc.Password)
	}

	headers["SOAPAction"] = soapAction
	// TODO: Content-Type is different between SOAP version 1.1/1.2
	headers["Content-Type"] = "text/xml; charset=\"utf-8\""
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: sc.InsecureSkip,
		},
	}

	client := &http.Client{Transport: tr}

	log.Print("*********************** REQUEST START ***********************")
	log.Printf("POST to <%s>", sc.Url)
	log.Printf("Headers: {%s}", utils.MapToString(headers))
	log.Print("Request Body: \n", soapBody)

	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return
	}
	defer res.Body.Close()

	sc.StatusCode = res.StatusCode
	log.Printf("Response status: <%d>", res.StatusCode)

	var data []byte
	if data, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}

	resBody := string(data)
	sc.ResBody = resBody
	log.Print("Response Body: \n", resBody)
	return
}

func (rc *RESTClient) DispatchReq() {

}
