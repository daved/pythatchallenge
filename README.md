# pyThat Challenge

    go get "github.com/daved/pythatchallenge"

From "[simple yet challenging problem]
(https://www.reddit.com/r/golang/comments/3y4onw/simple_yet_challenging_problem/)"
on [Reddit.com/r/golang](https://www.reddit.com/r/golang/), posted by 
[pyThat](https://www.reddit.com/user/pyThat).

## Benchmarks

BenchmarkA data:
```
[]string{"cp", `"file`, "name", "with", `spaces"`, `"file`, "name", `backup"`}
```

BenchmarkB data:
```
[]string{"one", "two", `"three`, "four", `five"`}
```

Results:
```
benchmark                   iter       time/iter   bytes alloc         allocs
---------                   ----       ---------   -----------         ------
BenchmarkA_DavedDev-8    2000000    664.00 ns/op      112 B/op    4 allocs/op
BenchmarkA_DJHerbis-8    5000000    280.00 ns/op       64 B/op    2 allocs/op
BenchmarkA_JasonMoo-8    1000000   1240.00 ns/op      144 B/op    3 allocs/op
BenchmarkA_SRT-8         1000000   1947.00 ns/op      336 B/op   16 allocs/op

BenchmarkB_DavedDev-8    3000000    453.00 ns/op       80 B/op    4 allocs/op
BenchmarkB_DJHerbis-8   10000000    158.00 ns/op       32 B/op    1 allocs/op
BenchmarkB_JasonMoo-8    2000000    816.00 ns/op      112 B/op    3 allocs/op
BenchmarkB_SRT-8         1000000   1166.00 ns/op      208 B/op    9 allocs/op
```
