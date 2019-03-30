package document

import "github.com/jensneuse/graphql-go-tools/pkg/lexing/position"

type Node interface {
	NodeName() ByteSliceReference
	NodeAlias() ByteSliceReference
	NodeDescription() ByteSliceReference
	NodeArgumentSet() int
	NodeArgumentsDefinition() int
	NodeDirectiveSet() int // Change Signature to int (DirectiveSet)
	NodeEnumValuesDefinition() EnumValueDefinitions
	NodeSelectionSet() int
	NodeFields() []int
	NodeFieldsDefinition() FieldDefinitions
	NodeFragmentSpreads() []int
	NodeInlineFragments() []int
	NodeVariableDefinitions() []int
	NodeType() int
	NodeOperationType() OperationType
	NodeValue() int
	NodeDefaultValue() int
	NodeImplementsInterfaces() []int
	InputValueDefinitionsNode
	TypeSystemDefinitionNode
	UnionTypeSystemDefinitionNode
	ValueNode
	PositionNode
	InputFieldsDefinitionNode
}

type TypeSystemDefinitionNode interface {
	NodeSchemaDefinition() SchemaDefinition
	NodeScalarTypeDefinitions() []int
	NodeObjectTypeDefinitions() []int
	NodeInterfaceTypeDefinitions() []int
	NodeUnionTypeDefinitions() []int
	NodeEnumTypeDefinitions() []int
	NodeInputObjectTypeDefinitions() []int
	NodeDirectiveDefinitions() []int
}

type UnionTypeSystemDefinitionNode interface {
	NodeUnionMemberTypes() []int
}

type ValueNode interface {
	NodeValueType() ValueType
	NodeValueReference() int
}

type PositionNode interface {
	NodePosition() position.Position
}

type InputValueDefinitionsNode interface {
	NodeInputValueDefinitions() InputValueDefinitions
}

type InputFieldsDefinitionNode interface {
	NodeInputFieldsDefinition() int
}
