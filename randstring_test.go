package randstring

import (
	"testing"

	"github.com/kaibox-git/randstring/benchmarks/generate"
	"github.com/kaibox-git/randstring/benchmarks/make"
	"github.com/kaibox-git/randstring/benchmarks/randombase64"
	"github.com/kaibox-git/randstring/benchmarks/randseq"
	"github.com/kaibox-git/randstring/benchmarks/randstr"
	"github.com/kaibox-git/randstring/benchmarks/randstring"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	require.Len(t, Create(32), 32)
	require.Len(t, Create(64), 64)
}

func BenchmarkCreate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Create(32)
	}
}

func BenchmarkRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randstr.RandStr(32)
	}
}

func BenchmarkMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make.Make(32)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generate.Generate(32)
	}
}

func BenchmarkRandomBase64String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randombase64.RandomBase64String(32)
	}
}

func BenchmarkRandSeq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randseq.RandSeq(32)
	}
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = randstring.RandomString(32)
	}
}
