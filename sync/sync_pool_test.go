package sync

import "testing"

func TestAlloc(t *testing.T) {
	//以下代码用于演示【错误释放空间问题】
	Free([]byte{})
	Alloc(1)
}
