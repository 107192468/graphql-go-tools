package parser

import (
	"bytes"
	. "github.com/franela/goblin"
	"github.com/jensneuse/graphql-go-tools/pkg/document"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"testing"
)

func TestInputFieldsDefinitionParser(t *testing.T) {

	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("parser.parseInputFieldsDefinition", func() {

		tests := []struct {
			it           string
			input        string
			expectErr    types.GomegaMatcher
			expectValues types.GomegaMatcher
		}{
			{
				it:        "should parse a simple InputFieldsDefinition",
				input:     `{inputValue: Int}`,
				expectErr: BeNil(),
				expectValues: Equal(document.InputFieldsDefinition{
					document.InputValueDefinition{
						Name: "inputValue",
						Type: document.NamedType{
							Name: "Int",
						},
					},
				}),
			},
			{
				it:           "should not parse an optional InputFieldsDefinition",
				input:        ` `,
				expectErr:    BeNil(),
				expectValues: Equal(document.InputFieldsDefinition(nil)),
			},
			{
				it:        "should be able to parse multiple InputValueDefinitions within an InputFieldsDefinition",
				input:     `{inputValue: Int, outputValue: String}`,
				expectErr: BeNil(),
				expectValues: Equal(document.InputFieldsDefinition{
					document.InputValueDefinition{
						Name: "inputValue",
						Type: document.NamedType{
							Name: "Int",
						},
					},
					document.InputValueDefinition{
						Name: "outputValue",
						Type: document.NamedType{
							Name: "String",
						},
					},
				}),
			},
			{
				it:           "should return empty when no CURLYBRACKETOPEN at beginning (since it can be optional)",
				input:        `inputValue: Int}`,
				expectErr:    BeNil(),
				expectValues: Equal(document.InputFieldsDefinition(nil)),
			},
			{
				it:           "should fail when double CURLYBRACKETOPEN at beginning",
				input:        `{{inputValue: Int}`,
				expectErr:    Not(BeNil()),
				expectValues: Equal(document.InputFieldsDefinition(nil)),
			},
			{
				it:        "should fail when no CURLYBRACKETCLOSE at the end",
				input:     `{inputValue: Int`,
				expectErr: Not(BeNil()),
				expectValues: Equal(document.InputFieldsDefinition(document.InputFieldsDefinition{
					document.InputValueDefinition{
						Name: "inputValue",
						Type: document.NamedType{
							Name: "Int",
						},
					},
				})),
			},
		}

		for _, test := range tests {
			test := test

			g.It(test.it, func() {

				reader := bytes.NewReader([]byte(test.input))
				parser := NewParser()
				parser.l.SetInput(reader)

				val, err := parser.parseInputFieldsDefinition()
				Expect(err).To(test.expectErr)
				Expect(val).To(test.expectValues)
			})
		}
	})
}