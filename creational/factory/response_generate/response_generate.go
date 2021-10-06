package response_generate

import (
	"encoding/json"
	"encoding/xml"
)

type ResponseType int

const (
	JSON ResponseType = iota
	XML
)

type Data struct {
	Name string
}

type ResponseGenerator interface {
	Generate(data Data) string
}

type JSONResponse struct{}

func (r *JSONResponse) Generate(data Data) string {
	res, _ := json.Marshal(data)

	return string(res)
}

type XMLResponse struct{}

func (r *XMLResponse) Generate(data Data) string {
	res, _ := xml.Marshal(data)

	return string(res)
}

func NewResponseGenerator(responseType ResponseType) ResponseGenerator {
	switch responseType {
	case JSON:
		return &JSONResponse{}
	case XML:
		return &XMLResponse{}
	}

	return nil
}
