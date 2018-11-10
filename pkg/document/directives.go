package document

// Directive as specified in:
// http://facebook.github.io/graphql/draft/#Directive
type Directive struct {
	Name      string
	Arguments Arguments
}

// Directives as specified in
// http://facebook.github.io/graphql/draft/#Directives
type Directives []Directive