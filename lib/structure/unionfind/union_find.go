package unionfind

type UnionFindIfs interface {
	Root(int) int
	Unite(int, int) (bool, error)
}

type UnionFind struct {
	vn  int
	par []int
}

func NewUnionFind(vn_ int) UnionFind {
	uf := UnionFind{
		vn: vn_,
	}
	for i := 0; i < vn_; i++ {
		uf.par = append(uf.par, -1)
	}
	return uf
}

func (uf UnionFind) Root(u int) int {
	if uf.par[u] == -1 {
		return u
	}
	return uf.Root(uf.par[u])
}

func (uf UnionFind) Unite(u int, v int) bool {
	u = uf.Root(u)
	v = uf.Root(v)
	if u == v {
		return false
	}
	uf.par[v] = u
	return true
}

func (uf UnionFind) IsSame(u int, v int) bool {
	par := uf.Root(u)
	par2 := uf.Root(v)
	return par == par2
}
