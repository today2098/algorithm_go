package unionfind

type UnionFindIfs interface {
	Root(int) int
	Unite(int, int) (bool, error)
}

type UnionFind struct {
	vn  int
	gn  int
	par []int
}

func NewUnionFind(vn_ int) *UnionFind {
	uf := &UnionFind{
		vn: vn_,
		gn: vn_,
	}
	for i := 0; i < vn_; i++ {
		uf.par = append(uf.par, -1)
	}
	return uf
}

func (uf UnionFind) Vn() int {
	return uf.vn
}

func (uf UnionFind) Gn() int {
	return uf.gn
}

func (uf *UnionFind) Root(v int) int {
	if uf.par[v] < 0 {
		return v
	}
	uf.par[v] = uf.Root(uf.par[v])
	return uf.par[v]
}

func (uf *UnionFind) Size(v int) int {
	return -uf.par[uf.Root(v)]
}

func (uf *UnionFind) IsSame(u int, v int) bool {
	par := uf.Root(u)
	par2 := uf.Root(v)
	return par == par2
}

func (uf *UnionFind) Unite(u int, v int) bool {
	u = uf.Root(u)
	v = uf.Root(v)
	if u == v {
		return false
	}
	if uf.Size(u) < uf.Size(v) {
		u, v = v, u
	}
	uf.par[u] += uf.par[v]
	uf.par[v] = u
	uf.gn--
	return true
}

func (uf *UnionFind) Reset() {
	uf.gn = uf.vn
	for i := 0; i < uf.vn; i++ {
		uf.par[i] = -1
	}
}
