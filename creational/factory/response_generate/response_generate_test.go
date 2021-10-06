package response_generate_test

import (
	"testing"

	"../response_generate"
)

func TestJSONGenerate(t *testing.T) {
	mockData := response_generate.Data{Name: "Bob"}
	expectedRes := `{"Name":"Bob"}`

	jsonGenerator := response_generate.NewResponseGenerator(response_generate.JSON)

	res := jsonGenerator.Generate(mockData)

	if res != expectedRes {
		t.Fail()
	}
}
