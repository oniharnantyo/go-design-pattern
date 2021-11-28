package response_generate_test

import (
	"testing"

	"github.com/oniharnantyo/go-design-pattern/creational/factory/response_generate"
)

func TestGenerate(t *testing.T) {
	t.Run("", func(t *testing.T) {
		mockData := response_generate.Data{Name: "Bob"}
		expectedRes := `{"Name":"Bob"}`

		jsonGenerator := response_generate.NewResponseGenerator(response_generate.JSON)

		res := jsonGenerator.Generate(mockData)

		if res != expectedRes {
			t.Fail()
		}
	})
}
