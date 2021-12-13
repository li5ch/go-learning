### BenchMark 使用

- 进行基准测试的文件必须以*_test.go的文件为结尾
- 参与Benchmark基准性能测试的方法必须以Benchmark为前缀 
- 性能测试命令为go test [参数]，比如go test -bench=. -benchmem



```markdown
-bench regexp	性能测试，支持表达式对测试函数进行筛选。
-benchmem	性能测试的时候显示测试函数的内存分配的统计信息
$ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: program/benchmark
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_test-12        7439091               152.0 ns/op           248 B/op          5 allocs/op
PASS
ok      promgram/benchmark    1.304s

```

上述结果各个字段的含义：

| 结果项            | 含义                                                         |
| ----------------- | :----------------------------------------------------------- |
| Benchmark_test-12 | **Benchmark_test** 是测试的函数名 **-12** 表示GOMAXPROCS（线程数）的值为12 |
| 7439091           | 表示一共执行了**7439091**次，即**b.N**的值                   |
| 152.0 ns/op       | 表示平均每次操作花费了**152.0纳秒**                          |
| 248B/op           | 表示每次操作申请了**248Byte**的内存申请                      |
| 5 allocs/op       | 表示每次操作申请了**5**次内存                                |

