---
data:
  _extendedDependsOn:
  - icon: ':warning:'
    path: lib/examples/sum_test.go
    title: lib/examples/sum_test.go
  - icon: ':warning:'
    path: lib/structure/unionfind/union_find.go
    title: lib/structure/unionfind/union_find.go
  - icon: ':warning:'
    path: lib/structure/unionfind/union_find_test.go
    title: lib/structure/unionfind/union_find_test.go
  _extendedRequiredBy:
  - icon: ':warning:'
    path: lib/examples/sum_test.go
    title: lib/examples/sum_test.go
  - icon: ':warning:'
    path: lib/structure/unionfind/union_find.go
    title: lib/structure/unionfind/union_find.go
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
    RuntimeError: bundler is not specified: lib/examples/sum.go\n"
  code: "package examples\n\nfunc Sum(a int, b int) int {\n\treturn a + b\n}\n"
  dependsOn:
  - lib/examples/sum_test.go
  - lib/structure/unionfind/union_find.go
  - lib/structure/unionfind/union_find_test.go
  isVerificationFile: false
  path: lib/examples/sum.go
  requiredBy:
  - lib/examples/sum_test.go
  - lib/structure/unionfind/union_find.go
  - lib/structure/unionfind/union_find_test.go
  timestamp: '2024-06-21 21:52:45+09:00'
  verificationStatus: LIBRARY_NO_TESTS
  verifiedWith: []
documentation_of: lib/examples/sum.go
layout: document
redirect_from:
- /library/lib/examples/sum.go
- /library/lib/examples/sum.go.html
title: lib/examples/sum.go
---
