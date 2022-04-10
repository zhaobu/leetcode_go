package check

type CaseIface interface {
	GetInput() interface{}
	GetName() string
	// Copy() CaseIface
	Print() string
}

type UnionFind interface {
	Find(v int) int
	Uinon(v1, v2 int)
	IsSame(v1, v2 int) bool
	GetCount() int
	Init(cap int)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
