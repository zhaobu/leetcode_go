package check

type QuickUnion struct {
	parent []int
	rank   []int
	count  int
}

// 查找v所属的集合(根节点)
func (p *QuickUnion) Find(v int) int {
	return 0
}

//合并v1,v2所属的集合
func (p *QuickUnion) Uinon(v1, v2 int) {
	return
}

//检查v1,v2是否属于同一集合
func (p *QuickUnion) IsSame(v1, v2 int) bool {
	return p.Find(v1) == p.Find(v2)
}

//获取连通分量个数
func (p *QuickUnion) GetCount() int {
	return p.count
}

func (p *QuickUnion) Init(cap int) {
	if cap < 0 {
		return
	}
	parent := make([]int, cap)
	for i, _ := range parent {
		parent[i] = i
	}
	return
}

func NewQuickUnion(cap int) *QuickUnion {
	union := &QuickUnion{}
	union.Init(cap)
	return union
}
