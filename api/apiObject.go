package api

import (
	"GoWithTestAutomation/utils"
	"github.com/beevik/etree"
	"github.com/bitly/go-simplejson"
)

type SOAPAttrs struct {
	EnvAttrs,
	HeaderAttrs,
	BodyAttrs map[string]string
}

type SOAPObject struct {
	SOAPBody string
	RequestObject
	ResponseObject
}

type RESTObject struct {
	Method string
	RequestObject
	ResponseObject
}

type RequestObject struct {
	Endpoint string
}

type ResponseObject struct {
	StatusCode int
	ResStr     string
	ResMap     map[string]interface{}
}

func (so *SOAPObject) CreateSOAPBody(envns string, soapHeaderInJSON, soapBodyInJSON *simplejson.Json,
	filePath, delim, dupSymbol string, attrs SOAPAttrs) (err error) {
	createAttrs := func(ele *etree.Element, m map[string]string) {
		if m != nil {
			for k, v := range m {
				ele.CreateAttr(k, v)
			}
		}
	}

	doc := etree.NewDocument()
	// TODO: Should this be hard-coded?
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	// Create SOAPEnvelop
	envelop := doc.CreateElement(envns + ":Envelop")
	envelop.CreateAttr("xmlns:"+envns, "http://schemas.xmlsoap.org/soap/envelope/")
	createAttrs(envelop, attrs.EnvAttrs)

	// Create SOAPHeader
	if soapHeaderInJSON != nil {
		header := envelop.CreateElement(envns + ":Header")
		// TODO: Does SOAP Header possibly have attributes?
		createAttrs(header, attrs.HeaderAttrs)

		var subDoc *etree.Document
		if subDoc, err = utils.FlatJSON2XML(soapHeaderInJSON, delim, dupSymbol); err != nil {
			return
		}
		header.AddChild(subDoc)
	}

	// Create SOAPBody
	body := envelop.CreateElement(envns + ":Body")
	createAttrs(body, attrs.BodyAttrs)

	var flatJSON *simplejson.Json
	if filePath != "" {
		if flatJSON, err = utils.ReadJSONFile(filePath); err != nil {
			return
		}
	} else {
		flatJSON = soapBodyInJSON
	}

	var subDoc *etree.Document
	if subDoc, err = utils.FlatJSON2XML(flatJSON, delim, dupSymbol); err != nil {
		return
	}
	body.AddChild(subDoc)

	if so.SOAPBody, err = doc.WriteToString(); err != nil {
		return
	}
	return nil
}
