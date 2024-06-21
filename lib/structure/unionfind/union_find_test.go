package unionfind

import (
	"reflect"
	"testing"
)

func TestNewUnionFind(t *testing.T) {
	type args struct {
		vn_ int
	}
	tests := []struct {
		name string
		args args
		want *UnionFind
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUnionFind(tt.args.vn_); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnionFind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_Vn(t *testing.T) {
	type fields struct {
		vn  int
		gn  int
		par []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := UnionFind{
				vn:  tt.fields.vn,
				gn:  tt.fields.gn,
				par: tt.fields.par,
			}
			if got := uf.Vn(); got != tt.want {
				t.Errorf("UnionFind.Vn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_Gn(t *testing.T) {
	type fields struct {
		vn  int
		gn  int
		par []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := UnionFind{
				vn:  tt.fields.vn,
				gn:  tt.fields.gn,
				par: tt.fields.par,
			}
			if got := uf.Gn(); got != tt.want {
				t.Errorf("UnionFind.Gn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_Root(t *testing.T) {
	type fields struct {
		vn  int
		gn  int
		par []int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := &UnionFind{
				vn:  tt.fields.vn,
				gn:  tt.fields.gn,
				par: tt.fields.par,
			}
			if got := uf.Root(tt.args.v); got != tt.want {
				t.Errorf("UnionFind.Root() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_Size(t *testing.T) {
	type fields struct {
		vn  int
		gn  int
		par []int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := &UnionFind{
				vn:  tt.fields.vn,
				gn:  tt.fields.gn,
				par: tt.fields.par,
			}
			if got := uf.Size(tt.args.v); got != tt.want {
				t.Errorf("UnionFind.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_IsSame(t *testing.T) {
	type fields struct {
		vn  int
		gn  int
		par []int
	}
	type args struct {
		u int
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := &UnionFind{
				vn:  tt.fields.vn,
				gn:  tt.fields.gn,
				par: tt.fields.par,
			}
			if got := uf.IsSame(tt.args.u, tt.args.v); got != tt.want {
				t.Errorf("UnionFind.IsSame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_Unite(t *testing.T) {
	type fields struct {
		vn  int
		gn  int
		par []int
	}
	type args struct {
		u int
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := &UnionFind{
				vn:  tt.fields.vn,
				gn:  tt.fields.gn,
				par: tt.fields.par,
			}
			if got := uf.Unite(tt.args.u, tt.args.v); got != tt.want {
				t.Errorf("UnionFind.Unite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_Reset(t *testing.T) {
	type fields struct {
		vn  int
		gn  int
		par []int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := &UnionFind{
				vn:  tt.fields.vn,
				gn:  tt.fields.gn,
				par: tt.fields.par,
			}
			uf.Reset()
		})
	}
}
