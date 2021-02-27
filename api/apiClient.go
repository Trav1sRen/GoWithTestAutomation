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
	BaseURL      string
	InsecureSkip bool
}

type auth struct {
	Username, Password string
}

type SOAPClient struct {
	RequestClient
}

type RESTClient struct {
	RequestClient
}

func (sc *SOAPClient) DispatchReq(so *SOAPObject) (err error) {
	reqURL := sc.BaseURL + so.Endpoint

	var req *http.Request
	if req, err = http.NewRequest("POST", reqURL, strings.NewReader(so.SOAPBody)); err != nil {
		return
	}

	if sc.auth != struct{ Username, Password string }{} {
		req.SetBasicAuth(sc.Username, sc.Password)
	}

	for k, v := range so.Headers {
		req.Header.Add(k, v.(string))
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: sc.InsecureSkip,
		},
	}

	client := &http.Client{Transport: tr}

	log.Print("*********************** REQUEST START ***********************")
	log.Printf("POST to <%s>", reqURL)
	log.Printf("Headers: {%s}", utils.MapToString(so.Headers))
	log.Print("Request Body: \n", so.SOAPBody)

	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return
	}
	defer res.Body.Close()

	so.StatusCode = res.StatusCode
	log.Printf("Response status: <%d>", res.StatusCode)

	var data []byte
	if data, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}

	resBody := string(data)
	so.ResStr = resBody
	so.ResMap, err = utils.XML2Map(resBody)
	if err != nil {
		return
	}

	log.Print("Response Body: \n", resBody)
	return
}

func (rc *RESTClient) DispatchReq() {

}
