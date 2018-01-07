package bls12

// #cgo CFLAGS: -I${SRCDIR}/relic/include -I${SRCDIR}/build/include
// #cgo LDFLAGS: ${SRCDIR}/build/lib/librelic_s.a
// #include "relic_core.h"
// #include "relic_err.h"
import "C"
import (
	"os"
)

func init() {
	C.core_init()
	C.ep_param_set_any_pairf()
	checkError()
}

// With CHECK on, the program exits on the second uncaught(?) error,
// and there are functions like ep_read_bin that will cause two errors
// in a row without returning.
//
// With CHECK off there is no err_get_msg.
//
// Basically there's nothing we can do beyond keeping CHECK on, so that
// we see log+exit, and treat all errors as irrecoverable. YOLO.
//
// But anyway, if by mistake we cause one error and not two, we need
// to detonate ourselves. Sigh.
//
// Ah, and https://github.com/relic-toolkit/relic/issues/59.

func checkError() {
	if C.err_get_code() != C.STS_OK {
		var e *C.err_t
		var msg **C.char
		C.err_get_msg(e, msg)
		// errors.New(C.GoString(*msg))
		os.Exit(int(*e))
	}
}
