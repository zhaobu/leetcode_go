package core

type CaseIface interface {
	GetInput() interface{}
	GetName() string
	// Copy() CaseIface
	Print() string
}
