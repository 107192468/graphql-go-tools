package parser

import (
	"bytes"
	. "github.com/franela/goblin"
	document "github.com/jensneuse/graphql-go-tools/pkg/document"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"testing"
)

func TestObjectFieldParser(t *testing.T) {

	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("parser.parseObjectField", func() {

		tests := []struct {
			it               string
			input            string
			expectErr        types.GomegaMatcher
			expectFieldName  types.GomegaMatcher
			expectFieldValue types.GomegaMatcher
		}{
			{
				it:              "should parse simple object field",
				input:           `foo: "bar"`,
				expectErr:       BeNil(),
				expectFieldName: Equal("foo"),
				expectFieldValue: Equal(document.StringValue{
					Val: "bar",
				}),
			},
		}

		for _, test := range tests {
			test := test

			g.It(test.it, func() {

				reader := bytes.NewReader([]byte(test.input))
				parser := NewParser()
				parser.l.SetInput(reader)

				field, err := parser.parseObjectField()
				Expect(err).To(test.expectErr)
				Expect(field.Name).To(test.expectFieldName)
				Expect(field.Value).To(test.expectFieldValue)
			})
		}
	})
}