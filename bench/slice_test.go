package bench

import "testing"

/*
go test -v -bench=BenchmarkAppendTo
goos: darwin
goarch: amd64
pkg: github.com/zhiruchen/go-common/bench
BenchmarkAppendToNOMakeSlice-8               100          12683145 ns/op
BenchmarkAppendToMakeSlice-8                 668           1909600 ns/op
BenchmarkAppendToCapSlice-8                  561           2250636 ns/op
*/
const numOfItems int32 = 1000000

/*
go test -v -bench=BenchmarkAppendToNOMakeSlice
goos: darwin
goarch: amd64
pkg: github.com/zhiruchen/go-common/bench
BenchmarkAppendToNOMakeSlice-8           1388947               801 ns/op

*/
func BenchmarkAppendToNOMakeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendToNOMakeSlice(numOfItems)
	}
}

/*
go test -v -bench=BenchmarkAppendToMakeSlice
goos: darwin
goarch: amd64
pkg: github.com/zhiruchen/go-common/bench
BenchmarkAppendToMakeSlice-8             3610646               304 ns/op
BenchmarkAppendToMakeSliceWithCap-8      4308080               302 ns/op

*/
func BenchmarkAppendToMakeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendToMakeSlice(numOfItems)
	}
}

/*

 */
func BenchmarkAppendToCapSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendToMakeSliceWithCap(numOfItems)
	}
}
