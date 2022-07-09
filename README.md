# JsonStruct

Json / JS dynamic structure. Play with Go.

[Docs](https://pkg.go.dev/github.com/Pencroff/JsonStruct)

[Standard](https://www.ecma-international.org/publications-and-standards/standards/ecma-404/)

### Parsing

* [My Golang JSON Evaluation](https://medium.com/geekculture/my-golang-json-evaluation-20a9ca6ef79c)
  * [Benchmarks](https://github.com/slaise/GoJsonBenchmark)

#### Byte array core

* https://github.com/buger/jsonparser
* https://github.com/tidwall/gjson

#### Marshallers

* https://github.com/mailru/easyjson
* https://github.com/pkg/json - abstractions for json
* https://github.com/francoispqt/gojay
* https://github.com/goccy/go-json
* "-------------------------------"
* https://github.com/minio/simdjson-go

### Data access speed

* Primitive values
  * Bool implemented in same way on both sides so no difference in speed - 0.5 - 0.6 ns (instantly)
  * Value implementation
    * Write: 0.5 - 1.2 ns
    * Read set type: 1.7 - 2.7 ns 
    * Check: 1.2 - 2.3 ns
  * Pointer implementation
    * Write: 15 - 36 ns
    * Read set type: 1.7 - 3.0 ns
    * Check: 1.2 - 2.3 ns
* Object operations
  * Value implementation
    * Set key: 392 - 497 ns (22% - 48% slower than go map)
    * Get key: 67 - 77 ns (43% - 57% slower than go map)
    * Check: 70 - 72 ns (47% - 49% slower than go map)
  * Pointer implementation
    * Set key: 409 - 541 ns
      * 27% - 61% slower than go map
      * 4% - 8% slower than value implementation
    * Get key: 71 - 79 ns
      * 51% - 61% slower than go map
      * 3% - 6% slower than value implementation
    * Check: 71 - 74 ns
      * 51% slower than go map
      * 1% - 3% slower than value implementation
  * Go Map implementation
    * Set key: 321 - 336 ns
    * Get key: 47 - 49 ns
    * Check: 47 - 49 ns
* Array operations
  * Value implementation
    * Push: 70 - 124 ns (84% - 143% slower than go slice)
    * Pop: 2.1 - 2.5 ns (similar to go slice)
  * Pointer implementation
    * Push: 91 - 132 ns (139% - 159% slower than go slice)
    * Pop: 29 - 31 ns (1100% - 1200% slower than value implementation)
  * Go slice implementation
    * Push: 38 - 51 ns
    * Pop: 2.3 ns (single measurement)

### Memory

* Primitive values:
  * Value implementation: 112 bytes
  * Pointer implementation: 16 - 40 bytes depend on the value
  
   From 2.8 to 7 times larger value then pointer implementation.

* Object values (7 elements with all primitive types):
    * Value implementation: 160 bytes empty object and 1232 bytes with set of fields 
    * Pointer implementation: 72 bytes empty object and 536 bytes with set of fields

     From 2.2 to 2.3 times larger value then pointer implementation.

* Array values:
    * Value implementation: 112 bytes empty array and 1136 bytes with set of elements
    * Pointer implementation: 40 bytes empty array and 624 bytes with set of elements

    From 1.8 to 2.8 times larger value then pointer implementation.

### Summary

The value implementation is faster than pointer implementation.
The difference is due to the fact that value implementation is less memory efficient.
Value implementation should expect to consume approximately 2 time more memory than pointer implementation (depends on stored values).
The performance of pointer implementation significantly slow in array implementation the rest operations in same range or similar to value implementation.
For performance please handle maps and slices.

## How to increase version

* commit all required changes
* git tag <version - v0.0.2>
* git push origin --tags
* done - check docs on [pkg.go.dev](https://pkg.go.dev/github.com/Pencroff/JsonStruct)
* install by `go get -u github.com/Pencroff/JsonStruct` 

## ToDo

* Parse json to struct
* ToJson from test to suite
* Make diff
* Evaluate circular dependencies and return error
* MsgPack
