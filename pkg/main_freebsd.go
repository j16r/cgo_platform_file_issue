package pkg

// #include <stdlib.h>
import "C"
import "unsafe"

type thing struct{}

var Thing thing

func (t thing) Splat() {
	message := C.CString("splat!")
	defer C.free(unsafe.Pointer(message))
}
