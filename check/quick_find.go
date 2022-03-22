package check

type QuickFind struct {
	parent []int
	count  int //连通分量的个数
}

// 查找v所属的集合(根节点)
func (p *QuickFind) Find(v int) int {
	if v < 0 || v > len(p.parent) {
		panic("索引超出范围")
	}
	return p.parent[v]
}

//合并v1,v2所属的集合
func (p *QuickFind) Uinon(v1, v2 int) {
	// Quick Find 的 union(v1, v2)：让 v1 所在集合的所有元素都指向 v2 的根节点

	p1 := p.Find(v1)
	p2 := p.Find(v2)
	for i, v := range p.parent {
		if v == p1 {
			p.parent[i] = p2
			p.count--
		}
	}
	return
}

//检查v1,v2是否属于同一集合
func (p *QuickFind) IsSame(v1, v2 int) bool {
	return p.Find(v1) == p.Find(v2)
}

//获取连通分量个数
func (p *QuickFind) GetCount() int {
	return p.count
}

//初始化
func (p *QuickFind) Init(cap int) {
	// 初始化时，每个元素各自属于一个单元素集合
	if cap < 0 {
		return
	}
	parent := make([]int, cap)
	for i, _ := range parent {
		parent[i] = i
	}
	p.parent = parent
	p.count = cap
	return
}

func (p *QuickFind) InitCustom(count int, parent []int) {
	// 初始化时，每个元素各自属于一个单元素集合
	if parent == nil || len(parent) < 0 {
		return
	}
	p.parent = parent
	p.count = count
	return
}

func NewQuickFind(cap int) *QuickFind {
	union := &QuickFind{}
	union.Init(cap)
	return union
}
