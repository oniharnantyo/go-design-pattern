package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonParser_Parse(t *testing.T) {
	xmlParser := NewXml()

	data := &jsonData{
		Name:    "Bob",
		Address: "A Street",
	}

	t.Run("xml", func(t *testing.T) {
		jsonToXmlParser := NewJsonToXml(xmlParser)
		jsonParser := NewJson(jsonToXmlParser)

		err := jsonParser.Parse(data)
		assert.NoError(t, err)
	})

	t.Run("json", func(t *testing.T) {
		jsonParser := NewJson(nil)
		err := jsonParser.Parse(data)

		assert.NoError(t, err)
	})

}
