# go-once

This repository contains benchmarks for various implementations of the
[sync.Once](https://pkg.go.dev/sync#Once) type.

Result from a 2023 Apple M2 Pro:

```
goos: darwin
goarch: arm64
pkg: aryan.app/go-once
Benchmark/standard-lib-once/routines=1/-10                       1516785               766.5 ns/op
Benchmark/standard-lib-once/routines=1000/-10                       1989            598515 ns/op
Benchmark/standard-lib-once/routines=100000/-10                       15          75617978 ns/op
Benchmark/mutex-only-once/routines=1/-10                         1450498               820.3 ns/op
Benchmark/mutex-only-once/routines=1000/-10                         1725            697687 ns/op
Benchmark/mutex-only-once/routines=100000/-10                         13          89216016 ns/op
Benchmark/mutex-and-int32-atomic-once/routines=1/-10             1456428               826.1 ns/op
Benchmark/mutex-and-int32-atomic-once/routines=1000/-10             1969            600127 ns/op
Benchmark/mutex-and-int32-atomic-once/routines=100000/-10             15          75732906 ns/op
Benchmark/mutex-and-int64-atomic-once/routines=1/-10             1475373               835.7 ns/op
Benchmark/mutex-and-int64-atomic-once/routines=1000/-10             1978            611932 ns/op
Benchmark/mutex-and-int64-atomic-once/routines=100000/-10             15          73864103 ns/op
PASS
ok      aryan.app/go-once       19.496s
```
