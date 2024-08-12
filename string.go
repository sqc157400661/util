package util

import (
	"reflect"
	"unsafe"
)

// BytesToString casts slice to string without copy
func BytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}

	p := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&b)).Data)
	var s string
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Data = uintptr(p)
	sh.Cap = len(b)
	sh.Len = len(b)
	return s
}

// StringToBytes casts string to slice without copy
func StringToBytes(s string) []byte {
	if len(s) == 0 {
		return []byte{}
	}

	p := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	var b []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh.Data = uintptr(p)
	sh.Cap = len(s)
	sh.Len = len(s)
	return b
}
