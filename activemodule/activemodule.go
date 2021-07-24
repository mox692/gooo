package main

import (
	"fmt"
	"unsafe"
)

//go:linkname activeModules runtime.activeModules
func activeModules() []*moduledata

func main() {
	modules := activeModules()
	fmt.Printf("len(modukes) = %d,  \n", len(modules))
	fmt.Printf("modukes = %+v,  \n", modules[0])
}

type moduledata struct {
	pclntable    []byte
	ftab         []functab
	filetab      []uint32
	findfunctab  uintptr
	minpc, maxpc uintptr

	text, etext           uintptr
	noptrdata, enoptrdata uintptr
	data, edata           uintptr
	bss, ebss             uintptr
	noptrbss, enoptrbss   uintptr
	end, gcdata, gcbss    uintptr
	types, etypes         uintptr

	textsectmap []textsect
	typelinks   []int32 // offsets from types
	itablinks   []*itab

	ptab []ptabEntry

	pluginpath string
	pkghashes  []modulehash

	modulename   string
	modulehashes []modulehash

	hasmain uint8 // 1 if module contains the main function, 0 otherwise

	gcdatamask, gcbssmask bitvector

	typemap map[typeOff]*_type // offset to *_rtype in previous module

	bad bool // module failed to load and should be ignored

	next *moduledata
}
type functab struct {
	entry   uintptr
	funcoff uintptr
}
type textsect struct {
	vaddr    uintptr // prelinked section vaddr
	length   uintptr // section length
	baseaddr uintptr // relocated section address
}
type ptabEntry struct {
	name nameOff
	typ  typeOff
}

type nameOff int32
type typeOff int32
type textOff int32
type itab struct {
	inter *interfacetype
	_type *_type
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}
type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}
type tflag uint8
type interfacetype struct {
	typ     _type
	pkgpath name
	mhdr    []imethod
}
type name struct {
	bytes *byte
}
type imethod struct {
	name nameOff
	ityp typeOff
}
type modulehash struct {
	modulename   string
	linktimehash string
	runtimehash  *string
}
type bitvector struct {
	n        int32 // # of bits
	bytedata *uint8
}
