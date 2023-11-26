# go-once

This repository contains benchmarks for various implementations of the
[sync.Once](https://pkg.go.dev/sync#Once) type.

Result from a 2023 Apple M2 Pro:

```
goos: darwin
goarch: arm64
pkg: aryan.app/go-once
Benchmark/standard-lib-once/routines=1-10                        1605355               728.6 ns/op
Benchmark/standard-lib-once/routines=1000-10                        3222            359892 ns/op
Benchmark/standard-lib-once/routines=100000-10                        28          37600542 ns/op
Benchmark/standard-lib-once/routines=10000000-10                       1        3831679083 ns/op
Benchmark/mutex-only-once/routines=1-10                          1636492               733.3 ns/op
Benchmark/mutex-only-once/routines=1000-10                          2941            402762 ns/op
Benchmark/mutex-only-once/routines=100000-10                          28          41678208 ns/op
Benchmark/mutex-only-once/routines=10000000-10                         1        4152651250 ns/op
Benchmark/mutex-and-int32-atomic-once/routines=1-10              1636999               726.3 ns/op
Benchmark/mutex-and-int32-atomic-once/routines=1000-10              3379            369858 ns/op
Benchmark/mutex-and-int32-atomic-once/routines=100000-10              28          37577216 ns/op
Benchmark/mutex-and-int32-atomic-once/routines=10000000-10             1        3789052083 ns/op
Benchmark/mutex-and-int64-atomic-once/routines=1-10              1629019               740.5 ns/op
Benchmark/mutex-and-int64-atomic-once/routines=1000-10              3258            366026 ns/op
Benchmark/mutex-and-int64-atomic-once/routines=100000-10              28          38992088 ns/op
Benchmark/mutex-and-int64-atomic-once/routines=10000000-10             1        3843423500 ns/op
PASS
ok      aryan.app/go-once       32.985s
```
