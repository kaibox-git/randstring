# randstring

randstring.Create () is fast and has minimal memory allocation. It returns a random alphanumeric string of a given length.

## Install

```
go get github.com/kaibox-git/randstring
```

## Usage

```go
import "github.com/kaibox-git/randstring"

...

for i:=0; i<3; i++{
    s := randstring.Create(32)
    println(s)
}
```

Example output:

```
drJEmUtKP9MDqlht50UBphGgPlDPypxJ
CAGukS8enZvT2d7a2VpqEzdqz8aCEceS
cIWj07dy3tk7Lc6i1e88TGWjX2N0Z2d1
```

## Benchmarks

See benchmark details in [randstring_test.go](github.com/kaibox-git/randstring/randstring_test.go)

```
BenchmarkCreate-8               	 7968715	       141.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkRandStr-8              	 7608330	       161.9 ns/op	      32 B/op	       1 allocs/op
BenchmarkMake-8                 	 6583093	       185.8 ns/op	      32 B/op	       1 allocs/op
BenchmarkGenerate-8             	 1575364	       778.5 ns/op	      32 B/op	       1 allocs/op
BenchmarkRandomBase64String-8   	 1228260	       936.3 ns/op	     128 B/op	       3 allocs/op
BenchmarkRandSeq-8              	  995644	      1182 ns/op	     176 B/op	       2 allocs/op
BenchmarkRandString-8           	  122017	      9434 ns/op	     120 B/op	       3 allocs/op
```