package hash

import "testing"

var src = "asdhfioadsyfhidsfhjvdshfiapsdhciasdhcaisdghfoiuvasdhciafhcdisuahfciasudhfviadsuhgfiaesdujvcoikahsdiugfhasdpiofhadisughfvaisdfhcaiusdhfvnciuasdhfvniuasdhfviausdf"

func BenchmarkMd5(b *testing.B) {

	b.Log("MD5基准测试")
	b.ReportAllocs() // 内存开销
	// b.N为常规写法
	for i := 0; i < b.N; i++ {
		BaseMd5(src)
	}
}

func BenchmarkSha224(b *testing.B) {

	b.Log("Sha224基准测试")
	b.ReportAllocs() // 内存开销
	// b.N为常规写法
	for i := 0; i < b.N; i++ {
		BaseSha224(src)
	}
}

func BenchmarkSha256(b *testing.B) {

	b.Log("Sha256基准测试")
	b.ReportAllocs() // 内存开销
	// b.N为常规写法
	for i := 0; i < b.N; i++ {
		BaseSha256(src)
	}
}

func BenchmarkSha384(b *testing.B) {

	b.Log("Sha384基准测试")
	b.ReportAllocs() // 内存开销
	// b.N为常规写法
	for i := 0; i < b.N; i++ {
		BaseSha384(src)
	}
}

func BenchmarkSha512(b *testing.B) {

	b.Log("Sha512基准测试")
	b.ReportAllocs() // 内存开销
	// b.N为常规写法
	for i := 0; i < b.N; i++ {
		BaseSha512(src)
	}
}
