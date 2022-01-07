package make

import (
	"math/rand"
	"sync"
	"time"
	"unsafe"
)

type rnd struct {
	pool *sync.Pool
}

var (
	mrand = &rnd{
		pool: &sync.Pool{
			New: func() interface{} {
				return rand.New(rand.NewSource(time.Now().UnixNano()))
			},
		},
	}
	abc = []byte(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`)
)

func (s *rnd) read(p []byte) (int, error) {
	r := s.pool.Get().(*rand.Rand)
	defer s.pool.Put(r)

	return r.Read(p)
}

func Make(len int) string {
	b := make([]byte, len)
	_, err := mrand.read(b)
	if err != nil {
		return ``
	}
	for i := 0; i < len; i++ {
		b[i] = abc[b[i]%(62)]
	}
	return *(*string)(unsafe.Pointer(&b))
}
