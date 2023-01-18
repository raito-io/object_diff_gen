package test

type AnotherType struct {
	AnotherA int
	AnotherB string
}

//go:generate go run github.com/raito-io/object_diff_gen TestStruct
type TestStruct struct {
	IntValue    int
	StringValue string
	TypeValue   AnotherType

	IntPtr    *int
	StringPtr *string
	TypePtr   *AnotherType

	ArrayInt    []int
	ArrayObject []AnotherType
}
