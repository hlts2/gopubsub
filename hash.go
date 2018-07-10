package gopubsub

import (
	"hash/fnv"
	"unsafe"
)

func generateHash(text string) uint32 {
	h := fnv.New32()
	h.Write(*(*[]byte)(unsafe.Pointer(&text)))
	return h.Sum32()
}
