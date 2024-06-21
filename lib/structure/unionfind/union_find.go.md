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
    path: lib/structure/unionfind/union_find_test.go
    title: lib/structure/unionfind/union_find_test.go
  _extendedRequiredBy:
  - icon: ':warning:'
    path: lib/examples/sum.go
    title: lib/examples/sum.go
  - icon: ':warning:'
    path: lib/examples/sum_test.go
    title: lib/examples/sum_test.go
  - icon: ':warning:'
    path: lib/structure/unionfind/union_find_test.go
    title: lib/structure/unionfind/union_find_test.go
  _extendedVerifiedWith: []
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':warning:'
  attributes: {}
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: lib/structure/unionfind/union_find.go\n"
  code: "package unionfind\n\ntype UnionFindIfs interface {\n\tRoot(int) int\n\tUnite(int,\
    \ int) (bool, error)\n}\n\ntype UnionFind struct {\n\tvn  int\n\tgn  int\n\tpar\
    \ []int\n}\n\nfunc NewUnionFind(vn_ int) *UnionFind {\n\tuf := &UnionFind{\n\t\
    \tvn: vn_,\n\t\tgn: vn_,\n\t}\n\tfor i := 0; i < vn_; i++ {\n\t\tuf.par = append(uf.par,\
    \ -1)\n\t}\n\treturn uf\n}\n\nfunc (uf UnionFind) Vn() int {\n\treturn uf.vn\n\
    }\n\nfunc (uf UnionFind) Gn() int {\n\treturn uf.gn\n}\n\nfunc (uf *UnionFind)\
    \ Root(v int) int {\n\tif uf.par[v] < 0 {\n\t\treturn v\n\t}\n\tuf.par[v] = uf.Root(uf.par[v])\n\
    \treturn uf.par[v]\n}\n\nfunc (uf *UnionFind) Size(v int) int {\n\treturn -uf.par[uf.Root(v)]\n\
    }\n\nfunc (uf *UnionFind) IsSame(u int, v int) bool {\n\tpar := uf.Root(u)\n\t\
    par2 := uf.Root(v)\n\treturn par == par2\n}\n\nfunc (uf *UnionFind) Unite(u int,\
    \ v int) bool {\n\tu = uf.Root(u)\n\tv = uf.Root(v)\n\tif u == v {\n\t\treturn\
    \ false\n\t}\n\tif uf.Size(u) < uf.Size(v) {\n\t\tu, v = v, u\n\t}\n\tuf.par[u]\
    \ += uf.par[v]\n\tuf.par[v] = u\n\tuf.gn--\n\treturn true\n}\n\nfunc (uf *UnionFind)\
    \ Reset() {\n\tuf.gn = uf.vn\n\tfor i := 0; i < uf.vn; i++ {\n\t\tuf.par[i] =\
    \ -1\n\t}\n}\n"
  dependsOn:
  - lib/examples/sum_test.go
  - lib/examples/sum.go
  - lib/structure/unionfind/union_find_test.go
  isVerificationFile: false
  path: lib/structure/unionfind/union_find.go
  requiredBy:
  - lib/examples/sum_test.go
  - lib/examples/sum.go
  - lib/structure/unionfind/union_find_test.go
  timestamp: '2024-06-21 21:52:45+09:00'
  verificationStatus: LIBRARY_NO_TESTS
  verifiedWith: []
documentation_of: lib/structure/unionfind/union_find.go
layout: document
redirect_from:
- /library/lib/structure/unionfind/union_find.go
- /library/lib/structure/unionfind/union_find.go.html
title: lib/structure/unionfind/union_find.go
---
