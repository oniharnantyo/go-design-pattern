package adapter

import (
	"encoding/json"
	"encoding/xml"
)

type XmlParser interface {
	toXML(data *xmlData) error
}

type xmlData struct {
	Name    string `xml:"name"`
	Address string `xml:"address"`
}

type xmlParse struct{}

func NewXml() XmlParser {
	return new(xmlParse)
}

func (p *xmlParse) toXML(data *xmlData) error {
	_, err := xml.Marshal(data)
	if err != nil {
		return err
	}

	return nil
}

type JsonToXmlAdapter interface {
	Parse(ata *jsonData) error
}

type jsonToXmlData struct {
	Name    string `json:"name" xml:"name"`
	Address string `json:"address" xml:"address"`
}

type jsonToXml struct {
	xml XmlParser
}

func NewJsonToXml(xml XmlParser) JsonToXmlAdapter {
	return &jsonToXml{xml}
}

func (j *jsonToXml) Parse(data *jsonData) error {
	xmlData := &xmlData{
		Name:    data.Name,
		Address: data.Address,
	}

	return j.xml.toXML(xmlData)
}

type jsonData struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type JsonParser interface {
	Parse(data *jsonData) error
}

type jsonParser struct {
	jsonToXml JsonToXmlAdapter
}

func NewJson(jsonToXml JsonToXmlAdapter) JsonParser {
	return &jsonParser{jsonToXml}
}

func (p *jsonParser) Parse(data *jsonData) error {
	if p.jsonToXml != nil {
		return p.jsonToXml.Parse(data)
	}

	_, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return nil
}
