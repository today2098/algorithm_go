---
data:
  _extendedDependsOn:
  - icon: ':warning:'
    path: lib/examples/sum.go
    title: lib/examples/sum.go
  - icon: ':warning:'
    path: lib/examples/sum_test.go
    title: lib/examples/sum_test.go
  - icon: ':warning:'
    path: lib/structure/unionfind/union_find.go
    title: lib/structure/unionfind/union_find.go
  _extendedRequiredBy:
  - icon: ':warning:'
    path: lib/examples/sum.go
    title: lib/examples/sum.go
  - icon: ':warning:'
    path: lib/examples/sum_test.go
    title: lib/examples/sum_test.go
  - icon: ':warning:'
    path: lib/structure/unionfind/union_find.go
    title: lib/structure/unionfind/union_find.go
  _extendedVerifiedWith: []
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':warning:'
  attributes: {}
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: lib/structure/unionfind/union_find_test.go\n"
  code: "package unionfind\n\nimport (\n\t\"reflect\"\n\t\"testing\"\n)\n\nfunc TestNewUnionFind(t\
    \ *testing.T) {\n\ttype args struct {\n\t\tvn_ int\n\t}\n\ttests := []struct {\n\
    \t\tname string\n\t\targs args\n\t\twant *UnionFind\n\t}{\n\t\t// TODO: Add test\
    \ cases.\n\t}\n\tfor _, tt := range tests {\n\t\tt.Run(tt.name, func(t *testing.T)\
    \ {\n\t\t\tif got := NewUnionFind(tt.args.vn_); !reflect.DeepEqual(got, tt.want)\
    \ {\n\t\t\t\tt.Errorf(\"NewUnionFind() = %v, want %v\", got, tt.want)\n\t\t\t\
    }\n\t\t})\n\t}\n}\n\nfunc TestUnionFind_Vn(t *testing.T) {\n\ttype fields struct\
    \ {\n\t\tvn  int\n\t\tgn  int\n\t\tpar []int\n\t}\n\ttests := []struct {\n\t\t\
    name   string\n\t\tfields fields\n\t\twant   int\n\t}{\n\t\t// TODO: Add test\
    \ cases.\n\t}\n\tfor _, tt := range tests {\n\t\tt.Run(tt.name, func(t *testing.T)\
    \ {\n\t\t\tuf := UnionFind{\n\t\t\t\tvn:  tt.fields.vn,\n\t\t\t\tgn:  tt.fields.gn,\n\
    \t\t\t\tpar: tt.fields.par,\n\t\t\t}\n\t\t\tif got := uf.Vn(); got != tt.want\
    \ {\n\t\t\t\tt.Errorf(\"UnionFind.Vn() = %v, want %v\", got, tt.want)\n\t\t\t\
    }\n\t\t})\n\t}\n}\n\nfunc TestUnionFind_Gn(t *testing.T) {\n\ttype fields struct\
    \ {\n\t\tvn  int\n\t\tgn  int\n\t\tpar []int\n\t}\n\ttests := []struct {\n\t\t\
    name   string\n\t\tfields fields\n\t\twant   int\n\t}{\n\t\t// TODO: Add test\
    \ cases.\n\t}\n\tfor _, tt := range tests {\n\t\tt.Run(tt.name, func(t *testing.T)\
    \ {\n\t\t\tuf := UnionFind{\n\t\t\t\tvn:  tt.fields.vn,\n\t\t\t\tgn:  tt.fields.gn,\n\
    \t\t\t\tpar: tt.fields.par,\n\t\t\t}\n\t\t\tif got := uf.Gn(); got != tt.want\
    \ {\n\t\t\t\tt.Errorf(\"UnionFind.Gn() = %v, want %v\", got, tt.want)\n\t\t\t\
    }\n\t\t})\n\t}\n}\n\nfunc TestUnionFind_Root(t *testing.T) {\n\ttype fields struct\
    \ {\n\t\tvn  int\n\t\tgn  int\n\t\tpar []int\n\t}\n\ttype args struct {\n\t\t\
    v int\n\t}\n\ttests := []struct {\n\t\tname   string\n\t\tfields fields\n\t\t\
    args   args\n\t\twant   int\n\t}{\n\t\t// TODO: Add test cases.\n\t}\n\tfor _,\
    \ tt := range tests {\n\t\tt.Run(tt.name, func(t *testing.T) {\n\t\t\tuf := &UnionFind{\n\
    \t\t\t\tvn:  tt.fields.vn,\n\t\t\t\tgn:  tt.fields.gn,\n\t\t\t\tpar: tt.fields.par,\n\
    \t\t\t}\n\t\t\tif got := uf.Root(tt.args.v); got != tt.want {\n\t\t\t\tt.Errorf(\"\
    UnionFind.Root() = %v, want %v\", got, tt.want)\n\t\t\t}\n\t\t})\n\t}\n}\n\nfunc\
    \ TestUnionFind_Size(t *testing.T) {\n\ttype fields struct {\n\t\tvn  int\n\t\t\
    gn  int\n\t\tpar []int\n\t}\n\ttype args struct {\n\t\tv int\n\t}\n\ttests :=\
    \ []struct {\n\t\tname   string\n\t\tfields fields\n\t\targs   args\n\t\twant\
    \   int\n\t}{\n\t\t// TODO: Add test cases.\n\t}\n\tfor _, tt := range tests {\n\
    \t\tt.Run(tt.name, func(t *testing.T) {\n\t\t\tuf := &UnionFind{\n\t\t\t\tvn:\
    \  tt.fields.vn,\n\t\t\t\tgn:  tt.fields.gn,\n\t\t\t\tpar: tt.fields.par,\n\t\t\
    \t}\n\t\t\tif got := uf.Size(tt.args.v); got != tt.want {\n\t\t\t\tt.Errorf(\"\
    UnionFind.Size() = %v, want %v\", got, tt.want)\n\t\t\t}\n\t\t})\n\t}\n}\n\nfunc\
    \ TestUnionFind_IsSame(t *testing.T) {\n\ttype fields struct {\n\t\tvn  int\n\t\
    \tgn  int\n\t\tpar []int\n\t}\n\ttype args struct {\n\t\tu int\n\t\tv int\n\t\
    }\n\ttests := []struct {\n\t\tname   string\n\t\tfields fields\n\t\targs   args\n\
    \t\twant   bool\n\t}{\n\t\t// TODO: Add test cases.\n\t}\n\tfor _, tt := range\
    \ tests {\n\t\tt.Run(tt.name, func(t *testing.T) {\n\t\t\tuf := &UnionFind{\n\t\
    \t\t\tvn:  tt.fields.vn,\n\t\t\t\tgn:  tt.fields.gn,\n\t\t\t\tpar: tt.fields.par,\n\
    \t\t\t}\n\t\t\tif got := uf.IsSame(tt.args.u, tt.args.v); got != tt.want {\n\t\
    \t\t\tt.Errorf(\"UnionFind.IsSame() = %v, want %v\", got, tt.want)\n\t\t\t}\n\t\
    \t})\n\t}\n}\n\nfunc TestUnionFind_Unite(t *testing.T) {\n\ttype fields struct\
    \ {\n\t\tvn  int\n\t\tgn  int\n\t\tpar []int\n\t}\n\ttype args struct {\n\t\t\
    u int\n\t\tv int\n\t}\n\ttests := []struct {\n\t\tname   string\n\t\tfields fields\n\
    \t\targs   args\n\t\twant   bool\n\t}{\n\t\t// TODO: Add test cases.\n\t}\n\t\
    for _, tt := range tests {\n\t\tt.Run(tt.name, func(t *testing.T) {\n\t\t\tuf\
    \ := &UnionFind{\n\t\t\t\tvn:  tt.fields.vn,\n\t\t\t\tgn:  tt.fields.gn,\n\t\t\
    \t\tpar: tt.fields.par,\n\t\t\t}\n\t\t\tif got := uf.Unite(tt.args.u, tt.args.v);\
    \ got != tt.want {\n\t\t\t\tt.Errorf(\"UnionFind.Unite() = %v, want %v\", got,\
    \ tt.want)\n\t\t\t}\n\t\t})\n\t}\n}\n\nfunc TestUnionFind_Reset(t *testing.T)\
    \ {\n\ttype fields struct {\n\t\tvn  int\n\t\tgn  int\n\t\tpar []int\n\t}\n\t\
    tests := []struct {\n\t\tname   string\n\t\tfields fields\n\t}{\n\t\t// TODO:\
    \ Add test cases.\n\t}\n\tfor _, tt := range tests {\n\t\tt.Run(tt.name, func(t\
    \ *testing.T) {\n\t\t\tuf := &UnionFind{\n\t\t\t\tvn:  tt.fields.vn,\n\t\t\t\t\
    gn:  tt.fields.gn,\n\t\t\t\tpar: tt.fields.par,\n\t\t\t}\n\t\t\tuf.Reset()\n\t\
    \t})\n\t}\n}\n"
  dependsOn:
  - lib/examples/sum_test.go
  - lib/examples/sum.go
  - lib/structure/unionfind/union_find.go
  isVerificationFile: false
  path: lib/structure/unionfind/union_find_test.go
  requiredBy:
  - lib/examples/sum_test.go
  - lib/examples/sum.go
  - lib/structure/unionfind/union_find.go
  timestamp: '2024-06-21 21:52:45+09:00'
  verificationStatus: LIBRARY_NO_TESTS
  verifiedWith: []
documentation_of: lib/structure/unionfind/union_find_test.go
layout: document
redirect_from:
- /library/lib/structure/unionfind/union_find_test.go
- /library/lib/structure/unionfind/union_find_test.go.html
title: lib/structure/unionfind/union_find_test.go
---
