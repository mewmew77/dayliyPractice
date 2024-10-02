package sync

import (
	"fmt"
	"sync"
)

var pools []*sync.Pool

func Alloc(size int) []byte {
	zeros := powerOf2(size)
	b := *pools[zeros].Get().(*[]byte)
	if cap(b) < size {
		panic(fmt.Sprintf("%d < %d", cap(b), size))
	}
	return b[:size]
}

func Free(b []byte) {
	pools[powerOf2(cap(b))].Put(&b)
}

func powerOf2(s int) int {
	var bits int
	p := 1
	for p < s {
		bits++
		p *= 2
	}
	return bits
}

func init() {
	pools = make([]*sync.Pool, 30)
	for i := 0; i < 30; i++ {
		func(bits int) {
			pools[i] = &sync.Pool{
				New: func() interface{} {
					b := make([]byte, 1<<bits)
					return &b
				},
			}
		}(i)
	}
}
