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
		want UnionFind
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

func TestUnionFind_Root(t *testing.T) {
	type fields struct {
		vn int
	}
	type args struct {
		u int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Case 1",
			fields: fields{
				vn: 10,
			},
			args: args{
				u: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := NewUnionFind(tt.fields.vn)
			if got := uf.Root(tt.args.u); got != tt.want {
				t.Errorf("UnionFind.Root() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_Unite(t *testing.T) {
	type fields struct {
		vn  int
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
			uf := UnionFind{
				vn:  tt.fields.vn,
				par: tt.fields.par,
			}
			if got := uf.Unite(tt.args.u, tt.args.v); got != tt.want {
				t.Errorf("UnionFind.Unite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionFind_IsSame(t *testing.T) {
	type fields struct {
		vn  int
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
			uf := UnionFind{
				vn:  tt.fields.vn,
				par: tt.fields.par,
			}
			if got := uf.IsSame(tt.args.u, tt.args.v); got != tt.want {
				t.Errorf("UnionFind.IsSame() = %v, want %v", got, tt.want)
			}
		})
	}
}
