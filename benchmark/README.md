
## Performance results

λ go test -bench=. -timeout 30m
goos: windows
goarch: amd64
pkg: github.com/Pencroff/JsonStruct/benchmark
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
Data size: 1.85 Mb
Benchmark_Unmarshal_code/Std_code-12                  37          32605884 ns/op          59.51 MB/s     3045344 B/op      92670 allocs/op
Benchmark_Unmarshal_code/Go_code-12                  186           5809410 ns/op         334.02 MB/s     3075540 B/op      13509 allocs/op
Benchmark_Unmarshal_code/Iter_code-12                126           9333640 ns/op         207.90 MB/s     2297118 B/op      52144 allocs/op
Benchmark_Unmarshal_code/JValue_code-12               58          20570531 ns/op          94.33 MB/s    18698258 B/op     128909 allocs/op
Benchmark_Unmarshal_code/Jay_code-12                 162           7322342 ns/op         265.01 MB/s     3346972 B/op      27874 allocs/op
Benchmark_Unmarshal_code/Simd_code-12                128           9343452 ns/op         207.68 MB/s    11929161 B/op         15 allocs/op
Data size: 2.15 Mb
Benchmark_Unmarshal_canada/StdJson___canada-12       28          40732946 ns/op          55.26 MB/s     5364424 B/op     173430 allocs/op
Benchmark_Unmarshal_canada/StdSimple_canada-12       28          38731521 ns/op          58.12 MB/s     5389324 B/op     172950 allocs/op
Benchmark_Unmarshal_canada/Go_canada-12              92          12796560 ns/op         175.91 MB/s     3669734 B/op      56537 allocs/op
Benchmark_Unmarshal_canada/GoSimple_canada-12        99          13390625 ns/op         168.11 MB/s     3691042 B/op      56539 allocs/op
Benchmark_Unmarshal_canada/IterSimple_canada-12      40          30309440 ns/op          74.27 MB/s     8007962 B/op     276069 allocs/op
Benchmark_Unmarshal_canada/JsonValue_canada-12       33          35055958 ns/op          64.21 MB/s    39460262 B/op     223513 allocs/op
Benchmark_Unmarshal_canada/Jay_canada-12             93          12426033 ns/op         181.16 MB/s     5724903 B/op     170259 allocs/op
Benchmark_Unmarshal_canada/Simd_canada-12            85          13035580 ns/op         172.69 MB/s     3040830 B/op         10 allocs/op
PASS
ok      github.com/Pencroff/JsonStruct/benchmark        21.620s


## JSON Data

Source from:

* https://github.com/miloyip/nativejson-benchmark
  * jsonchecker - folder with failed tests
  * roundtrip - folder with simple json validations
  * canada.json.gz
  * citm_catalog.json.gz
  * twitter.json.gz
* https://github.com/goccy/go-json/tree/master/benchmarks/testdata
  * code.json.gz
* https://github.com/mailru/easyjson/blob/master/benchmark
  * example.json.gz
* https://github.com/ultrajson/ultrajson/tree/main/tests
  * sample.json.gz - a lot of unicode characters
* https://github.com/json-iterator/test-data
  * large-file.json.gz
