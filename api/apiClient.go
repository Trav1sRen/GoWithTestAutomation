package api

import (
	"GoWithTestAutomation/utils"
	"crypto/tls"
	"github.com/bitly/go-simplejson"
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
	if err = sc.dispatchReq(&so.RequestObject, &so.ResponseObject); err != nil {
		return
	}
	if so.ResMap, err = utils.XML2Map(so.ResStr); err != nil {
		return
	}
	return nil
}

func (rc *RESTClient) DispatchReq(ro *RESTObject) (err error) {
	if err = rc.dispatchReq(&ro.RequestObject, &ro.ResponseObject); err != nil {
		return
	}

	switch strings.ToUpper(ro.DataFormat) {
	case "XML":
		if ro.ResMap, err = utils.XML2Map(ro.ResStr); err != nil {
			return
		}
		break
	case "JSON":
		var j *simplejson.Json
		if j, err = utils.Str2JSON(ro.ResStr); err != nil {
			return
		}
		if ro.ResMap, err = j.Map(); err != nil {
			return
		}
	}

	return nil
}

func (rc *RequestClient) dispatchReq(reqObj *RequestObject, resObj *ResponseObject) (err error) {
	reqURL := rc.BaseURL + reqObj.Endpoint

	var req *http.Request
	if req, err = http.NewRequest(reqObj.Method, reqURL, strings.NewReader(reqObj.RequestBody)); err != nil {
		return
	}

	if rc.auth != struct{ Username, Password string }{} {
		req.SetBasicAuth(rc.Username, rc.Password)
	}

	for k, v := range reqObj.Headers {
		req.Header.Add(k, v.(string))
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: rc.InsecureSkip,
		},
	}

	client := &http.Client{Transport: tr}

	log.Print("*********************** REQUEST START ***********************")
	log.Printf("POST to <%s>", reqURL)
	log.Printf("Headers: {%s}", utils.MapToString(reqObj.Headers))
	log.Print("Request Body: \n", reqObj.RequestBody)

	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return
	}
	defer res.Body.Close()

	resObj.StatusCode = res.StatusCode
	log.Printf("Response status: <%d>", res.StatusCode)

	var data []byte
	if data, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}

	resBody := string(data)
	resObj.ResStr = resBody

	log.Print("Response Body: \n", resBody)
	return
}
