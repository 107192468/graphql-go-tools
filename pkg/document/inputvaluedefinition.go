package document

// InputValueDefinition as specified in:
// http://facebook.github.io/graphql/draft/#InputValueDefinition
type InputValueDefinition struct {
	Description  string
	Name         string
	Type         Type
	DefaultValue Value
	Directives   Directives
}

// DefaultValue as specified in:
// http://facebook.github.io/graphql/draft/#DefaultValue
type DefaultValue Value