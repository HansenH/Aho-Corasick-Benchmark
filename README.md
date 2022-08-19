## [Aho-Corasick](github.com/HansenH/Aho-Corasick) Benchmark

Benchmark for  
[Aho-Corasick Multi-pattern Matching Algorithm](github.com/HansenH/Aho-Corasick)  
vs.  
iterative [KMP](github.com/paddie/gokmp)  
vs.  
iterative regexp  

Command:
```
go test -bench . -benchmem -count 1 -cpu 1 -benchtime 10s
```

Output example (for github.com/HansenH/Aho-Corasick v1.0.2):
```
goos: windows
goarch: amd64
pkg: test
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
BenchmarkAC                  297          40347125 ns/op         5757016 B/op      68265 allocs/op
BenchmarkIKMP                 36         340573469 ns/op        199766752 B/op   2466851 allocs/op
BenchmarkIRegexp              48         245563777 ns/op        12925600 B/op      97698 allocs/op
PASS
ok      test    40.994s
```

## LICENSE

MIT
