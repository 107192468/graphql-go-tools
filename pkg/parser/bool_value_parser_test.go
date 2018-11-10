package parser

import (
	"bytes"
	. "github.com/franela/goblin"
	"github.com/jensneuse/graphql-go-tools/pkg/document"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"testing"
)

func TestBoolValueParser(t *testing.T) {

	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("parser.parseBoolValue", func() {

		tests := []struct {
			it           string
			input        string
			expectErr    types.GomegaMatcher
			expectValues types.GomegaMatcher
		}{
			{
				it:        "should parse a true boolean",
				input:     "true",
				expectErr: BeNil(),
				expectValues: Equal(document.BooleanValue{
					Val: true,
				}),
			},
			{
				it:        "should parse a false boolean",
				input:     "false",
				expectErr: BeNil(),
				expectValues: Equal(document.BooleanValue{
					Val: false,
				}),
			},
		}

		for _, test := range tests {
			test := test

			g.It(test.it, func() {

				reader := bytes.NewReader([]byte(test.input))
				parser := NewParser()
				parser.l.SetInput(reader)

				val, err := parser.parseBoolValue()
				Expect(err).To(test.expectErr)
				Expect(val).To(test.expectValues)
			})
		}
	})
}