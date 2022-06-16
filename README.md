# JsonStruct

Json / JS dynamic structure. Play with Go.

[Docs](https://pkg.go.dev/github.com/Pencroff/JsonStruct)

### How to increase version

* commit all required changes
* git tag <version - v0.0.2>
* git push origin --tags
* done - check docs on [pkg.go.dev](https://pkg.go.dev/github.com/Pencroff/JsonStruct)
* install by `go get -u github.com/Pencroff/JsonStruct` 

## ToDo

* ToJson from test to suite
* Full set of functionality

## Benchmarks

### String to Int/Uint

    λ go test -bench=BenchmarkStringToIntConverter -benchtime=60s
    goos: windows
    goarch: amd64
    pkg: github.com/Pencroff/JsonStruct/helper
    cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
    BenchmarkStringToIntConverter/StringToUint-12           1000000000              18.85 ns/op
    BenchmarkStringToIntConverter/ParseUint-12              1000000000              34.32 ns/op
    BenchmarkStringToIntConverter/ParseUint(0)-12           1000000000              34.21 ns/op
    -------------------------------------------------------------------------------------------
    BenchmarkStringToIntConverter/StringToInt-12            1000000000              19.44 ns/op
    BenchmarkStringToIntConverter/strconv.Atoi-12           1000000000              37.32 ns/op
    BenchmarkStringToIntConverter/ParseInt-12               1000000000              31.19 ns/op
    BenchmarkStringToIntConverter/ParseInt(0)-12            1000000000              31.23 ns/op
    PASS
    ok      github.com/Pencroff/JsonStruct/helper   228.182s

    λ go test -bench=BenchmarkStringToIntConverter -benchtime=60s
    goos: windows
    goarch: amd64
    pkg: github.com/Pencroff/JsonStruct/helper
    cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
    BenchmarkStringToIntConverter/StringToUint-12           1000000000              18.87 ns/op
    BenchmarkStringToIntConverter/ParseUint-12              1000000000              34.68 ns/op
    BenchmarkStringToIntConverter/ParseUint(0)-12           1000000000              33.35 ns/op
    -------------------------------------------------------------------------------------------
    BenchmarkStringToIntConverter/StringToInt-12            1000000000              18.03 ns/op
    BenchmarkStringToIntConverter/strconv.Atoi-12           1000000000              38.97 ns/op
    BenchmarkStringToIntConverter/ParseInt-12               1000000000              36.22 ns/op
    BenchmarkStringToIntConverter/ParseInt(0)-12            1000000000              31.24 ns/op
    PASS
    ok      github.com/Pencroff/JsonStruct/helper   232.982s

    λ go test -bench=BenchmarkStringToIntConverter -benchtime=60s
    goos: windows
    goarch: amd64
    pkg: github.com/Pencroff/JsonStruct/helper
    cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
    BenchmarkStringToIntConverter/StringToUint-12           1000000000              20.16 ns/op
    BenchmarkStringToIntConverter/ParseUint-12              1000000000              34.57 ns/op
    BenchmarkStringToIntConverter/ParseUint(0)-12           1000000000              33.71 ns/op
    -------------------------------------------------------------------------------------------
    BenchmarkStringToIntConverter/StringToInt-12            1000000000              18.33 ns/op
    BenchmarkStringToIntConverter/strconv.Atoi-12           1000000000              37.69 ns/op
    BenchmarkStringToIntConverter/ParseInt-12               1000000000              36.57 ns/op
    BenchmarkStringToIntConverter/ParseInt(0)-12            1000000000              33.81 ns/op
    PASS
    ok      github.com/Pencroff/JsonStruct/helper   236.663s

    λ go test -bench=BenchmarkStringToIntConverter
    goos: windows
    goarch: amd64
    pkg: github.com/Pencroff/JsonStruct/helper
    cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
    BenchmarkStringToIntConverter/StringToUint-12           75411937                17.45 ns/op
    BenchmarkStringToIntConverter/ParseUint-12              36364627                34.23 ns/op
    BenchmarkStringToIntConverter/ParseUint(0)-12           33652942                29.78 ns/op
    -------------------------------------------------------------------------------------------
    BenchmarkStringToIntConverter/StringToInt-12            63043752                19.29 ns/op
    BenchmarkStringToIntConverter/strconv.Atoi-12           31566403                40.02 ns/op
    BenchmarkStringToIntConverter/ParseInt-12               31126708                37.28 ns/op
    BenchmarkStringToIntConverter/ParseInt(0)-12            34734181                34.11 ns/op
    PASS
    ok      github.com/Pencroff/JsonStruct/helper   9.090s