package effectivego
import (
	"sync"
	"bytes"
)
// new  or make

// new →　メモリーを初期化するのではなく0を埋める（？）ポインターを返す。
type SyncedBuffer struct {
	lock sync.Mutex
	buffer bytes.Buffer
}

//type *SyncedBufferを返す
p := new(SyncedBuffer)

// type SyncedBufferを返す
// 違いはなんぞ・・・？
var p2 SyncedBuffer

func newFile(fd int, name string) *File{
	if fd < 0 {
		return nil
	}

	// fd, name, number, somthingというラベルがついたものを返す
	// 複合リテラルという
	return &File{fd: fd, name: name, number: nil, somthing: 0}
}

// make →　newと違って、具体的な値を持って初期化される。ただ、ポインターは返らない。map, slice, channelに利用される
// 100個まで値が入る。
make([]int, 10, 100)