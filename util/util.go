package util

// #cgo LDFLAGS: json-fortran-lib/libjsonfortran.a -lm
// #include "wrapper.h"
import "C"

func Run(input string) {
	inputStr := C.CString(input)
	C.Run(inputStr)
}
