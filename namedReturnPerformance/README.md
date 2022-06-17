Performance measurement of named vs. unnamed return types.

Based on https://medium.com/@ramseyjiang_22278/90-of-gophers-dont-know-6-simplest-ways-to-optimise-code-2-83d5f1f06272

Unlike the conclusions in that article, I find no difference in performance
among the various implementations.

Note that I removed the `time` and `*.printf` invocations as they get in the
way of the benchmarking. Note that this makes `declareReturnNameTypeWithFmt`
identical to `declareReturnNameTypeWithLog`, so you can see the consistency in
the benchmark measurements.

Example output:

```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: namedReturnPerformance
cpu: Intel(R) Core(TM) i7-5557U CPU @ 3.10GHz
BenchmarkReturnTypeOnly-4                6815211               171.8 ns/op
BenchmarkReturnNameTypeWithFmt-4         7033896               168.5 ns/op
BenchmarkReturnNameTypeWithLog-4         7192170               166.2 ns/op
PASS
ok      namedReturnPerformance  4.264s
$ 
```
