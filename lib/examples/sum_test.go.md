---
data:
  _extendedDependsOn:
  - icon: ':warning:'
    path: lib/examples/sum.go
    title: lib/examples/sum.go
  - icon: ':warning:'
    path: lib/structure/unionfind/union_find.go
    title: lib/structure/unionfind/union_find.go
  - icon: ':warning:'
    path: lib/structure/unionfind/union_find_test.go
    title: lib/structure/unionfind/union_find_test.go
  _extendedRequiredBy:
  - icon: ':warning:'
    path: lib/examples/sum.go
    title: lib/examples/sum.go
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
    RuntimeError: bundler is not specified: lib/examples/sum_test.go\n"
  code: "package examples\n\nimport \"testing\"\n\nfunc TestSum(t *testing.T) {\n\t\
    type args struct {\n\t\ta int\n\t\tb int\n\t}\n\ttests := []struct {\n\t\tname\
    \ string\n\t\targs args\n\t\twant int\n\t}{\n\t\t{\n\t\t\tname: \"Case 1\",\n\t\
    \t\targs: args{\n\t\t\t\ta: 1,\n\t\t\t\tb: 10,\n\t\t\t},\n\t\t\twant: 11,\n\t\t\
    },\n\t\t{\n\t\t\tname: \"Case 2\",\n\t\t\targs: args{\n\t\t\t\ta: -3,\n\t\t\t\t\
    b: 10,\n\t\t\t},\n\t\t\twant: 7,\n\t\t},\n\t\t{\n\t\t\tname: \"Case 3\",\n\t\t\
    \targs: args{\n\t\t\t\ta: 0,\n\t\t\t\tb: 10,\n\t\t\t},\n\t\t\twant: 10,\n\t\t\
    },\n\t}\n\tfor _, tt := range tests {\n\t\tt.Run(tt.name, func(t *testing.T) {\n\
    \t\t\tif got := Sum(tt.args.a, tt.args.b); got != tt.want {\n\t\t\t\tt.Errorf(\"\
    Sum() = %v, want %v\", got, tt.want)\n\t\t\t}\n\t\t})\n\t}\n}\n"
  dependsOn:
  - lib/examples/sum.go
  - lib/structure/unionfind/union_find.go
  - lib/structure/unionfind/union_find_test.go
  isVerificationFile: false
  path: lib/examples/sum_test.go
  requiredBy:
  - lib/examples/sum.go
  - lib/structure/unionfind/union_find.go
  - lib/structure/unionfind/union_find_test.go
  timestamp: '2024-06-21 21:52:45+09:00'
  verificationStatus: LIBRARY_NO_TESTS
  verifiedWith: []
documentation_of: lib/examples/sum_test.go
layout: document
redirect_from:
- /library/lib/examples/sum_test.go
- /library/lib/examples/sum_test.go.html
title: lib/examples/sum_test.go
---
