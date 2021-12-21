package check

type CaseIface interface {
	GetInput() interface{}
	GetName() string
	// Copy() CaseIface
	Print() string
}
